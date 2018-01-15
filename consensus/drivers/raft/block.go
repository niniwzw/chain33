package raft

import (
	"errors"
	"sync"
	"time"

	"code.aliyun.com/chain33/chain33/common/merkle"
	//"code.aliyun.com/chain33/chain33/consensus/drivers"
	"code.aliyun.com/chain33/chain33/queue"
	"code.aliyun.com/chain33/chain33/types"
	"code.aliyun.com/chain33/chain33/util"
	"github.com/coreos/etcd/snap"
	"github.com/golang/protobuf/proto"
	log "github.com/inconshreveable/log15"

	//"sync/atomic"
	//"golang.org/x/tools/go/gcimporter15/testdata"
	"fmt"
)

var (
	listSize     int = 10000
	zeroHash     [32]byte
	currentBlock *types.Block
	mulock       sync.Mutex
)

//type Miner interface {
//	CreateGenesisTx() []*types.Transaction
//	CreateBlock()
//}
type RaftClient struct {
	//TODO: BaseClient中有些参数访问不了，所以暂时不用baseClient
	//*drivers.BaseClient

	proposeC chan<- *types.Block
	qclient  queue.IClient
	q        *queue.Queue
	//minerStart   int32
	once         sync.Once
	Cfg          *types.Consensus
	currentBlock *types.Block
	mulock       sync.RWMutex
	//child        Miner

	//mu          sync.RWMutex
	//blockstore  *types.Block
	//qclient     queue.IClient
	//q           *queue.Queue
	commitC     <-chan *types.Block
	errorC      <-chan error
	snapshotter *snap.Snapshotter
	validatorC  <-chan bool
}

func NewBlockstore(cfg *types.Consensus, snapshotter *snap.Snapshotter, proposeC chan<- *types.Block, commitC <-chan *types.Block, errorC <-chan error, validatorC <-chan bool) *RaftClient {
	//var flag int32
	//if cfg.Minerstart {
	//	flag = 1
	//}
	//c := drivers.NewBaseClient(cfg)
	client := &RaftClient{proposeC: proposeC, snapshotter: snapshotter, validatorC: validatorC, commitC: commitC, errorC: errorC}
	client.Cfg = cfg
	log.Info("Enter consensus raft")
	//c.SetChild(b)
	//TODO:多起一个readCommit，可能时raft包有bug
	//go client.readCommits(commitC, errorC)

	return client
}

func (client *RaftClient) getSnapshot() ([]byte, error) {
	//这里可能导致死锁
	client.mulock.Lock()
	defer client.mulock.Unlock()
	return proto.Marshal(client.currentBlock)
}
func (client *RaftClient) recoverFromSnapshot(snapshot []byte) error {
	var block types.Block
	if err := proto.Unmarshal(snapshot, &block); err != nil {
		return err
	}
	client.mulock.Lock()
	defer client.mulock.Unlock()
	client.currentBlock = &block
	//client.mulock.Unlock()
	return nil
}

//func (client *RaftClient) CreateGenesisTx() (ret []*types.Transaction) {
//	var tx types.Transaction
//	tx.Execer = []byte("coins")
//	tx.To = client.Cfg.Genesis
//	//gen payload
//	g := &types.CoinsAction_Genesis{}
//	g.Genesis = &types.CoinsGenesis{}
//	g.Genesis.Amount = 1e8 * types.Coin
//	tx.Payload = types.Encode(&types.CoinsAction{Value: g, Ty: types.CoinsActionGenesis})
//	ret = append(ret, &tx)
//	return
//}

//这个isMinning 在raft中用不到，默认是leader节点进行挖矿打包
//func (client *RaftClient) IsMining() bool {
//	return atomic.LoadInt32(&client.minerStart) == 1
//}

func (client *RaftClient) RequestBlock(start int64) (*types.Block, error) {
	if client.qclient == nil {
		panic("client not bind message queue.")
	}
	msg := client.qclient.NewMessage("blockchain", types.EventGetBlocks, &types.ReqBlocks{start, start, false})
	client.qclient.Send(msg, true)
	resp, err := client.qclient.Wait(msg)
	if err != nil {
		return nil, err
	}
	blocks := resp.GetData().(*types.BlockDetails)
	return blocks.Items[0].Block, nil
}

// 初始化时，取一次区块高度放在内存中，后面自增长，不用再重复去blockchain取
func (client *RaftClient) getInitHeight() int64 {

	msg := client.qclient.NewMessage("blockchain", types.EventGetBlockHeight, nil)

	client.qclient.Send(msg, true)
	replyHeight, err := client.qclient.Wait(msg)
	h := replyHeight.GetData().(*types.ReplyBlockHeight).Height
	rlog.Info("init = ", "height", h)
	if err != nil {
		panic("error happens when get height from blockchain")
	}
	return h
}

// 向blockchain写区块
func (client *RaftClient) writeBlock(prevHash []byte, block *types.Block) error {
	blockdetail, err := util.ExecBlock(client.q, prevHash, block, false)
	if err != nil { //never happen
		panic(err)
	}
	if len(blockdetail.Block.Txs) == 0 {
		return errors.New("ErrNoTxs")
	}
	msg := client.qclient.NewMessage("blockchain", types.EventAddBlockDetail, blockdetail)
	client.qclient.Send(msg, true)
	resp, err := client.qclient.Wait(msg)
	if err != nil {
		return err
	}
	if resp.GetData().(*types.Reply).IsOk {
		client.setCurrentBlock(block)
	} else {
		//TODO:
		//把txs写回mempool
		reply := resp.GetData().(*types.Reply)
		return errors.New(string(reply.GetMsg()))
	}
	return nil
}

func (client *RaftClient) setCurrentBlock(b *types.Block) {
	client.mulock.Lock()
	if client.currentBlock == nil || client.currentBlock.Height <= b.Height {
		client.currentBlock = b
	}
	client.mulock.Unlock()
}

func (client *RaftClient) getCurrentBlock() (b *types.Block) {
	client.mulock.Lock()
	defer client.mulock.Unlock()
	return client.currentBlock
}

func (client *RaftClient) getCurrentHeight() int64 {
	client.mulock.Lock()
	start := client.currentBlock.Height
	client.mulock.Unlock()
	return start
}

func (client *RaftClient) SetQueue(q *queue.Queue) {
	log.Info("Enter SetQueue method of consensus")
	// 只有主节点打包区块，其余节点接受p2p广播过来的区块
	//TODO:这里应该要改成轮询执行，当本台节点为Validator节点时则执行
	go client.readCommits(client.commitC, client.errorC)
	client.qclient = q.GetClient()
	client.q = q
	//client.initBlock()
	//TODO：当raft集群中的leader节点突然发生故障，此时另外的节点已经选举出新的leader，
	// 老的leader中运行的打包程此刻应该被终止？
	//isValidator = true
	//go client.eventLoop()
	//go client.createBlock()
	go client.pollingTask(q)
}

func (client *RaftClient) Close() {
	rlog.Info("consensus raft closed")
}

func (client *RaftClient) initBlock() {
	height := client.getInitHeight()
	if height == -1 {
		// 创世区块
		newblock := &types.Block{}
		newblock.Height = 0
		newblock.BlockTime = client.Cfg.GenesisBlockTime
		// TODO: 下面这些值在创世区块中赋值nil，是否合理？
		newblock.ParentHash = zeroHash[:]
		//tx := client.child.CreateGenesisTx()
		tx := client.CreateGenesisTx()
		newblock.Txs = tx
		newblock.TxHash = merkle.CalcMerkleRoot(newblock.Txs)
		// 通过propose channel把block传到raft核心
		client.propose(newblock)
		// 把区块放在内存中
		//TODO:这里要等确认后才能把当前的块设置为新块
		client.setCurrentBlock(newblock)
		//client.writeBlock(zeroHash[:], newblock)
	} else {
		block, err := client.RequestBlock(height)
		if err != nil {
			panic(err)
		}
		client.setCurrentBlock(block)
	}
}
func (client *RaftClient) checkTxDup(txs []*types.Transaction) (transactions []*types.Transaction) {
	var checkHashList types.TxHashList
	txMap := make(map[string]*types.Transaction)
	for _, tx := range txs {
		hash := tx.Hash()
		txMap[string(hash)] = tx
		checkHashList.Hashes = append(checkHashList.Hashes, hash)
	}
	// 发送Hash过后的交易列表给blockchain模块
	hashList := client.qclient.NewMessage("blockchain", types.EventTxHashList, &checkHashList)
	client.qclient.Send(hashList, true)
	dupTxList, _ := client.qclient.Wait(hashList)

	// 取出blockchain返回的重复交易列表
	dupTxs := dupTxList.GetData().(*types.TxHashList).Hashes

	for _, hash := range dupTxs {
		delete(txMap, string(hash))
	}

	for _, tx := range txMap {
		transactions = append(transactions, tx)
	}
	return transactions
}

func (client *RaftClient) createBlock() {
	issleep := true
	for {
		//如果leader节点突然挂了，不是打包节点，需要退出
		if !isValidator {
			log.Warn("I'm not the validator node anymore,exit.=============================")
			break
		}
		log.Info("==================start create new block!=====================")
		if issleep {
			time.Sleep(time.Second)
		}
		txs := client.RequestTx()
		if len(txs) == 0 {
			issleep = true
			continue
		}
		issleep = false
		//check dup
		txs = client.checkTxDup(txs)
		fmt.Println(len(txs))
		lastBlock := client.getCurrentBlock()
		rlog.Warn("**************************************lastBlockHeight:", lastBlock.Height)
		fmt.Println(lastBlock.Height)
		var newblock types.Block
		newblock.ParentHash = lastBlock.Hash()
		newblock.Height = lastBlock.Height + 1
		newblock.Txs = txs
		newblock.TxHash = merkle.CalcMerkleRoot(newblock.Txs)
		newblock.BlockTime = time.Now().Unix()
		if lastBlock.BlockTime >= newblock.BlockTime {
			newblock.BlockTime = lastBlock.BlockTime + 1
		}
		rlog.Info("Send block to raft core")
		client.propose(&newblock)
		err := client.writeBlock(lastBlock.StateHash, &newblock)
		if err != nil {
			issleep = true
			continue
		}
	}
}

// 准备新区块
//func (client *RaftClient) EventLoop() {
//	// 监听blockchain模块，获取当前最高区块
//	client.qclient.Sub("consensus")
//	go func() {
//		for msg := range client.qclient.Recv() {
//			rlog.Info("consensus recv", "msg", msg)
//			if msg.Ty == types.EventAddBlock {
//				block := msg.GetData().(*types.BlockDetail).Block
//				client.SetCurrentBlock(block)
//			} else if msg.Ty == types.EventCheckBlock {
//				block := msg.GetData().(*types.BlockDetail)
//				err := client.CheckBlock(block)
//				msg.ReplyErr("EventCheckBlock", err)
//			} else if msg.Ty == types.EventMinerStart {
//				if !atomic.CompareAndSwapInt32(&client.minerStart, 0, 1) {
//					msg.ReplyErr("EventMinerStart", types.ErrMinerIsStared)
//				} else {
//					client.once.Do(func() {
//						client.InitBlock()
//					})
//					msg.ReplyErr("EventMinerStart", nil)
//				}
//			} else if msg.Ty == types.EventMinerStop {
//				if !atomic.CompareAndSwapInt32(&client.minerStart, 1, 0) {
//					msg.ReplyErr("EventMinerStop", types.ErrMinerNotStared)
//				} else {
//					msg.ReplyErr("EventMinerStop", nil)
//				}
//			}
//		}
//	}()
//}
// 准备新区块
func (client *RaftClient) eventLoop() {
	// 监听blockchain模块，获取当前最高区块
	client.qclient.Sub("consensus")
	go func() {
		for msg := range client.qclient.Recv() {
			//当本节点不再时leader节点时，需要退出。
			if !isValidator {
				log.Warn("I'm not the validator node anymore,exit.=============================")
				break
			}
			rlog.Info("consensus recv", "msg", msg)
			if msg.Ty == types.EventAddBlock {
				block := msg.GetData().(*types.BlockDetail).Block
				client.setCurrentBlock(block)
			}
			//} else if msg.Ty == types.EventCheckBlock {
			//	block := msg.GetData().(*types.BlockDetail)
			//	err := client.CheckBlock(block)
			//	msg.ReplyErr("EventCheckBlock", err)
			//}
		}
	}()
}
func (client *RaftClient) CheckBlock(block *types.BlockDetail) error {
	//check parent
	if block.Block.Height == 0 { //genesis block not check
		return nil
	}
	parent, err := client.RequestBlock(block.Block.Height - 1)
	if err != nil {
		return err
	}
	//check base info
	if parent.Height+1 != block.Block.Height {
		return types.ErrBlockHeight
	}
	//check by drivers
	//err = client.child.CheckBlock(parent, block)

	return nil
}

// Mempool中取交易列表
func (client *RaftClient) RequestTx() []*types.Transaction {
	if client.qclient == nil {
		panic("client not bind message queue.")
	}
	msg := client.qclient.NewMessage("mempool", types.EventTxList, listSize)
	client.qclient.Send(msg, true)
	resp, _ := client.qclient.Wait(msg)
	return resp.GetData().(*types.ReplyTxList).GetTxs()
}
func (client *RaftClient) CreateGenesisTx() (ret []*types.Transaction) {
	var tx types.Transaction
	tx.Execer = []byte("coins")
	tx.To = client.Cfg.Genesis
	//gen payload
	g := &types.CoinsAction_Genesis{}
	g.Genesis = &types.CoinsGenesis{}
	g.Genesis.Amount = 1e8 * types.Coin
	tx.Payload = types.Encode(&types.CoinsAction{Value: g, Ty: types.CoinsActionGenesis})
	ret = append(ret, &tx)
	return
}

// 向raft底层发送block
func (client *RaftClient) propose(block *types.Block) {
	client.proposeC <- block
}

// 从receive channel中读leader发来的block
func (client *RaftClient) readCommits(commitC <-chan *types.Block, errorC <-chan error) {
	//var prevHash []byte
	for {
		select {
		case data := <-commitC:
			rlog.Info("Get block from commit channel")
			if data == nil {
				rlog.Info("data is nil===================================")
				snapshot, err := client.snapshotter.Load()
				if err == snap.ErrNoSnapshot {
					return
				}

				log.Info("loading snapshot at term %d and index %d", snapshot.Metadata.Term, snapshot.Metadata.Index)
				if err := client.recoverFromSnapshot(snapshot.Data); err != nil {
					panic(err)
				}
				continue
			}
			//TODO:这里有个极端的情况，就是readCommits 比validatorC早接收到消息，这样的话，
			// 在程序刚开始启动的时候有可能存在丢失数据的问题
			//if !isValidator {
			//	log.Warn("I'm not the validator node, I don't need to writeBlock!==========================")
			//	continue
			//}
			log.Warn("I'm the validator node, I need to writeBlock!==========================")
			client.setCurrentBlock(data)
			//lastBlock := client.getCurrentBlock()
			//if lastBlock == nil {
			//	prevHash = zeroHash[:]
			//} else {
			//	prevHash = lastBlock.StateHash
			//}
			//client.mulock.Lock()
			//err:=client.writeBlock(prevHash, data)
			//client.mulock.Unlock()
			//if err !=nil {
			//log.Error("=========writeBlock  failed!=======")
			//continue
			//}
			log.Info("============writeBlock successfully!======")

		case err, ok := <-errorC:
			log.Info("============error======")
			if ok {
				panic(err)
			}

		}
	}
}

//轮询任务，去检测本机器是否为validator节点，如果是，则执行打包任务
func (client *RaftClient) pollingTask(q *queue.Queue) {

	for {
		select {
		case value := <-client.validatorC:
			if !value {
				log.Warn("I'm not the validator node!=====")
				isValidator = false
			} else if !isValidator && value {
				log.Info("==================start init block========================")
				client.initBlock()
				//TODO：当raft集群中的leader节点突然发生故障，此时另外的节点已经选举出新的leader，
				// 老的leader中运行的打包程此刻应该被终止？
				isValidator = true
				go client.eventLoop()
				go client.createBlock()

			}

		}
	}
}
