package executor

import (
	tokenty "gitlab.33.cn/chain33/chain33/plugin/dapp/token/types"
	"gitlab.33.cn/chain33/chain33/types"
)

func (t *token) Query_GetTokens(in *tokenty.ReqTokens) (types.Message, error) {
	if in == nil {
		return nil, types.ErrInvalidParam
	}
	return t.GetTokens(in)
}

func (t *token) Query_GetTokenInfo(in *types.ReqString) (types.Message, error) {
	if in == nil {
		return nil, types.ErrInvalidParam
	}
	return t.GetTokenInfo(in.GetData())
}

func (t *token) Query_GetAddrReceiverforTokens(in *tokenty.ReqAddressToken) (types.Message, error) {
	if in == nil {
		return nil, types.ErrInvalidParam
	}
	return t.GetAddrReceiverforTokens(in)
}

func (t *token) Query_GetAccountTokenAssets(in *tokenty.ReqAccountTokenAssets) (types.Message, error) {
	if in == nil {
		return nil, types.ErrInvalidParam
	}
	return t.GetAccountTokenAssets(in)
}

func (t *token) Query_GetTxByToken(in *tokenty.ReqTokenTx) (types.Message, error) {
	if in == nil {
		return nil, types.ErrInvalidParam
	}
	if !types.GetSaveTokenTxList() {
		return nil, types.ErrActionNotSupport
	}
	return t.GetTxByToken(in)
}