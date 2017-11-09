package mempool

import (
	"container/list"
	"sync"

	"code.aliyun.com/chain33/chain33/queue"
	"code.aliyun.com/chain33/chain33/types"
	log "github.com/inconshreveable/log15"
)

const poolCacheSize = 300000 // mempool容量

var mlog = log.New("module", "mempool")

func SetLogLevel(level int) {

}

func DisableLog() {
	mlog.SetHandler(log.DiscardHandler())
}

type MClient interface {
	SetQueue(q *queue.Queue)
	SendTx(tx *types.Transaction) queue.Message
}

type channelClient struct {
	qclient queue.IClient
}

type Mempool struct {
	proxyMtx sync.Mutex
	cache    *txCache
}

type txCache struct {
	// mtx  sync.Mutex
	size   int
	txMap  map[string]*list.Element
	txList *list.List
}

// NewTxCache初始化txCache
func newTxCache(cacheSize int) *txCache {
	return &txCache{
		size:   cacheSize,
		txMap:  make(map[string]*list.Element, cacheSize),
		txList: list.New(),
	}
}

// txCache.Exists判断txCache中是否存在给定tx
func (cache *txCache) Exists(tx *types.Transaction) bool {
	_, exists := cache.txMap[string(tx.Hash())]
	return exists
}

// txCache.Push把给定tx添加到txCache并返回true；如果tx已经存在txCache中则返回false
func (cache *txCache) Push(tx *types.Transaction) bool {
	if _, exists := cache.txMap[string(tx.Hash())]; exists {
		return false
	}

	if cache.txList.Len() >= cache.size {
		popped := cache.txList.Front()
		poppedTx := popped.Value.(*types.Transaction)
		delete(cache.txMap, string(poppedTx.Hash()))
		cache.txList.Remove(popped)
	}

	txElement := cache.txList.PushBack(tx)
	cache.txMap[string(tx.Hash())] = txElement

	return true
}

// txCache.Remove移除txCache中给定tx
func (cache *txCache) Remove(tx *types.Transaction) {
	cache.txList.Remove(cache.txMap[string(tx.Hash())])
	delete(cache.txMap, string(tx.Hash()))
}

// txCache.Size返回txCache中已存tx数目
func (cache *txCache) Size() int {
	return cache.txList.Len()
}

// txCache.SetMempoolSize用来设置Mempool容量
func (cache *txCache) SetSize(newSize int) {
	if cache.txList.Len() > 0 {
		panic("only can set a empty size")
	}
	cache.size = newSize
}

func New() *Mempool {
	pool := &Mempool{}
	pool.cache = newTxCache(poolCacheSize)
	return pool
}

func (mem *Mempool) Resize(size int) {
	mem.cache.SetSize(size)
}

// Mempool.GetTxList从txCache中返回给定数目的tx并从txCache中删除返回的tx
func (mem *Mempool) GetTxList(txListSize int) []*types.Transaction {
	mem.proxyMtx.Lock()
	defer mem.proxyMtx.Unlock()

	txsSize := mem.cache.Size()
	var result []*types.Transaction
	var i int
	var minSize int

	if txsSize <= txListSize {
		minSize = txsSize
	} else {
		minSize = txListSize
	}

	for i = 0; i < minSize; i++ {
		popped := mem.cache.txList.Front()
		poppedTx := popped.Value.(*types.Transaction)
		result = append(result, poppedTx)
		mem.cache.txList.Remove(popped)
		delete(mem.cache.txMap, string(poppedTx.Hash()))
	}

	return result
}

// Mempool.Size返回Mempool中txCache大小
func (mem *Mempool) Size() int {
	mem.proxyMtx.Lock()
	defer mem.proxyMtx.Unlock()
	return mem.cache.Size()
}

// Mempool.CheckTx坚持tx有效性并加入Mempool中
func (mem *Mempool) CheckTx(tx *types.Transaction) bool {
	mem.proxyMtx.Lock()
	defer mem.proxyMtx.Unlock()

	if mem.cache.Exists(tx) {
		return false
	}
	mem.cache.Push(tx)

	return true
}

//Mempool.RemoveTxsOfBlock移除Mempool中已被Blockchain打包的tx
func (mem *Mempool) RemoveTxsOfBlock(block *types.Block) bool {
	mem.proxyMtx.Lock()
	defer mem.proxyMtx.Unlock()
	for tx := range block.Txs {
		exist := mem.cache.Exists(block.Txs[tx])
		if exist {
			mem.cache.Remove(block.Txs[tx])
		}
	}
	return true
}

// channelClient.SendTx向"p2p"发送消息
func (client *channelClient) SendTx(tx *types.Transaction) queue.Message {
	if client.qclient == nil {
		panic("client not bind message queue.")
	}

	msg := client.qclient.NewMessage("p2p", types.EventTxBroadcast, tx)
	client.qclient.Send(msg, true)
	resp, err := client.qclient.Wait(msg)

	if err != nil {
		resp.Data = err
	}

	return resp
}

func (mem *Mempool) SetQueue(q *queue.Queue) {
	client := q.GetClient()
	client.Sub("mempool")
	go func() {
		for msg := range client.Recv() {
			mlog.Info("mempool recv", "msg", msg)
			if msg.Ty == types.EventTx {
				if mem.CheckTx(msg.GetData().(*types.Transaction)) {
					msg.Reply(client.NewMessage("rpc", types.EventReply,
						&types.Reply{true, nil}))
				} else {
					msg.Reply(client.NewMessage("rpc", types.EventReply,
						&types.Reply{false, []byte("transaction exists")}))
				}
			} else if msg.Ty == types.EventTxAddMempool {
				if mem.CheckTx(msg.GetData().(*types.Transaction)) {
					msg.Reply(client.NewMessage("rpc", types.EventReply,
						&types.Reply{true, nil}))
				} else {
					msg.Reply(client.NewMessage("rpc", types.EventReply,
						&types.Reply{false, []byte("transaction exists")}))
				}
			} else if msg.Ty == types.EventTxList {
				msg.Reply(client.NewMessage("consensus", types.EventTxListReply,
					&types.ReplyTxList{mem.GetTxList(10000)}))
			} else if msg.Ty == types.EventAddBlock {
				mem.RemoveTxsOfBlock(msg.GetData().(*types.Block))
			} else if msg.Ty == types.EventGetMempoolSize {
				msg.Reply(client.NewMessage("rpc", types.EventMempoolSize,
					&types.MempoolSize{int64(mem.Size())}))
			}
		}
	}()
}
