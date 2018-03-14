package mavl

//store package store the world - state data
import (
	"code.aliyun.com/chain33/chain33/common"
	"code.aliyun.com/chain33/chain33/common/mavl"
	"code.aliyun.com/chain33/chain33/queue"
	"code.aliyun.com/chain33/chain33/store/drivers"
	"code.aliyun.com/chain33/chain33/types"
	lru "github.com/hashicorp/golang-lru"
	log "github.com/inconshreveable/log15"
)

var mlog = log.New("module", "mavl")

func SetLogLevel(level string) {
	common.SetLogLevel(level)
}

func DisableLog() {
	mlog.SetHandler(log.DiscardHandler())
}

type MavlStore struct {
	*drivers.BaseStore
	trees map[string]*mavl.MAVLTree
	cache *lru.Cache
}

//driver
//dbpath
func New(cfg *types.Store) *MavlStore {
	bs := drivers.NewBaseStore(cfg)
	mavls := &MavlStore{bs, make(map[string]*mavl.MAVLTree), nil}
	mavls.cache, _ = lru.New(10)
	bs.SetChild(mavls)
	return mavls
}

func (mavls *MavlStore) Close() {
	mavls.BaseStore.Close()
	mlog.Info("store mavl closed")
}

func (mavls *MavlStore) Set(datas *types.StoreSet) []byte {
	hash := mavl.SetKVPair(mavls.GetDB(), datas)
	return hash
}

func (mavls *MavlStore) Get(datas *types.StoreGet) [][]byte {
	var tree *mavl.MAVLTree
	var err error
	values := make([][]byte, len(datas.Keys))
	if data, ok := mavls.cache.Get(string(datas.StateHash)); ok {
		tree = data.(*mavl.MAVLTree)
	} else {
		tree = mavl.NewMAVLTree(mavls.GetDB())
		err = tree.Load(datas.StateHash)
		mavls.cache.Add(string(datas.StateHash), tree)
	}
	if err == nil {
		for i := 0; i < len(datas.Keys); i++ {
			_, value, exit := tree.Get(datas.Keys[i])
			if exit {
				values[i] = value
			}
		}
	}
	return values
}

func (mavls *MavlStore) MemSet(datas *types.StoreSet) []byte {
	tree := mavl.NewMAVLTree(mavls.GetDB())
	tree.Load(datas.StateHash)
	for i := 0; i < len(datas.KV); i++ {
		tree.Set(datas.KV[i].Key, datas.KV[i].Value)
	}
	hash := tree.Hash()
	mavls.trees[string(hash)] = tree
	return hash
}

func (mavls *MavlStore) Commit(req *types.ReqHash) []byte {
	tree, ok := mavls.trees[string(req.Hash)]
	if !ok {
		mlog.Error("store mavl commit", "err", types.ErrHashNotFound)
		return nil
	}
	tree.Save()
	delete(mavls.trees, string(req.Hash))
	return req.Hash
}

func (mavls *MavlStore) Rollback(req *types.ReqHash) []byte {
	_, ok := mavls.trees[string(req.Hash)]
	if !ok {
		mlog.Error("store mavl rollback", "err", types.ErrHashNotFound)
		return nil
	}
	delete(mavls.trees, string(req.Hash))
	return req.Hash
}

func (mavls *MavlStore) ProcEvent(msg queue.Message) {
	qclient := mavls.GetQueueClient()
	if msg.Ty == types.EventStoreRollback {
		req := msg.GetData().(*types.ReqHash)
		hash := mavls.Rollback(req)
		if hash == nil {
			msg.Reply(qclient.NewMessage("", types.EventStoreRollback, types.ErrHashNotFound))
		} else {
			msg.Reply(qclient.NewMessage("", types.EventStoreRollback, &types.ReplyHash{hash}))
		}
	} else {
		msg.ReplyErr("MavlStore", types.ErrActionNotSupport)
	}
}
