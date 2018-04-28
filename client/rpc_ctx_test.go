package client_test

import (
	"encoding/json"
	"fmt"
	"os"

	jsonrpc "gitlab.33.cn/chain33/chain33/rpc"
	"time"
)

// TODO: SetPostRunCb()
type RpcCtx struct {
	Addr   string
	Method string
	Params interface{}
	Res    interface{}

	cb Callback
}

type Callback func(res interface{}) (interface{}, error)

func NewRpcCtx(methed string, params, res interface{}) *RpcCtx {
	time.Sleep(1 * time.Millisecond) // 这里sleep下避免因为速度太快导致没来的记清理端口宠儿引发端口无法使用的错误
	return &RpcCtx{
		Addr:   "http://localhost:8801",
		Method: methed,
		Params: params,
		Res:    res,
	}
}

func (c *RpcCtx) SetResultCb(cb Callback) {
	c.cb = cb
}

func (c *RpcCtx) Run() (err error) {
	rpc, err := jsonrpc.NewJSONClient(c.Addr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	err = rpc.Call(c.Method, c.Params, c.Res)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// maybe format rpc result
	var result interface{}
	if c.cb != nil {
		result, err = c.cb(c.Res)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	} else {
		result = c.Res
	}

	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println(string(data))
	return
}
