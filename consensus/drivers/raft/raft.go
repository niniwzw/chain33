package raft

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"

	"context"

	"code.aliyun.com/chain33/chain33/types"
	"github.com/coreos/etcd/etcdserver/stats"
	"github.com/coreos/etcd/pkg/fileutil"
	typec "github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/raft"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/rafthttp"
	"github.com/coreos/etcd/snap"
	"github.com/coreos/etcd/wal"
	"github.com/coreos/etcd/wal/walpb"
	"github.com/golang/protobuf/proto"
	log "github.com/inconshreveable/log15"
)

const (
	IsLeader   = "0"
	LeaderIsOK = "1"
)

var (
	defaultSnapCount        uint64 = 10000
	snapshotCatchUpEntriesN uint64 = 10000
	leaderCache             uint64
	isReady                 bool
)

type raftNode struct {
	proposeC         <-chan *types.Block
	confChangeC      <-chan raftpb.ConfChange
	commitC          chan<- *types.Block
	errorC           chan<- error
	id               int
	bootstrapPeers   []string
	readOnlyPeers    []string
	addPeers         []string
	join             bool
	waldir           string
	snapdir          string
	getSnapshot      func() ([]byte, error)
	lastIndex        uint64
	confState        raftpb.ConfState
	snapshotIndex    uint64
	appliedIndex     uint64
	node             raft.Node
	raftStorage      *raft.MemoryStorage
	wal              *wal.WAL
	snapshotter      *snap.Snapshotter
	snapshotterReady chan *snap.Snapshotter
	snapCount        uint64
	transport        *rafthttp.Transport
	stopMu           sync.RWMutex
	stopc            chan struct{}
	httpstopc        chan struct{}
	httpdonec        chan struct{}
	validatorC       chan map[string]bool
	//用于判断该节点是否重启过
	restartC chan struct{}
}

func NewRaftNode(id int, join bool, peers []string, readOnlyPeers []string, addPeers []string, getSnapshot func() ([]byte, error), proposeC <-chan *types.Block,
	confChangeC <-chan raftpb.ConfChange) (<-chan *types.Block, <-chan error, <-chan *snap.Snapshotter, <-chan map[string]bool) {

	log.Info("Enter consensus raft")
	// commit channel
	commitC := make(chan *types.Block)
	errorC := make(chan error)
	//storage := raft.NewMemoryStorage()
	//var readOnlyPeers []string
	//readOnlyPeers :=[]string{"http://172.31.2.210:9021"}
	rc := &raftNode{
		proposeC:         proposeC,
		confChangeC:      confChangeC,
		commitC:          commitC,
		errorC:           errorC,
		id:               id,
		join:             join,
		bootstrapPeers:   peers,
		readOnlyPeers:    readOnlyPeers,
		addPeers:         addPeers,
		waldir:           fmt.Sprintf("chain33_raft-%d", id),
		snapdir:          fmt.Sprintf("chain33_raft-%d-snap", id),
		getSnapshot:      getSnapshot,
		snapCount:        defaultSnapCount,
		stopc:            make(chan struct{}),
		httpstopc:        make(chan struct{}),
		httpdonec:        make(chan struct{}),
		validatorC:       make(chan map[string]bool),
		snapshotterReady: make(chan *snap.Snapshotter, 1),
		restartC:         make(chan struct{}, 1),
	}
	go rc.startRaft()

	return commitC, errorC, rc.snapshotterReady, rc.validatorC
}

//  启动raft节点
func (rc *raftNode) startRaft() {
	// 有snapshot就打开，没有则创建
	if !fileutil.Exist(rc.snapdir) {
		if err := os.Mkdir(rc.snapdir, 0750); err != nil {
			log.Error("chain33_raft: cannot create dir for snapshot (%v)", err)
		}
	}

	rc.snapshotter = snap.New(rc.snapdir)
	rc.snapshotterReady <- rc.snapshotter

	oldwal := wal.Exist(rc.waldir)
	rc.wal = rc.replayWAL()

	rpeers := make([]raft.Peer, len(rc.bootstrapPeers))
	for i := range rpeers {
		rpeers[i] = raft.Peer{ID: uint64(i + 1)}
	}
	c := &raft.Config{
		ID:              uint64(rc.id),
		ElectionTick:    10,
		HeartbeatTick:   1,
		Storage:         rc.raftStorage,
		MaxSizePerMsg:   1024 * 1024,
		MaxInflightMsgs: 256,
		//设置成预投票，当节点重连时，可以快速地重新加入集群
		PreVote:     true,
		CheckQuorum: false,
	}
	if len(rc.readOnlyPeers) > 0 && rc.id > len(rc.bootstrapPeers) {
		rc.join = true
	}
	if oldwal {
		rc.restartC <- struct{}{}
		rc.node = raft.RestartNode(c)
	} else {
		startPeers := rpeers
		if rc.join {
			startPeers = nil
		}
		rc.node = raft.StartNode(c, startPeers)
	}

	rc.transport = &rafthttp.Transport{
		ID:          typec.ID(rc.id),
		ClusterID:   0x1000,
		Raft:        rc,
		ServerStats: stats.NewServerStats("", ""),
		LeaderStats: stats.NewLeaderStats(strconv.Itoa(rc.id)),
		ErrorC:      make(chan error),
	}

	rc.transport.Start()
	for i := range rc.bootstrapPeers {
		if i+1 != rc.id {
			rc.transport.AddPeer(typec.ID(i+1), []string{rc.bootstrapPeers[i]})
		}
	}

	// 启动网络监听
	go rc.serveRaft()
	go rc.serveChannels()

	//定时轮询watch leader 状态是否改变，更新validator
	go rc.updateValidator()
}

// 网络监听
func (rc *raftNode) serveRaft() {
	var peers []string
	peers = append(peers, rc.bootstrapPeers...)
	peers = append(peers, rc.readOnlyPeers...)
	peers = append(peers, rc.addPeers...)
	nodeURL, err := url.Parse(peers[rc.id-1])
	if err != nil {
		log.Error("raft: Failed parsing URL (%v)", err)
	}

	ln, err := newStoppableListener(nodeURL.Host, rc.httpstopc)
	if err != nil {
		log.Error("raft: Failed to listen rafthttp (%v)", err)
	}

	err = (&http.Server{Handler: rc.transport.Handler()}).Serve(ln)
	select {
	case <-rc.httpstopc:
	default:
		log.Error("raft: Failed to serve rafthttp (%v)", err)
	}
	close(rc.httpdonec)
}

func (rc *raftNode) serveChannels() {
	snapShot, err := rc.raftStorage.Snapshot()
	if err != nil {
		panic(err)
	}
	rc.confState = snapShot.Metadata.ConfState
	rc.snapshotIndex = snapShot.Metadata.Index
	rc.appliedIndex = snapShot.Metadata.Index

	defer rc.wal.Close()

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	go func() {
		var confChangeCount uint64 = 0
		// 通过propose和proposeConfchange方法往RaftNode发通知
		for rc.proposeC != nil && rc.confChangeC != nil {
			select {
			case prop, ok := <-rc.proposeC:
				if !ok {
					rc.proposeC = nil
				} else {
					out, err := proto.Marshal(prop)
					if err != nil {
						log.Error("failed to marshal block: ", err)
					}
					rc.node.Propose(context.TODO(), out)
				}

			case cc, ok := <-rc.confChangeC:
				if !ok {
					rc.confChangeC = nil
				} else {
					confChangeCount += 1
					cc.ID = confChangeCount
					rc.node.ProposeConfChange(context.TODO(), cc)
				}
			}
		}
		close(rc.stopc)
	}()
	// 从Ready()中接收数据
	for {
		select {
		case <-ticker.C:
			rc.node.Tick()
		case rd := <-rc.node.Ready():
			rc.wal.Save(rd.HardState, rd.Entries)
			if !raft.IsEmptySnap(rd.Snapshot) {
				rc.saveSnap(rd.Snapshot)
				rc.raftStorage.ApplySnapshot(rd.Snapshot)
				rc.publishSnapshot(rd.Snapshot)
			}
			rc.raftStorage.Append(rd.Entries)
			rc.transport.Send(rd.Messages)
			if ok := rc.publishEntries(rc.entriesToApply(rd.CommittedEntries)); !ok {
				rc.stop()
				return
			}
			rc.maybeTriggerSnapshot()

			rc.node.Advance()
		case err := <-rc.transport.ErrorC:
			rc.writeError(err)
			return

		case <-rc.stopc:
			rc.stop()
			return
		}
	}
}

func (rc *raftNode) updateValidator() {
	var validatorMap map[string]bool
	//这块监听后期需要根据场景进行优化
	time.Sleep(5 * time.Second)

	//用于标记readOnlyPeers是否已经被添加到集群中了
	flag := false
	isRestart := false
	ticker := time.NewTicker(50 * time.Millisecond)
	select {
	case <-rc.restartC:
		isRestart = true
		close(rc.restartC)
	case <-ticker.C:
		ticker.Stop()
	}
	for {
		validatorMap = make(map[string]bool)
		if rc.Leader() == raft.None {
			//TODO 当节点被删除时需要如下提示信息需要优化，删除节点是不可逆操作，因此需要记录被删节点信息
			rlog.Info("==============Leader is not ready==============")
			validatorMap[LeaderIsOK] = false
		} else {

			// 获取到leader Id,选主成功
			validatorMap[LeaderIsOK] = true
			if rc.id == int(rc.Leader()) {
				//leader选举出来之后即可添加addReadOnlyPeers
				if !flag && !isRestart {
					go rc.addReadOnlyPeers()
				}
				validatorMap[IsLeader] = true
			} else {
				validatorMap[IsLeader] = false
			}
			flag = true
		}
		time.Sleep(time.Second)
		rc.validatorC <- validatorMap
	}
}
func (rc *raftNode) Status() raft.Status {
	rc.stopMu.RLock()
	defer rc.stopMu.RUnlock()
	return rc.node.Status()
}

// isLeader checks if we are the leader or not, without the protection of lock
func (rc *raftNode) leader() uint64 {
	return rc.node.Status().Lead
}

// IsLeader checks if we are the leader or not, with the protection of lock
func (rc *raftNode) Leader() uint64 {
	rc.stopMu.RLock()
	defer rc.stopMu.RUnlock()
	return rc.leader()
}
func (rc *raftNode) replayWAL() *wal.WAL {
	log.Info("replaying WAL of member %d", rc.id)
	snapshot := rc.loadSnapshot()
	w := rc.openWAL(snapshot)
	_, st, ents, err := w.ReadAll()
	if err != nil {
		log.Error("chain33_raft: failed to read WAL (%v)", err)
	}
	rc.raftStorage = raft.NewMemoryStorage()
	if snapshot != nil {
		rc.raftStorage.ApplySnapshot(*snapshot)
	}
	rc.raftStorage.SetHardState(st)

	// append to storage so raft starts at the right place in log
	rc.raftStorage.Append(ents)
	// send nil once lastIndex is published so client knows commit channel is current
	if len(ents) > 0 {
		rc.lastIndex = ents[len(ents)-1].Index
	} else {
		rc.commitC <- nil
	}
	return w
}

func (rc *raftNode) loadSnapshot() *raftpb.Snapshot {
	snapshot, err := rc.snapshotter.Load()
	if err != nil && err != snap.ErrNoSnapshot {
		log.Error("chain33_raft: error loading snapshot (%v)", err)
	}
	return snapshot
}

func (rc *raftNode) saveSnap(snap raftpb.Snapshot) error {
	walSnap := walpb.Snapshot{
		Index: snap.Metadata.Index,
		Term:  snap.Metadata.Term,
	}
	if err := rc.wal.SaveSnapshot(walSnap); err != nil {
		return err
	}
	if err := rc.snapshotter.SaveSnap(snap); err != nil {
		return err
	}
	return rc.wal.ReleaseLockTo(snap.Metadata.Index)
}

func (rc *raftNode) maybeTriggerSnapshot() {
	if rc.appliedIndex-rc.snapshotIndex <= rc.snapCount {
		return
	}

	log.Info("start snapshot [applied index: %d | last snapshot index: %d]", rc.appliedIndex, rc.snapshotIndex)
	data, err := rc.getSnapshot()
	if err != nil {
		log.Error("Err happened when get snapshot")
	}
	snapShot, err := rc.raftStorage.CreateSnapshot(rc.appliedIndex, &rc.confState, data)
	if err != nil {
		panic(err)
	}
	if err := rc.saveSnap(snapShot); err != nil {
		panic(err)
	}

	compactIndex := uint64(1)
	if rc.appliedIndex > snapshotCatchUpEntriesN {
		compactIndex = rc.appliedIndex - snapshotCatchUpEntriesN
	}
	if err := rc.raftStorage.Compact(compactIndex); err != nil {
		panic(err)
	}

	log.Info("compacted log at index %d", compactIndex)
	rc.snapshotIndex = rc.appliedIndex
}

func (rc *raftNode) publishSnapshot(snapshotToSave raftpb.Snapshot) {
	if raft.IsEmptySnap(snapshotToSave) {
		return
	}

	log.Info("publishing snapshot at index %d", rc.snapshotIndex)
	defer log.Info("finished publishing snapshot at index %d", rc.snapshotIndex)

	if snapshotToSave.Metadata.Index <= rc.appliedIndex {
		log.Error("snapshot index [%d] should > progress.appliedIndex [%d] + 1", snapshotToSave.Metadata.Index, rc.appliedIndex)
	}
	rc.commitC <- nil // trigger kvstore to load snapshot

	rc.confState = snapshotToSave.Metadata.ConfState
	rc.snapshotIndex = snapshotToSave.Metadata.Index
	rc.appliedIndex = snapshotToSave.Metadata.Index
}

func (rc *raftNode) openWAL(snapshot *raftpb.Snapshot) *wal.WAL {
	if !wal.Exist(rc.waldir) {
		if err := os.Mkdir(rc.waldir, 0750); err != nil {
			log.Error("chain33_raft: cannot create dir for wal (%v)", err)
		}

		w, err := wal.Create(rc.waldir, nil)
		if err != nil {
			log.Error("chain33_raft: create wal error (%v)", err)
		}
		w.Close()
	}

	walsnap := walpb.Snapshot{}
	if snapshot != nil {
		walsnap.Index, walsnap.Term = snapshot.Metadata.Index, snapshot.Metadata.Term
	}
	log.Info("loading WAL at term %d and index %d", walsnap.Term, walsnap.Index)
	w, err := wal.Open(rc.waldir, walsnap)
	if err != nil {
		log.Error("chain33_raft: error loading wal (%v)", err)
	}

	return w
}

// 关闭http连接和channel
func (rc *raftNode) stop() {
	rc.stopHTTP()
	close(rc.commitC)
	close(rc.errorC)
	rc.node.Stop()
}

func (rc *raftNode) stopHTTP() {
	rc.transport.Stop()
	close(rc.httpstopc)
	<-rc.httpdonec
}

func (rc *raftNode) writeError(err error) {
	rc.stopHTTP()
	close(rc.commitC)
	rc.errorC <- err
	close(rc.errorC)
	rc.node.Stop()
}

// 往commit channel中写入commit log
func (rc *raftNode) publishEntries(ents []raftpb.Entry) bool {
	for i := range ents {
		switch ents[i].Type {
		case raftpb.EntryNormal:
			if len(ents[i].Data) == 0 {
				break
			}
			// 解码
			block := &types.Block{}
			if err := proto.Unmarshal(ents[i].Data, block); err != nil {
				log.Error("failed to unmarshal: ", err)
			}
			select {
			case rc.commitC <- block:
			case <-rc.stopc:
				return false
			}

		case raftpb.EntryConfChange:
			var cc raftpb.ConfChange
			cc.Unmarshal(ents[i].Data)
			rc.confState = *rc.node.ApplyConfChange(cc)
			switch cc.Type {
			case raftpb.ConfChangeAddNode:
				if len(cc.Context) > 0 {
					rc.transport.AddPeer(typec.ID(cc.NodeID), []string{string(cc.Context)})
				}
			case raftpb.ConfChangeRemoveNode:
				if cc.NodeID == uint64(rc.id) {
					log.Info("I've been removed from the cluster! Shutting down.")
					return false
				}
				rc.transport.RemovePeer(typec.ID(cc.NodeID))
			case raftpb.ConfChangeAddLearnerNode:
				if len(cc.Context) > 0 {
					rc.transport.AddPeer(typec.ID(cc.NodeID), []string{string(cc.Context)})
				}
				isReady = true
			}

		}

		rc.appliedIndex = ents[i].Index

		if ents[i].Index == rc.lastIndex {
			select {
			case rc.commitC <- nil:
			case <-rc.stopc:
				return false
			}
		}
	}
	return true
}

func (rc *raftNode) entriesToApply(ents []raftpb.Entry) (nents []raftpb.Entry) {
	if len(ents) == 0 {
		return
	}
	firstIdx := ents[0].Index
	if firstIdx > rc.appliedIndex+1 {
		log.Error("first index of committed entry[%d] should <= progress.appliedIndex[%d] 1", firstIdx, rc.appliedIndex)
	}
	if rc.appliedIndex-firstIdx+1 < uint64(len(ents)) {
		nents = ents[rc.appliedIndex-firstIdx+1:]
	}
	return
}

func (rc *raftNode) Process(ctx context.Context, m raftpb.Message) error {
	return rc.node.Step(ctx, m)
}
func (rc *raftNode) IsIDRemoved(id uint64) bool                           { return false }
func (rc *raftNode) ReportUnreachable(id uint64)                          {}
func (rc *raftNode) ReportSnapshot(id uint64, status raft.SnapshotStatus) {}
func (rc *raftNode) addReadOnlyPeers() {
	isReady = true
	//信息校验，防止是空数组
	if len(rc.readOnlyPeers) == 1 && rc.readOnlyPeers[0] == "" {
		return
	}
	for i, peer := range rc.readOnlyPeers {
		cc := raftpb.ConfChange{
			Type:    raftpb.ConfChangeAddLearnerNode,
			NodeID:  uint64(len(rc.bootstrapPeers) + i + 1),
			Context: []byte(peer),
		}
		//节点变更一次只能变更一个，因此需要检查是否添加成功
		if isReady {
			confChangeC <- cc
			isReady = false
			for {
				time.Sleep(500 * time.Millisecond)
				if isReady {
					break
				}
			}
		}
	}
}
