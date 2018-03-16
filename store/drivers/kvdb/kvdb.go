package kvdb

import (
	"bytes"

	"code.aliyun.com/chain33/chain33/common"
	"code.aliyun.com/chain33/chain33/queue"
	"code.aliyun.com/chain33/chain33/store/drivers"
	"code.aliyun.com/chain33/chain33/types"
	log "github.com/inconshreveable/log15"
)

var klog = log.New("module", "kvdb")

func SetLogLevel(level string) {
	common.SetLogLevel(level)
}

func DisableLog() {
	klog.SetHandler(log.DiscardHandler())
}

type KVStore struct {
	*drivers.BaseStore
	cache map[string]map[string][]byte
}

func New(cfg *types.Store) *KVStore {
	bs := drivers.NewBaseStore(cfg)
	kvs := &KVStore{bs, make(map[string]map[string][]byte)}
	bs.SetChild(kvs)
	return kvs
}

func (kvs *KVStore) Close() {
	kvs.BaseStore.Close()
	klog.Info("store kvdb closed")
}

func (kvs *KVStore) Set(datas *types.StoreSet) []byte {
	hash := calcHash(datas)
	kvmap := make(map[string][]byte)
	for _, kv := range datas.KV {
		kvmap[string(kv.Key)] = kv.Value
	}
	kvs.save(kvmap)
	return hash
}

func (kvs *KVStore) Get(datas *types.StoreGet) [][]byte {
	values := make([][]byte, len(datas.Keys))
	if kvmap, ok := kvs.cache[string(datas.StateHash)]; ok {
		for i := 0; i < len(datas.Keys); i++ {
			value := kvmap[string(datas.Keys[i])]
			if value != nil {
				values[i] = value
			}
		}
	} else {
		db := kvs.GetDB()
		for i := 0; i < len(datas.Keys); i++ {
			value := db.Get(datas.Keys[i])
			if value != nil {
				values[i] = value
			}
		}
	}
	return values
}

func (kvs *KVStore) MemSet(datas *types.StoreSet) []byte {
	hash := calcHash(datas)
	kvmap := make(map[string][]byte)
	for _, kv := range datas.KV {
		kvmap[string(kv.Key)] = kv.Value
	}
	kvs.cache[string(hash)] = kvmap
	if len(kvs.cache) > 100 {
		klog.Error("too many items in cache")
	}
	return hash
}

func (kvs *KVStore) Commit(req *types.ReqHash) []byte {
	kvmap, ok := kvs.cache[string(req.Hash)]
	if !ok {
		klog.Error("store kvdb commit", "err", types.ErrHashNotFound)
		return nil
	}
	kvs.save(kvmap)
	delete(kvs.cache, string(req.Hash))
	return req.Hash
}

func (kvs *KVStore) Rollback(req *types.ReqHash) []byte {
	_, ok := kvs.cache[string(req.Hash)]
	if !ok {
		klog.Error("store kvdb rollback", "err", types.ErrHashNotFound)
		return nil
	}
	delete(kvs.cache, string(req.Hash))
	return req.Hash
}

func (kvs *KVStore) ProcEvent(msg queue.Message) {
	qclient := kvs.GetQueueClient()
	if msg.Ty == types.EventStoreRollback {
		req := msg.GetData().(*types.ReqHash)
		hash := kvs.Rollback(req)
		if hash == nil {
			msg.Reply(qclient.NewMessage("", types.EventStoreRollback, types.ErrHashNotFound))
		} else {
			msg.Reply(qclient.NewMessage("", types.EventStoreRollback, &types.ReplyHash{hash}))
		}
	} else {
		msg.ReplyErr("KVStore", types.ErrActionNotSupport)
	}
}

func (kvs *KVStore) save(kvmap map[string][]byte) {
	storeBatch := kvs.GetDB().NewBatch(true)
	for key, value := range kvmap {
		if value == nil {
			storeBatch.Delete([]byte(key))
		} else {
			storeBatch.Set([]byte(key), value)
		}
	}
	storeBatch.Write()
}

func calcHash(datas *types.StoreSet) []byte {
	var hashes [][]byte
	hashes = append(hashes, datas.StateHash)
	kvset := datas.KV
	for _, kv := range kvset {
		hashes = append(hashes, kv.Key, kv.Value)
	}
	data := bytes.Join(hashes, []byte(""))
	return common.Sha256(data)
}
