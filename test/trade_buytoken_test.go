package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"testing"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"gitlab.33.cn/chain33/chain33/common"
	"gitlab.33.cn/chain33/chain33/common/address"
	"gitlab.33.cn/chain33/chain33/types"
)

var (
	tokenName = "FZMKJ"
	privKey   = "0xc4553726d9d04b6a58ff99ba4b4aeb47055f97f04514d30535cde686365c2af2"
)

func TestBuytoken(t *testing.T) {
	// 构造交易1
	var tx *types.Transaction
	transfer := &types.CoinsAction{}
	toAddr := address.ExecAddress("trade")
	v := &types.CoinsAction_Transfer{Transfer: &types.CoinsTransfer{Amount: 10 * types.Coin, Note: "trans to trade", To: toAddr}}
	transfer.Value = v
	transfer.Ty = types.CoinsActionTransfer
	tx = &types.Transaction{Execer: []byte("coins"), Payload: types.Encode(transfer), To: toAddr}
	tx.Fee, _ = tx.GetRealFee(types.MinFee)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	tx.Nonce = random.Int63()
	txHex1 := common.ToHex(types.Encode(tx))[2:]
	fmt.Println("tx1:", txHex1)

	// 查询sellID
	var queryMap = make(map[string]interface{})
	queryMap["method"] = "Chain33.Query"
	queryMap["jsonrpc"] = "2.0"
	queryMap["id"] = 2
	var paramsMap = make([]interface{}, 0)
	var param = make(map[string]interface{})
	param["execer"] = "trade"
	param["funcName"] = "GetTokenSellOrderByStatus"
	var payloadMap = make(map[string]interface{})
	payloadMap["tokenSymbol"] = tokenName
	payloadMap["count"] = 1
	payloadMap["direction"] = 1
	payloadMap["status"] = types.TradeOrderStatusOnSale
	param["payload"] = payloadMap
	paramsMap = append(paramsMap, param)
	queryMap["params"] = paramsMap
	queryByte, err := json.Marshal(queryMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := http.Post("http://localhost:8801", "application/json", bytes.NewBufferString(string(queryByte)))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	res, _ := simplejson.NewJson(b)
	queryRes := res.Get("result").Get("sellOrders").MustArray()
	if len(queryRes) == 0 {
		fmt.Println("no sell item")
		return
	}
	sellId, ok := queryRes[0].(map[string]interface{})["sellID"]
	if !ok {
		return
	}
	fmt.Println("sellID:", sellId)

	// 构造交易2
	poststr := fmt.Sprintf(`{"jsonrpc":"2.0","id":2,"method":"Chain33.CreateRawTradeBuyTx","params":[{"boardlotCnt":1,"sellID":"%v"}]}`, sellId)
	resp, err = http.Post("http://localhost:8801", "application/json", bytes.NewBufferString(poststr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	res, _ = simplejson.NewJson(b)
	txHex2 := res.Get("result").MustString()
	fmt.Println("tx2:", txHex2)

	// 构造交易组
	txsArr := []string{txHex1, txHex2}
	var transactions []*types.Transaction
	for _, t := range txsArr {
		txByte, err := hex.DecodeString(t)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		var transaction types.Transaction
		types.Decode(txByte, &transaction)
		transactions = append(transactions, &transaction)
	}
	group, err := types.CreateTxGroup(transactions)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	err = group.Check(0, types.MinFee)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	newtx := group.Tx()
	grouptx := hex.EncodeToString(types.Encode(newtx))
	fmt.Println("txGroup:", grouptx)

	// 签名交易组
	poststr = fmt.Sprintf(`{"jsonrpc":"2.0","id":2,"method":"Chain33.SignRawTx","params":[{"privkey":"%v","txHex":"%v","expire":"%v","index":%v}]}`, privKey, grouptx, "2h", 0)
	resp, err = http.Post("http://localhost:8801", "application/json", bytes.NewBufferString(poststr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	res, _ = simplejson.NewJson(b)
	signedTx := res.Get("result").MustString()
	fmt.Println("signedTx:", signedTx)

	// 发送交易组
	poststr = fmt.Sprintf(`{"jsonrpc":"2.0","id":2,"method":"Chain33.SendTransaction","params":[{"data":"%v","token":"%v"}]}`, signedTx, types.BTY)
	resp, err = http.Post("http://localhost:8801", "application/json", bytes.NewBufferString(poststr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
