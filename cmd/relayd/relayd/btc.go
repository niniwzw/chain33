package relayd

import (
	"errors"
	"sync"
	"time"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/valyala/fasthttp"
	"gitlab.33.cn/chain33/chain33/common/merkle"
	"gitlab.33.cn/chain33/chain33/types"
)

type (
	BlockStamp struct {
		Height int32
		Hash   chainhash.Hash
	}

	BlockMeta struct {
		BlockStamp
		Time time.Time
	}

	ClientConnected   struct{}
	BlockConnected    BlockMeta
	BlockDisconnected BlockMeta
)

type Params struct {
	*chaincfg.Params
	RPCClientPort string
	RPCServerPort string
}

var MainNetParams = Params{
	Params:        &chaincfg.MainNetParams,
	RPCClientPort: "8334",
	RPCServerPort: "8332",
}

type BTCClient struct {
	*rpcclient.Client
	httpClient          *fasthttp.Client
	connConfig          *rpcclient.ConnConfig
	chainParams         *chaincfg.Params
	reconnectAttempts   int
	enqueueNotification chan interface{}
	dequeueNotification chan interface{}
	currentBlock        chan *BlockStamp
	quit                chan struct{}
	wg                  sync.WaitGroup
	started             bool
	quitMtx             sync.Mutex
}

func NewBTCClient(config *rpcclient.ConnConfig, reconnectAttempts int) (*BTCClient, error) {
	if reconnectAttempts < 0 {
		return nil, errors.New("ReconnectAttempts must be positive")
	}
	client := &BTCClient{
		connConfig:          config,
		chainParams:         MainNetParams.Params,
		reconnectAttempts:   reconnectAttempts,
		enqueueNotification: make(chan interface{}),
		dequeueNotification: make(chan interface{}),
		currentBlock:        make(chan *BlockStamp),
		quit:                make(chan struct{}),
	}
	ntfnCallbacks := &rpcclient.NotificationHandlers{
		OnClientConnected:   client.onClientConnect,
		OnBlockConnected:    client.onBlockConnected,
		OnBlockDisconnected: client.onBlockDisconnected,
	}
	rpcClient, err := rpcclient.New(client.connConfig, ntfnCallbacks)
	if err != nil {
		return nil, err
	}
	client.Client = rpcClient
	return client, nil
}

func (c *BTCClient) Start() error {
	err := c.Connect(c.reconnectAttempts)
	if err != nil {
		return err
	}

	// Verify that the server is running on the expected network.
	net, err := c.GetCurrentNet()
	if err != nil {
		c.Disconnect()
		return err
	}
	if net != c.chainParams.Net {
		c.Disconnect()
		return errors.New("mismatched networks")
	}

	c.quitMtx.Lock()
	c.started = true
	c.quitMtx.Unlock()

	c.wg.Add(1)
	go c.handler()
	return nil
}

func (c *BTCClient) Stop() {
	c.quitMtx.Lock()
	select {
	case <-c.quit:
	default:
		close(c.quit)
		c.Client.Shutdown()

		if !c.started {
			close(c.dequeueNotification)
		}
	}
	c.quitMtx.Unlock()
}

func (c *BTCClient) WaitForShutdown() {
	c.Client.WaitForShutdown()
	c.wg.Wait()
}

func (c *BTCClient) Notifications() <-chan interface{} {
	return c.dequeueNotification
}

func (c *BTCClient) BlockStamp() (*BlockStamp, error) {
	select {
	case bs := <-c.currentBlock:
		return bs, nil
	case <-c.quit:
		return nil, errors.New("disconnected")
	}
}

func parseBlock(block *btcjson.BlockDetails) (*BlockMeta, error) {
	if block == nil {
		return nil, nil
	}
	blkHash, err := chainhash.NewHashFromStr(block.Hash)
	if err != nil {
		return nil, err
	}
	blk := &BlockMeta{
		BlockStamp: BlockStamp{
			Height: block.Height,
			Hash:   *blkHash,
		},
		Time: time.Unix(block.Time, 0),
	}
	return blk, nil
}

func (c *BTCClient) onClientConnect() {
	select {
	case c.enqueueNotification <- ClientConnected{}:
	case <-c.quit:
	}
}

func (c *BTCClient) onBlockConnected(hash *chainhash.Hash, height int32, time time.Time) {
	select {
	case c.enqueueNotification <- BlockConnected{
		BlockStamp: BlockStamp{
			Hash:   *hash,
			Height: height,
		},
		Time: time,
	}:
	case <-c.quit:
	}
}

func (c *BTCClient) onBlockDisconnected(hash *chainhash.Hash, height int32, time time.Time) {
	select {
	case c.enqueueNotification <- BlockDisconnected{
		BlockStamp: BlockStamp{
			Hash:   *hash,
			Height: height,
		},
		Time: time,
	}:
	case <-c.quit:
	}
}

func (c *BTCClient) handler() {
	hash, height, err := c.GetBestBlock()
	if err != nil {
		c.Stop()
		c.wg.Done()
		return
	}

	bs := &BlockStamp{Hash: *hash, Height: height}
	var notifications []interface{}
	enqueue := c.enqueueNotification
	var dequeue chan interface{}
	var next interface{}
	pingChan := time.After(time.Minute)
out:
	for {
		select {
		case n, ok := <-enqueue:
			if !ok {
				// If no notifications are queued for handling,
				// the queue is finished.
				if len(notifications) == 0 {
					break out
				}
				// nil channel so no more reads can occur.
				enqueue = nil
				continue
			}
			if len(notifications) == 0 {
				next = n
				dequeue = c.dequeueNotification
			}
			notifications = append(notifications, n)
			pingChan = time.After(time.Minute)

		case dequeue <- next:
			if n, ok := next.(BlockConnected); ok {
				bs = &BlockStamp{
					Height: n.Height,
					Hash:   n.Hash,
				}
			}

			notifications[0] = nil
			notifications = notifications[1:]
			if len(notifications) != 0 {
				next = notifications[0]
			} else {
				// If no more notifications can be enqueued, the
				// queue is finished.
				if enqueue == nil {
					break out
				}
				dequeue = nil
			}

		case <-pingChan:
			type sessionResult struct {
				err error
			}
			sessionResponse := make(chan sessionResult, 1)
			go func() {
				_, err := c.Session()
				sessionResponse <- sessionResult{err}
			}()

			select {
			case resp := <-sessionResponse:
				if resp.err != nil {
					// log.Errorf("Failed to receive session "+"result: %v", resp.err)
					c.Stop()
					break out
				}
				pingChan = time.After(time.Minute)

			case <-time.After(time.Minute):
				// log.Errorf("Timeout waiting for session RPC")
				c.Stop()
				break out
			}

		case c.currentBlock <- bs:

		case <-c.quit:
			break out
		}
	}

	c.Stop()
	close(c.dequeueNotification)
	c.wg.Done()
}

// POSTClient creates the equivalent HTTP POST rpcclient.Client.
func (c *BTCClient) POSTClient() (*rpcclient.Client, error) {
	configCopy := *c.connConfig
	configCopy.HTTPPostMode = true
	return rpcclient.New(&configCopy, nil)
}

func (b *BTCClient) GetSPV(height uint64, txHash string) (*types.BtcSpv, error) {
	hash, err := chainhash.NewHashFromStr(txHash)
	if err != nil {
		return nil, err
	}

	ret, err := b.GetRawTransactionVerbose(hash)
	if err != nil {
		return nil, err
	}

	blockHash, err := chainhash.NewHashFromStr(ret.BlockHash)
	if err != nil {
		return nil, err
	}

	block, err := b.GetBlockVerbose(blockHash)
	if err != nil {
		return nil, err
	}
	var txIndex uint32
	txs := make([][]byte, 0, len(block.Tx))
	for index, tx := range block.Tx {
		if txHash == tx {
			txIndex = uint32(index)
		}
		hash, err := merkle.NewHashFromStr(tx)
		if err != nil {
			return nil, err
		}
		txs = append(txs, hash.CloneBytes())
	}
	proof := merkle.GetMerkleBranch(txs, txIndex)
	spv := &types.BtcSpv{
		Hash:        txHash,
		Time:        block.Time,
		Height:      uint64(block.Height),
		BlockHash:   block.Hash,
		TxIndex:     uint64(txIndex),
		BranchProof: proof,
	}
	return spv, nil
}

func (b *BTCClient) GetTransaction(hash string) (*types.BtcTransaction, error) {
	txHash, err := chainhash.NewHashFromStr(hash)
	if err != nil {
		return nil, err
	}

	tx, err := b.GetRawTransactionVerbose(txHash)
	if err != nil {
		return nil, err
	}
	btxTx := &types.BtcTransaction{}
	btxTx.Hash = hash
	btxTx.Time = tx.Time
	// TODO
	// btxTx.BlockHeight = tx.BlockHash
	// vin := make([]*types.Vin, len(tx.Vin))
	// for index, in := range tx.Vin {
	// 	vin[index].Value = in.Vout
	// 	vin[index].Address = in.Address
	// }
	// btcTx.Vin = vin

	vout := make([]*types.Vout, len(tx.Vout))
	for index, in := range tx.Vout {
		vout[index].Value = uint64(in.Value)
		vout[index].Address = in.ScriptPubKey.Addresses[0]
	}
	btxTx.Vout = vout

	return btxTx, nil
}

func (b *BTCClient) GetBlockHeader(height uint64) (*types.BtcHeader, error) {
	hash, err := b.GetBlockHash(int64(height))
	if err != nil {
		return nil, err
	}
	header, err := b.GetBlockHeaderVerbose(hash)
	if err != nil {
		return nil, err
	}

	h := &types.BtcHeader{
		Hash:          header.Hash,
		Confirmations: header.Confirmations,
		Height:        uint64(header.Height),
		MerkleRoot:    header.MerkleRoot,
		Time:          header.Time,
		Nonce:         int64(header.Nonce),
		// TODO
		// Bits:header.Bits,
		// Difficulty:header.Difficulty,
		PreviousHash: header.PreviousHash,
		NextHash:     header.NextHash,
	}
	return h, nil

}
