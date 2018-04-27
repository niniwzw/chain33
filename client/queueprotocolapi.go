package client

import (
	"gitlab.33.cn/chain33/chain33/types"
)

// 消息通道交互API接口定义
type QueueProtocolAPI interface {
	Close()
	Version() (*types.Reply, error)
	// 同步发送交易信息到指定模块，获取应答消息 types.EventTx
	SendTx(param *types.Transaction) (*types.Reply, error)
	// types.EventTxList
	GetTxList(param *types.TxHashList) (*types.ReplyTxList, error)
	// types.EventGetMempool
	GetMempool() (*types.ReplyTxList, error)
	// types.EventGetLastMempool
	GetLastMempool() (*types.ReplyTxList, error)

	// types.EventGetBlocks
	GetBlocks(param *types.ReqBlocks) (*types.BlockDetails, error)
	// types.EventQueryTx
	QueryTx(param *types.ReqHash) (*types.TransactionDetail, error)
	// types.EventGetTransactionByAddr
	GetTransactionByAddr(param *types.ReqAddr) (*types.ReplyTxInfos, error)
	// types.EventGetTransactionByHash
	GetTransactionByHash(param *types.ReqHashes) (*types.TransactionDetails, error)
	// types.EventGetHeaders
	GetHeaders(param *types.ReqBlocks) (*types.Headers, error)
	// types.EventGetBlockOverview
	GetBlockOverview(param *types.ReqHash) (*types.BlockOverview, error)
	// types.EventGetAddrOverview
	GetAddrOverview(param *types.ReqAddr) (*types.AddrOverview, error)
	// types.EventGetBlockHash
	GetBlockHash(param *types.ReqInt) (*types.ReplyHash, error)
	// types.EventIsSync
	IsSync() (*types.Reply, error)
	// types.EventIsNtpClockSync
	IsNtpClockSync() (*types.Reply, error)
	// types.EventLocalGet
	LocalGet(param *types.ReqHash) (*types.LocalReplyValue, error)
	// types.EventGetLastHeader
	GetLastHeader() (*types.Header, error)
	// types.EventQuery
	Query(param *types.Query) (*types.Message, error)

	// types.EventWalletGetAccountList
	WalletGetAccountList() (*types.WalletAccounts, error)
	// types.EventNewAccount
	NewAccount(param *types.ReqNewAccount) (*types.WalletAccount, error)
	// types.EventWalletTransactionList
	WalletTransactionList(param *types.ReqWalletTransactionList) (*types.WalletTxDetails, error)
	// types.EventWalletImportprivkey
	WalletImportprivkey(param *types.ReqWalletImportPrivKey) (*types.WalletAccount, error)
	// types.EventWalletSendToAddress
	WalletSendToAddress(param *types.ReqWalletSendToAddress) (*types.ReplyHash, error)
	// types.EventWalletSetFee
	WalletSetFee(param *types.ReqWalletSetFee) (*types.Reply, error)
	// types.EventWalletSetLabel
	WalletSetLabel(param *types.ReqWalletSetLabel) (*types.WalletAccount, error)
	// types.EventWalletMergeBalance
	WalletMergeBalance(param *types.ReqWalletMergeBalance) (*types.ReplyHashes, error)
	// types.EventWalletSetPasswd
	WalletSetPasswd(param *types.ReqWalletSetPasswd) (*types.Reply, error)
	// types.EventWalletLock
	WalletLock() (*types.Reply, error)
	// types.EventWalletUnLock
	WalletUnLock(param *types.WalletUnLock) (*types.Reply, error)
	// types.EventGenSeed
	GenSeed(param *types.GenSeedLang) (*types.ReplySeed, error)
	// types.EventSaveSeed
	SaveSeed(param *types.SaveSeedByPw) (*types.Reply, error)
	// types.EventGetSeed
	GetSeed(param *types.GetSeedByPw) (*types.ReplySeed, error)
	// types.EventGetWalletStatus
	GetWalletStatus() (*types.WalletStatus, error)
	// types.EventWalletAutoMiner
	WalletAutoMiner(param *types.MinerFlag) (*types.Reply, error)
	// types.EventDumpPrivkey
	DumpPrivkey(param *types.ReqStr) (*types.ReplyStr, error)
	// types.EventCloseTickets
	CloseTickets() (*types.ReplyHashes, error)
	// types.EventSignRawTx
	SignRawTx(param *types.ReqSignRawTx) (*types.ReplySignRawTx, error)

	// types.EventPeerInfo
	PeerInfo() (*types.PeerList, error)

	// types.EventGetTicketCount
	GetTicketCount() (*types.Int64, error)

	CreateRawTransaction(param *types.CreateTx) ([]byte, error)
	SendRawTransaction(param *types.SignedTx) (*types.Reply, error)

	GetBalance(param *types.ReqBalance) ([]*types.Account, error)
	GetTokenBalance(param *types.ReqTokenBalance) ([]*types.Account, error)
	// types.EventGetNetInfo
	GetNetInfo() (*types.NodeNetInfo, error)
	GetTotalCoins(param *types.ReqGetTotalCoins) (*types.ReplyGetTotalCoins, error)

	CreateRawTokenPreCreateTx(parm *TokenPreCreateTx) (*types.Transaction, error)
	CreateRawTokenFinishTx(parm *TokenFinishTx) (*types.Transaction, error)
	CreateRawTokenRevokeTx(parm *TokenRevokeTx) (*types.Transaction, error)
	CreateRawTradeSellTx(parm *TradeSellTx) (*types.Transaction, error)
	CreateRawTradeRevokeTx(param *TradeRevokeTx) (*types.Transaction, error)
	CreateRawTradeBuyTx(param *TradeBuyTx) (*types.Transaction, error)
}
