// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	account.proto
	blockchain.proto
	common.proto
	config.proto
	db.proto
	executor.proto
	p2p.proto
	rpc.proto
	transaction.proto
	wallet.proto

It has these top-level messages:
	Account
	ReceiptExecAccountTransfer
	ReceiptAccountTransfer
	ReceiptExecUTXOTrans4Privacy
	ReqBalance
	Accounts
	ReqTokenBalance
	Header
	Block
	Blocks
	BlockDetails
	Headers
	BlockOverview
	BlockDetail
	Receipts
	PrivacyKV
	PrivacyKVToken
	LocalDBSetWithPrivacy
	ReceiptsAndPrivacyKV
	ReceiptCheckTxList
	ChainStatus
	ReqBlocks
	MempoolSize
	ReplyBlockHeight
	BlockBody
	IsCaughtUp
	IsNtpClockSync
	Reply
	ReqString
	ReplyString
	ReplyStrings
	ReqInt
	Int64
	ReqHash
	ReplyHash
	ReqNil
	ReqHashes
	ReplyHashes
	KeyValue
	Config
	Log
	MemPool
	Consensus
	Wallet
	Store
	LocalStore
	BlockChain
	P2P
	Rpc
	Exec
	LeafNode
	InnerNode
	MAVLProof
	StoreNode
	LocalDBSet
	LocalDBList
	LocalDBGet
	LocalReplyValue
	StoreSet
	StoreSetWithSync
	StoreGet
	StoreReplyValue
	Genesis
	CoinsAction
	CoinsGenesis
	CoinsTransfer
	CoinsWithdraw
	Hashlock
	HashlockAction
	HashlockLock
	HashlockUnlock
	HashlockSend
	HashRecv
	Hashlockquery
	Ticket
	TicketAction
	TicketMiner
	TicketBind
	TicketOpen
	TicketGenesis
	TicketClose
	TicketList
	TicketInfos
	ReplyTicketList
	ReplyWalletTickets
	ReceiptTicket
	ReceiptTicketBind
	ExecTxList
	Query
	RetrievePara
	Retrieve
	RetrieveAction
	BackupRetrieve
	PreRetrieve
	PerformRetrieve
	CancelRetrieve
	ReqRetrieveInfo
	RetrieveQuery
	TokenAction
	TokenPreCreate
	TokenFinishCreate
	TokenRevokeCreate
	Token
	ReqTokens
	ReplyTokens
	ReceiptToken
	Trade
	TradeForSell
	SellOrder
	SellOrderReceipt
	ReqAddrTokens
	ReplySellOrders
	TokenRecv
	ReplyAddrRecvForTokens
	TradeForBuy
	TradeForRevokeSell
	ReceiptTradeBase
	ReceiptTradeSell
	ReceiptTradeBuy
	ReceiptTradeRevoke
	TradeBuyDone
	ReplyTradeBuyOrders
	ArrayConfig
	StringConfig
	Int32Config
	ConfigItem
	ModifyConfig
	ManageAction
	ReceiptConfig
	ReplyConfig
	ReqTokenSellOrder
	PrivacyAction
	Public2Privacy
	Privacy2Privacy
	Privacy2Public
	UTXOGlobalIndex
	KeyInput
	PrivacyInput
	KeyOutput
	PrivacyOutput
	GroupUTXOGlobalIndex
	LocalUTXOItem
	ReqUTXOPubKeys
	GroupUTXOPubKey
	ResUTXOPubKeys
	ReqPrivacyToken
	ReplyPrivacyAmounts
	ReplyUTXOsOfAmount
	ReceiptPrivacyOutput
	AmountsOfUTXO
	TokenNamesOfUTXO
	P2PGetPeerInfo
	P2PPeerInfo
	P2PVersion
	P2PVerAck
	P2PPing
	P2PPong
	P2PGetAddr
	P2PAddr
	P2PExternalInfo
	P2PGetBlocks
	P2PGetMempool
	P2PInv
	Inventory
	P2PGetData
	P2PTx
	P2PBlock
	BroadCastData
	P2PGetHeaders
	P2PHeaders
	InvData
	InvDatas
	Peer
	PeerList
	CreateTx
	UnsignTx
	SignedTx
	Transaction
	Signature
	AddrOverview
	ReqAddr
	ReqPrivacy
	HexTx
	ReplyTxInfo
	ReqTxList
	ReplyTxList
	TxHashList
	ReplyTxInfos
	ReceiptLog
	Receipt
	ReceiptData
	TxResult
	TransactionDetail
	TransactionDetails
	ReqAddrs
	WalletTxDetail
	WalletTxDetails
	WalletAccountStore
	WalletAccountPrivacy
	WalletPwHash
	WalletStatus
	WalletAccounts
	WalletAccount
	WalletUnLock
	GenSeedLang
	GetSeedByPw
	SaveSeedByPw
	ReplySeed
	ReqWalletSetPasswd
	ReqNewAccount
	MinerFlag
	ReqWalletTransactionList
	ReqWalletImportPrivKey
	ReqWalletSendToAddress
	ReqWalletSetFee
	ReqWalletSetLabel
	ReqWalletMergeBalance
	ReqStr
	ReplyStr
	ReqTokenPreCreate
	ReqTokenFinishCreate
	ReqTokenRevokeCreate
	ReqSellToken
	ReqBuyToken
	ReqRevokeSell
	ReqModifyConfig
	ReqPub2Pri
	ReqPri2Pri
	ReqPri2Pub
	ReqCreateUTXOs
	ReplyPrivacyPkPair
	ReqPrivacyBalance
	ReqPrivBal4AddrToken
	ReplyPrivacyBalance
	PrivacyDBStore
	UTXO
	ReqUTXOGlobalIndex
	UTXOBasic
	UTXOIndex4Amount
	ResUTXOGlobalIndex
	FTXOsSTXOsInOneTx
*/
package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Account 的信息
type Account struct {
	Currency int32  `protobuf:"varint,1,opt,name=currency" json:"currency,omitempty"`
	Balance  int64  `protobuf:"varint,2,opt,name=balance" json:"balance,omitempty"`
	Frozen   int64  `protobuf:"varint,3,opt,name=frozen" json:"frozen,omitempty"`
	Addr     string `protobuf:"bytes,4,opt,name=addr" json:"addr,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Account) GetCurrency() int32 {
	if m != nil {
		return m.Currency
	}
	return 0
}

func (m *Account) GetBalance() int64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Account) GetFrozen() int64 {
	if m != nil {
		return m.Frozen
	}
	return 0
}

func (m *Account) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

// 账户余额改变的一个交易回报（合约内）
type ReceiptExecAccountTransfer struct {
	ExecAddr string   `protobuf:"bytes,1,opt,name=execAddr" json:"execAddr,omitempty"`
	Prev     *Account `protobuf:"bytes,2,opt,name=prev" json:"prev,omitempty"`
	Current  *Account `protobuf:"bytes,3,opt,name=current" json:"current,omitempty"`
}

func (m *ReceiptExecAccountTransfer) Reset()                    { *m = ReceiptExecAccountTransfer{} }
func (m *ReceiptExecAccountTransfer) String() string            { return proto.CompactTextString(m) }
func (*ReceiptExecAccountTransfer) ProtoMessage()               {}
func (*ReceiptExecAccountTransfer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReceiptExecAccountTransfer) GetExecAddr() string {
	if m != nil {
		return m.ExecAddr
	}
	return ""
}

func (m *ReceiptExecAccountTransfer) GetPrev() *Account {
	if m != nil {
		return m.Prev
	}
	return nil
}

func (m *ReceiptExecAccountTransfer) GetCurrent() *Account {
	if m != nil {
		return m.Current
	}
	return nil
}

// 账户余额改变的一个交易回报（coins内）
type ReceiptAccountTransfer struct {
	Prev    *Account `protobuf:"bytes,1,opt,name=prev" json:"prev,omitempty"`
	Current *Account `protobuf:"bytes,2,opt,name=current" json:"current,omitempty"`
}

func (m *ReceiptAccountTransfer) Reset()                    { *m = ReceiptAccountTransfer{} }
func (m *ReceiptAccountTransfer) String() string            { return proto.CompactTextString(m) }
func (*ReceiptAccountTransfer) ProtoMessage()               {}
func (*ReceiptAccountTransfer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReceiptAccountTransfer) GetPrev() *Account {
	if m != nil {
		return m.Prev
	}
	return nil
}

func (m *ReceiptAccountTransfer) GetCurrent() *Account {
	if m != nil {
		return m.Current
	}
	return nil
}

// 隐私合约内部通过UTXO进行转账的一个交易回报
type ReceiptExecUTXOTrans4Privacy struct {
	Type   int32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	Amount int64 `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
}

func (m *ReceiptExecUTXOTrans4Privacy) Reset()                    { *m = ReceiptExecUTXOTrans4Privacy{} }
func (m *ReceiptExecUTXOTrans4Privacy) String() string            { return proto.CompactTextString(m) }
func (*ReceiptExecUTXOTrans4Privacy) ProtoMessage()               {}
func (*ReceiptExecUTXOTrans4Privacy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReceiptExecUTXOTrans4Privacy) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReceiptExecUTXOTrans4Privacy) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// 查询一个地址列表在某个执行器中余额
type ReqBalance struct {
	Addresses []string `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty"`
	Execer    string   `protobuf:"bytes,2,opt,name=execer" json:"execer,omitempty"`
}

func (m *ReqBalance) Reset()                    { *m = ReqBalance{} }
func (m *ReqBalance) String() string            { return proto.CompactTextString(m) }
func (*ReqBalance) ProtoMessage()               {}
func (*ReqBalance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReqBalance) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *ReqBalance) GetExecer() string {
	if m != nil {
		return m.Execer
	}
	return ""
}

// Account 的列表
type Accounts struct {
	Acc []*Account `protobuf:"bytes,1,rep,name=acc" json:"acc,omitempty"`
}

func (m *Accounts) Reset()                    { *m = Accounts{} }
func (m *Accounts) String() string            { return proto.CompactTextString(m) }
func (*Accounts) ProtoMessage()               {}
func (*Accounts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Accounts) GetAcc() []*Account {
	if m != nil {
		return m.Acc
	}
	return nil
}

type ReqTokenBalance struct {
	Addresses   []string `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty"`
	TokenSymbol string   `protobuf:"bytes,2,opt,name=tokenSymbol" json:"tokenSymbol,omitempty"`
	Execer      string   `protobuf:"bytes,3,opt,name=execer" json:"execer,omitempty"`
}

func (m *ReqTokenBalance) Reset()                    { *m = ReqTokenBalance{} }
func (m *ReqTokenBalance) String() string            { return proto.CompactTextString(m) }
func (*ReqTokenBalance) ProtoMessage()               {}
func (*ReqTokenBalance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ReqTokenBalance) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *ReqTokenBalance) GetTokenSymbol() string {
	if m != nil {
		return m.TokenSymbol
	}
	return ""
}

func (m *ReqTokenBalance) GetExecer() string {
	if m != nil {
		return m.Execer
	}
	return ""
}

func init() {
	proto.RegisterType((*Account)(nil), "types.Account")
	proto.RegisterType((*ReceiptExecAccountTransfer)(nil), "types.ReceiptExecAccountTransfer")
	proto.RegisterType((*ReceiptAccountTransfer)(nil), "types.ReceiptAccountTransfer")
	proto.RegisterType((*ReceiptExecUTXOTrans4Privacy)(nil), "types.ReceiptExecUTXOTrans4Privacy")
	proto.RegisterType((*ReqBalance)(nil), "types.ReqBalance")
	proto.RegisterType((*Accounts)(nil), "types.Accounts")
	proto.RegisterType((*ReqTokenBalance)(nil), "types.ReqTokenBalance")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6a, 0xc2, 0x40,
	0x10, 0xc6, 0x59, 0xe3, 0xbf, 0x8c, 0xb4, 0x85, 0x3d, 0x48, 0x10, 0x0f, 0x61, 0x4f, 0x39, 0x14,
	0x0f, 0x6d, 0x5f, 0xa0, 0x42, 0x2f, 0xbd, 0xb4, 0x6c, 0x2d, 0xf4, 0xba, 0xae, 0x23, 0x88, 0x9a,
	0xc4, 0x4d, 0x94, 0xa6, 0x0f, 0xd0, 0xe7, 0x2e, 0x33, 0x6e, 0x6c, 0xa8, 0x50, 0xbc, 0xed, 0x37,
	0xb3, 0x33, 0xdf, 0x6f, 0x3f, 0x16, 0xae, 0x8c, 0xb5, 0xd9, 0x3e, 0x2d, 0x27, 0xb9, 0xcb, 0xca,
	0x4c, 0x76, 0xca, 0x2a, 0xc7, 0x42, 0xad, 0xa1, 0xf7, 0x78, 0xac, 0xcb, 0x11, 0xf4, 0xed, 0xde,
	0x39, 0x4c, 0x6d, 0x15, 0x89, 0x58, 0x24, 0x1d, 0x7d, 0xd2, 0x32, 0x82, 0xde, 0xdc, 0x6c, 0x4c,
	0x6a, 0x31, 0x6a, 0xc5, 0x22, 0x09, 0x74, 0x2d, 0xe5, 0x10, 0xba, 0x4b, 0x97, 0x7d, 0x61, 0x1a,
	0x05, 0xdc, 0xf0, 0x4a, 0x4a, 0x68, 0x9b, 0xc5, 0xc2, 0x45, 0xed, 0x58, 0x24, 0xa1, 0xe6, 0xb3,
	0xfa, 0x16, 0x30, 0xd2, 0x68, 0x71, 0x95, 0x97, 0x4f, 0x9f, 0x68, 0xbd, 0xf1, 0xcc, 0x99, 0xb4,
	0x58, 0xa2, 0x23, 0x00, 0xa4, 0x32, 0x8d, 0x09, 0x1e, 0x3b, 0x69, 0xa9, 0xa0, 0x9d, 0x3b, 0x3c,
	0xb0, 0xfb, 0xe0, 0xee, 0x7a, 0xc2, 0xf4, 0x13, 0xbf, 0x41, 0x73, 0x4f, 0x26, 0xd0, 0x3b, 0x02,
	0x97, 0xcc, 0x72, 0x7e, 0xad, 0x6e, 0xab, 0x25, 0x0c, 0x3d, 0xc7, 0x5f, 0x86, 0xda, 0x47, 0x5c,
	0xe6, 0xd3, 0xfa, 0xdf, 0xe7, 0x19, 0xc6, 0x8d, 0xf7, 0xbe, 0xcf, 0x3e, 0x5e, 0xd8, 0xe8, 0xe1,
	0xd5, 0xad, 0x0e, 0xc6, 0x56, 0x14, 0x12, 0x4d, 0xfa, 0xb8, 0xf9, 0x4c, 0x81, 0x9a, 0x2d, 0xad,
	0xf1, 0x49, 0x7b, 0xa5, 0xa6, 0x00, 0x1a, 0x77, 0x53, 0x1f, 0xfb, 0x18, 0x42, 0x8a, 0x14, 0x8b,
	0x02, 0x8b, 0x48, 0xc4, 0x41, 0x12, 0xea, 0xdf, 0x02, 0xed, 0xa0, 0xe4, 0xd0, 0xf1, 0x8e, 0x50,
	0x7b, 0xa5, 0x6e, 0xa1, 0xef, 0x19, 0x0b, 0x19, 0x43, 0x60, 0xac, 0xe5, 0xd9, 0xf3, 0x17, 0x50,
	0x4b, 0xad, 0xe0, 0x46, 0xe3, 0x6e, 0x96, 0xad, 0x31, 0xbd, 0xcc, 0x36, 0x86, 0x41, 0x49, 0xb7,
	0xdf, 0xaa, 0xed, 0x3c, 0xdb, 0x78, 0xef, 0x66, 0xa9, 0x01, 0x16, 0x34, 0xc1, 0xe6, 0x5d, 0xfe,
	0x94, 0xf7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xae, 0x64, 0xc6, 0xa5, 0x02, 0x00, 0x00,
}
