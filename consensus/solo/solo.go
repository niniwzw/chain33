package solo

import (
	"fmt"
	"log"
)
import "code.aliyun.com/chain33/chain33/queue"
import "code.aliyun.com/chain33/chain33/types"

type SoloClient struct {
	qclient queue.IClient
}

func NewSolo() {
	log.Println("consensus/solo")
}

func (client *SoloClient) SetQueue(q *queue.Queue) {

	// TODO: solo模式下判断当前节点是否主节点，主节点打包区块，其余节点不用做
	client.qclient = q.GetClient()
	// Get transaction list size
	listSize = 10000

	go func() {
		for {
			// 循环从mempool里取交易列表
			resp, err := RequestTx(10000)
			if err != nil {
				log.Fatal("error")
			}

			// TODO: 需要评估取数据的时间和下面处理数据的时间
			// TODO:剔除账本中已经存在的重复交易
			QueryTransaction(resp.Data)

			// TODO:对交易列表中交易进行排序
			// sortTransaction(txlist)

			// 组成新区块，并广播
		}
	}()
}

// 向mempool发起RequestTxList请求消息，返回对应的交易列表
func (client *SoloClient) RequestTx(txNum int64) (queue.Message, err) {
	if client.qclient == nil {
		panic("client not bind message queue.")
	}
	msg := client.qclient.NewMessage("mempool", types.RequestTxList, 0, txNum)
	client.qclient.Send(msg, true)
	return client.qclient.Wait(msg.Id)
}

func (client *SoloClient) GetPreBlock() (int, err) {
	//	msg := client.qclient.NewMessage("blockchain", types., 0, )
	//	client.qclient.Send(msg, true)
	//	resp, err := client.qclient.Wait(msg.Id)
}

func (client *SoloClient) QueryTransaction(txs []*types.Transaction) (proof *types.MerkleProof, err error) {

}

func (client *SoloClient) sortTransaction() {

}

func (client *SoloClient) BroadcastBlock() {
	msg := client.qclient.NewMessage("p2p", types.EventGetBlocks, 0, &types.RequestBlocks{start, end})
	client.qclient.Send(msg, true)
	resp, err := client.qclient.Wait(msg.Id)
}
