// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Config struct {
	Title      string      `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Log        *Log        `protobuf:"bytes,2,opt,name=log" json:"log,omitempty"`
	Store      *Store      `protobuf:"bytes,3,opt,name=store" json:"store,omitempty"`
	Consensus  *Consensus  `protobuf:"bytes,5,opt,name=consensus" json:"consensus,omitempty"`
	MemPool    *MemPool    `protobuf:"bytes,6,opt,name=memPool" json:"memPool,omitempty"`
	BlockChain *BlockChain `protobuf:"bytes,7,opt,name=blockChain" json:"blockChain,omitempty"`
	Wallet     *Wallet     `protobuf:"bytes,8,opt,name=wallet" json:"wallet,omitempty"`
	P2P        *P2P        `protobuf:"bytes,9,opt,name=p2p" json:"p2p,omitempty"`
	Rpc        *Rpc        `protobuf:"bytes,10,opt,name=rpc" json:"rpc,omitempty"`
	Exec       *Exec       `protobuf:"bytes,11,opt,name=exec" json:"exec,omitempty"`
	TestNet    bool        `protobuf:"varint,12,opt,name=testNet" json:"testNet,omitempty"`
	FixTime    bool        `protobuf:"varint,13,opt,name=fixTime" json:"fixTime,omitempty"`
	Pprof      *Pprof      `protobuf:"bytes,14,opt,name=pprof" json:"pprof,omitempty"`
	Auth       *Authority  `protobuf:"bytes,15,opt,name=auth" json:"auth,omitempty"`
	Game       *GameExec   `protobuf:"bytes,16,opt,name=game" json:"game,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Config) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Config) GetLog() *Log {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *Config) GetStore() *Store {
	if m != nil {
		return m.Store
	}
	return nil
}

func (m *Config) GetConsensus() *Consensus {
	if m != nil {
		return m.Consensus
	}
	return nil
}

func (m *Config) GetMemPool() *MemPool {
	if m != nil {
		return m.MemPool
	}
	return nil
}

func (m *Config) GetBlockChain() *BlockChain {
	if m != nil {
		return m.BlockChain
	}
	return nil
}

func (m *Config) GetWallet() *Wallet {
	if m != nil {
		return m.Wallet
	}
	return nil
}

func (m *Config) GetP2P() *P2P {
	if m != nil {
		return m.P2P
	}
	return nil
}

func (m *Config) GetRpc() *Rpc {
	if m != nil {
		return m.Rpc
	}
	return nil
}

func (m *Config) GetExec() *Exec {
	if m != nil {
		return m.Exec
	}
	return nil
}

func (m *Config) GetTestNet() bool {
	if m != nil {
		return m.TestNet
	}
	return false
}

func (m *Config) GetFixTime() bool {
	if m != nil {
		return m.FixTime
	}
	return false
}

func (m *Config) GetPprof() *Pprof {
	if m != nil {
		return m.Pprof
	}
	return nil
}

func (m *Config) GetAuth() *Authority {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *Config) GetGame() *GameExec {
	if m != nil {
		return m.Game
	}
	return nil
}

type Log struct {
	// 日志级别，支持debug(dbug)/info/warn/error(eror)/crit
	Loglevel        string `protobuf:"bytes,1,opt,name=loglevel" json:"loglevel,omitempty"`
	LogConsoleLevel string `protobuf:"bytes,2,opt,name=logConsoleLevel" json:"logConsoleLevel,omitempty"`
	// 日志文件名，可带目录，所有生成的日志文件都放到此目录下
	LogFile string `protobuf:"bytes,3,opt,name=logFile" json:"logFile,omitempty"`
	// 单个日志文件的最大值（单位：兆）
	MaxFileSize uint32 `protobuf:"varint,4,opt,name=maxFileSize" json:"maxFileSize,omitempty"`
	// 最多保存的历史日志文件个数
	MaxBackups uint32 `protobuf:"varint,5,opt,name=maxBackups" json:"maxBackups,omitempty"`
	// 最多保存的历史日志消息（单位：天）
	MaxAge uint32 `protobuf:"varint,6,opt,name=maxAge" json:"maxAge,omitempty"`
	// 日志文件名是否使用本地事件（否则使用UTC时间）
	LocalTime bool `protobuf:"varint,7,opt,name=localTime" json:"localTime,omitempty"`
	// 历史日志文件是否压缩（压缩格式为gz）
	Compress bool `protobuf:"varint,8,opt,name=compress" json:"compress,omitempty"`
	// 是否打印调用源文件和行号
	CallerFile bool `protobuf:"varint,9,opt,name=callerFile" json:"callerFile,omitempty"`
	// 是否打印调用方法
	CallerFunction bool `protobuf:"varint,10,opt,name=callerFunction" json:"callerFunction,omitempty"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *Log) GetLoglevel() string {
	if m != nil {
		return m.Loglevel
	}
	return ""
}

func (m *Log) GetLogConsoleLevel() string {
	if m != nil {
		return m.LogConsoleLevel
	}
	return ""
}

func (m *Log) GetLogFile() string {
	if m != nil {
		return m.LogFile
	}
	return ""
}

func (m *Log) GetMaxFileSize() uint32 {
	if m != nil {
		return m.MaxFileSize
	}
	return 0
}

func (m *Log) GetMaxBackups() uint32 {
	if m != nil {
		return m.MaxBackups
	}
	return 0
}

func (m *Log) GetMaxAge() uint32 {
	if m != nil {
		return m.MaxAge
	}
	return 0
}

func (m *Log) GetLocalTime() bool {
	if m != nil {
		return m.LocalTime
	}
	return false
}

func (m *Log) GetCompress() bool {
	if m != nil {
		return m.Compress
	}
	return false
}

func (m *Log) GetCallerFile() bool {
	if m != nil {
		return m.CallerFile
	}
	return false
}

func (m *Log) GetCallerFunction() bool {
	if m != nil {
		return m.CallerFunction
	}
	return false
}

type MemPool struct {
	PoolCacheSize      int64 `protobuf:"varint,1,opt,name=poolCacheSize" json:"poolCacheSize,omitempty"`
	MinTxFee           int64 `protobuf:"varint,2,opt,name=minTxFee" json:"minTxFee,omitempty"`
	ForceAccept        bool  `protobuf:"varint,3,opt,name=forceAccept" json:"forceAccept,omitempty"`
	MaxTxNumPerAccount int64 `protobuf:"varint,4,opt,name=maxTxNumPerAccount" json:"maxTxNumPerAccount,omitempty"`
}

func (m *MemPool) Reset()                    { *m = MemPool{} }
func (m *MemPool) String() string            { return proto.CompactTextString(m) }
func (*MemPool) ProtoMessage()               {}
func (*MemPool) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *MemPool) GetPoolCacheSize() int64 {
	if m != nil {
		return m.PoolCacheSize
	}
	return 0
}

func (m *MemPool) GetMinTxFee() int64 {
	if m != nil {
		return m.MinTxFee
	}
	return 0
}

func (m *MemPool) GetForceAccept() bool {
	if m != nil {
		return m.ForceAccept
	}
	return false
}

func (m *MemPool) GetMaxTxNumPerAccount() int64 {
	if m != nil {
		return m.MaxTxNumPerAccount
	}
	return 0
}

type Consensus struct {
	Name                 string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Genesis              string `protobuf:"bytes,2,opt,name=genesis" json:"genesis,omitempty"`
	Minerstart           bool   `protobuf:"varint,3,opt,name=minerstart" json:"minerstart,omitempty"`
	GenesisBlockTime     int64  `protobuf:"varint,4,opt,name=genesisBlockTime" json:"genesisBlockTime,omitempty"`
	HotkeyAddr           string `protobuf:"bytes,5,opt,name=hotkeyAddr" json:"hotkeyAddr,omitempty"`
	ForceMining          bool   `protobuf:"varint,6,opt,name=forceMining" json:"forceMining,omitempty"`
	NodeId               int64  `protobuf:"varint,7,opt,name=NodeId" json:"NodeId,omitempty"`
	PeersURL             string `protobuf:"bytes,8,opt,name=PeersURL" json:"PeersURL,omitempty"`
	ClientAddr           string `protobuf:"bytes,9,opt,name=ClientAddr" json:"ClientAddr,omitempty"`
	RaftApiPort          int64  `protobuf:"varint,15,opt,name=raftApiPort" json:"raftApiPort,omitempty"`
	IsNewJoinNode        bool   `protobuf:"varint,16,opt,name=isNewJoinNode" json:"isNewJoinNode,omitempty"`
	ReadOnlyPeersURL     string `protobuf:"bytes,17,opt,name=readOnlyPeersURL" json:"readOnlyPeersURL,omitempty"`
	AddPeersURL          string `protobuf:"bytes,18,opt,name=addPeersURL" json:"addPeersURL,omitempty"`
	DefaultSnapCount     int64  `protobuf:"varint,19,opt,name=defaultSnapCount" json:"defaultSnapCount,omitempty"`
	WriteBlockSeconds    int64  `protobuf:"varint,20,opt,name=writeBlockSeconds" json:"writeBlockSeconds,omitempty"`
	HeartbeatTick        int32  `protobuf:"varint,21,opt,name=heartbeatTick" json:"heartbeatTick,omitempty"`
	ParaRemoteGrpcClient string `protobuf:"bytes,22,opt,name=paraRemoteGrpcClient" json:"paraRemoteGrpcClient,omitempty"`
	StartHeight          int64  `protobuf:"varint,23,opt,name=startHeight" json:"startHeight,omitempty"`
	EmptyBlockInterval   int64  `protobuf:"varint,24,opt,name=emptyBlockInterval" json:"emptyBlockInterval,omitempty"`
	AuthAccount          string `protobuf:"bytes,25,opt,name=authAccount" json:"authAccount,omitempty"`
	WaitBlocks4CommitMsg int32  `protobuf:"varint,26,opt,name=waitBlocks4CommitMsg" json:"waitBlocks4CommitMsg,omitempty"`
}

func (m *Consensus) Reset()                    { *m = Consensus{} }
func (m *Consensus) String() string            { return proto.CompactTextString(m) }
func (*Consensus) ProtoMessage()               {}
func (*Consensus) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *Consensus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Consensus) GetGenesis() string {
	if m != nil {
		return m.Genesis
	}
	return ""
}

func (m *Consensus) GetMinerstart() bool {
	if m != nil {
		return m.Minerstart
	}
	return false
}

func (m *Consensus) GetGenesisBlockTime() int64 {
	if m != nil {
		return m.GenesisBlockTime
	}
	return 0
}

func (m *Consensus) GetHotkeyAddr() string {
	if m != nil {
		return m.HotkeyAddr
	}
	return ""
}

func (m *Consensus) GetForceMining() bool {
	if m != nil {
		return m.ForceMining
	}
	return false
}

func (m *Consensus) GetNodeId() int64 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *Consensus) GetPeersURL() string {
	if m != nil {
		return m.PeersURL
	}
	return ""
}

func (m *Consensus) GetClientAddr() string {
	if m != nil {
		return m.ClientAddr
	}
	return ""
}

func (m *Consensus) GetRaftApiPort() int64 {
	if m != nil {
		return m.RaftApiPort
	}
	return 0
}

func (m *Consensus) GetIsNewJoinNode() bool {
	if m != nil {
		return m.IsNewJoinNode
	}
	return false
}

func (m *Consensus) GetReadOnlyPeersURL() string {
	if m != nil {
		return m.ReadOnlyPeersURL
	}
	return ""
}

func (m *Consensus) GetAddPeersURL() string {
	if m != nil {
		return m.AddPeersURL
	}
	return ""
}

func (m *Consensus) GetDefaultSnapCount() int64 {
	if m != nil {
		return m.DefaultSnapCount
	}
	return 0
}

func (m *Consensus) GetWriteBlockSeconds() int64 {
	if m != nil {
		return m.WriteBlockSeconds
	}
	return 0
}

func (m *Consensus) GetHeartbeatTick() int32 {
	if m != nil {
		return m.HeartbeatTick
	}
	return 0
}

func (m *Consensus) GetParaRemoteGrpcClient() string {
	if m != nil {
		return m.ParaRemoteGrpcClient
	}
	return ""
}

func (m *Consensus) GetStartHeight() int64 {
	if m != nil {
		return m.StartHeight
	}
	return 0
}

func (m *Consensus) GetEmptyBlockInterval() int64 {
	if m != nil {
		return m.EmptyBlockInterval
	}
	return 0
}

func (m *Consensus) GetAuthAccount() string {
	if m != nil {
		return m.AuthAccount
	}
	return ""
}

func (m *Consensus) GetWaitBlocks4CommitMsg() int32 {
	if m != nil {
		return m.WaitBlocks4CommitMsg
	}
	return 0
}

type Wallet struct {
	MinFee         int64    `protobuf:"varint,1,opt,name=minFee" json:"minFee,omitempty"`
	Driver         string   `protobuf:"bytes,2,opt,name=driver" json:"driver,omitempty"`
	DbPath         string   `protobuf:"bytes,3,opt,name=dbPath" json:"dbPath,omitempty"`
	DbCache        int32    `protobuf:"varint,4,opt,name=dbCache" json:"dbCache,omitempty"`
	SignType       string   `protobuf:"bytes,5,opt,name=signType" json:"signType,omitempty"`
	ForceMining    bool     `protobuf:"varint,6,opt,name=forceMining" json:"forceMining,omitempty"`
	Minerdisable   bool     `protobuf:"varint,7,opt,name=minerdisable" json:"minerdisable,omitempty"`
	Minerwhitelist []string `protobuf:"bytes,8,rep,name=minerwhitelist" json:"minerwhitelist,omitempty"`
}

func (m *Wallet) Reset()                    { *m = Wallet{} }
func (m *Wallet) String() string            { return proto.CompactTextString(m) }
func (*Wallet) ProtoMessage()               {}
func (*Wallet) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *Wallet) GetMinFee() int64 {
	if m != nil {
		return m.MinFee
	}
	return 0
}

func (m *Wallet) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *Wallet) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
}

func (m *Wallet) GetDbCache() int32 {
	if m != nil {
		return m.DbCache
	}
	return 0
}

func (m *Wallet) GetSignType() string {
	if m != nil {
		return m.SignType
	}
	return ""
}

func (m *Wallet) GetForceMining() bool {
	if m != nil {
		return m.ForceMining
	}
	return false
}

func (m *Wallet) GetMinerdisable() bool {
	if m != nil {
		return m.Minerdisable
	}
	return false
}

func (m *Wallet) GetMinerwhitelist() []string {
	if m != nil {
		return m.Minerwhitelist
	}
	return nil
}

type Store struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Driver  string `protobuf:"bytes,2,opt,name=driver" json:"driver,omitempty"`
	DbPath  string `protobuf:"bytes,3,opt,name=dbPath" json:"dbPath,omitempty"`
	DbCache int32  `protobuf:"varint,4,opt,name=dbCache" json:"dbCache,omitempty"`
}

func (m *Store) Reset()                    { *m = Store{} }
func (m *Store) String() string            { return proto.CompactTextString(m) }
func (*Store) ProtoMessage()               {}
func (*Store) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *Store) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Store) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *Store) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
}

func (m *Store) GetDbCache() int32 {
	if m != nil {
		return m.DbCache
	}
	return 0
}

type BlockChain struct {
	DefCacheSize          int64  `protobuf:"varint,1,opt,name=defCacheSize" json:"defCacheSize,omitempty"`
	MaxFetchBlockNum      int64  `protobuf:"varint,2,opt,name=maxFetchBlockNum" json:"maxFetchBlockNum,omitempty"`
	TimeoutSeconds        int64  `protobuf:"varint,3,opt,name=timeoutSeconds" json:"timeoutSeconds,omitempty"`
	BatchBlockNum         int64  `protobuf:"varint,4,opt,name=batchBlockNum" json:"batchBlockNum,omitempty"`
	Driver                string `protobuf:"bytes,5,opt,name=driver" json:"driver,omitempty"`
	DbPath                string `protobuf:"bytes,6,opt,name=dbPath" json:"dbPath,omitempty"`
	DbCache               int32  `protobuf:"varint,7,opt,name=dbCache" json:"dbCache,omitempty"`
	IsStrongConsistency   bool   `protobuf:"varint,8,opt,name=isStrongConsistency" json:"isStrongConsistency,omitempty"`
	SingleMode            bool   `protobuf:"varint,9,opt,name=singleMode" json:"singleMode,omitempty"`
	Batchsync             bool   `protobuf:"varint,10,opt,name=batchsync" json:"batchsync,omitempty"`
	IsRecordBlockSequence bool   `protobuf:"varint,11,opt,name=isRecordBlockSequence" json:"isRecordBlockSequence,omitempty"`
	IsParaChain           bool   `protobuf:"varint,12,opt,name=isParaChain" json:"isParaChain,omitempty"`
	EnableTxQuickIndex    bool   `protobuf:"varint,13,opt,name=enableTxQuickIndex" json:"enableTxQuickIndex,omitempty"`
}

func (m *BlockChain) Reset()                    { *m = BlockChain{} }
func (m *BlockChain) String() string            { return proto.CompactTextString(m) }
func (*BlockChain) ProtoMessage()               {}
func (*BlockChain) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *BlockChain) GetDefCacheSize() int64 {
	if m != nil {
		return m.DefCacheSize
	}
	return 0
}

func (m *BlockChain) GetMaxFetchBlockNum() int64 {
	if m != nil {
		return m.MaxFetchBlockNum
	}
	return 0
}

func (m *BlockChain) GetTimeoutSeconds() int64 {
	if m != nil {
		return m.TimeoutSeconds
	}
	return 0
}

func (m *BlockChain) GetBatchBlockNum() int64 {
	if m != nil {
		return m.BatchBlockNum
	}
	return 0
}

func (m *BlockChain) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *BlockChain) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
}

func (m *BlockChain) GetDbCache() int32 {
	if m != nil {
		return m.DbCache
	}
	return 0
}

func (m *BlockChain) GetIsStrongConsistency() bool {
	if m != nil {
		return m.IsStrongConsistency
	}
	return false
}

func (m *BlockChain) GetSingleMode() bool {
	if m != nil {
		return m.SingleMode
	}
	return false
}

func (m *BlockChain) GetBatchsync() bool {
	if m != nil {
		return m.Batchsync
	}
	return false
}

func (m *BlockChain) GetIsRecordBlockSequence() bool {
	if m != nil {
		return m.IsRecordBlockSequence
	}
	return false
}

func (m *BlockChain) GetIsParaChain() bool {
	if m != nil {
		return m.IsParaChain
	}
	return false
}

func (m *BlockChain) GetEnableTxQuickIndex() bool {
	if m != nil {
		return m.EnableTxQuickIndex
	}
	return false
}

type P2P struct {
	SeedPort        int32    `protobuf:"varint,1,opt,name=seedPort" json:"seedPort,omitempty"`
	Driver          string   `protobuf:"bytes,2,opt,name=driver" json:"driver,omitempty"`
	DbPath          string   `protobuf:"bytes,3,opt,name=dbPath" json:"dbPath,omitempty"`
	DbCache         int32    `protobuf:"varint,4,opt,name=dbCache" json:"dbCache,omitempty"`
	GrpcLogFile     string   `protobuf:"bytes,5,opt,name=grpcLogFile" json:"grpcLogFile,omitempty"`
	IsSeed          bool     `protobuf:"varint,6,opt,name=isSeed" json:"isSeed,omitempty"`
	ServerStart     bool     `protobuf:"varint,7,opt,name=serverStart" json:"serverStart,omitempty"`
	Seeds           []string `protobuf:"bytes,8,rep,name=seeds" json:"seeds,omitempty"`
	Enable          bool     `protobuf:"varint,9,opt,name=enable" json:"enable,omitempty"`
	MsgCacheSize    int32    `protobuf:"varint,10,opt,name=msgCacheSize" json:"msgCacheSize,omitempty"`
	Version         int32    `protobuf:"varint,11,opt,name=version" json:"version,omitempty"`
	VerMix          int32    `protobuf:"varint,12,opt,name=verMix" json:"verMix,omitempty"`
	VerMax          int32    `protobuf:"varint,13,opt,name=verMax" json:"verMax,omitempty"`
	InnerSeedEnable bool     `protobuf:"varint,14,opt,name=innerSeedEnable" json:"innerSeedEnable,omitempty"`
	InnerBounds     int32    `protobuf:"varint,15,opt,name=innerBounds" json:"innerBounds,omitempty"`
	UseGithub       bool     `protobuf:"varint,16,opt,name=useGithub" json:"useGithub,omitempty"`
}

func (m *P2P) Reset()                    { *m = P2P{} }
func (m *P2P) String() string            { return proto.CompactTextString(m) }
func (*P2P) ProtoMessage()               {}
func (*P2P) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

func (m *P2P) GetSeedPort() int32 {
	if m != nil {
		return m.SeedPort
	}
	return 0
}

func (m *P2P) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *P2P) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
}

func (m *P2P) GetDbCache() int32 {
	if m != nil {
		return m.DbCache
	}
	return 0
}

func (m *P2P) GetGrpcLogFile() string {
	if m != nil {
		return m.GrpcLogFile
	}
	return ""
}

func (m *P2P) GetIsSeed() bool {
	if m != nil {
		return m.IsSeed
	}
	return false
}

func (m *P2P) GetServerStart() bool {
	if m != nil {
		return m.ServerStart
	}
	return false
}

func (m *P2P) GetSeeds() []string {
	if m != nil {
		return m.Seeds
	}
	return nil
}

func (m *P2P) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *P2P) GetMsgCacheSize() int32 {
	if m != nil {
		return m.MsgCacheSize
	}
	return 0
}

func (m *P2P) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *P2P) GetVerMix() int32 {
	if m != nil {
		return m.VerMix
	}
	return 0
}

func (m *P2P) GetVerMax() int32 {
	if m != nil {
		return m.VerMax
	}
	return 0
}

func (m *P2P) GetInnerSeedEnable() bool {
	if m != nil {
		return m.InnerSeedEnable
	}
	return false
}

func (m *P2P) GetInnerBounds() int32 {
	if m != nil {
		return m.InnerBounds
	}
	return 0
}

func (m *P2P) GetUseGithub() bool {
	if m != nil {
		return m.UseGithub
	}
	return false
}

type Rpc struct {
	JrpcBindAddr      string   `protobuf:"bytes,1,opt,name=jrpcBindAddr" json:"jrpcBindAddr,omitempty"`
	GrpcBindAddr      string   `protobuf:"bytes,2,opt,name=grpcBindAddr" json:"grpcBindAddr,omitempty"`
	Whitlist          []string `protobuf:"bytes,3,rep,name=whitlist" json:"whitlist,omitempty"`
	Whitelist         []string `protobuf:"bytes,4,rep,name=whitelist" json:"whitelist,omitempty"`
	JrpcFuncWhitelist []string `protobuf:"bytes,5,rep,name=jrpcFuncWhitelist" json:"jrpcFuncWhitelist,omitempty"`
	GrpcFuncWhitelist []string `protobuf:"bytes,6,rep,name=grpcFuncWhitelist" json:"grpcFuncWhitelist,omitempty"`
	JrpcFuncBlacklist []string `protobuf:"bytes,7,rep,name=jrpcFuncBlacklist" json:"jrpcFuncBlacklist,omitempty"`
	GrpcFuncBlacklist []string `protobuf:"bytes,8,rep,name=grpcFuncBlacklist" json:"grpcFuncBlacklist,omitempty"`
	MainnetJrpcAddr   string   `protobuf:"bytes,9,opt,name=mainnetJrpcAddr" json:"mainnetJrpcAddr,omitempty"`
}

func (m *Rpc) Reset()                    { *m = Rpc{} }
func (m *Rpc) String() string            { return proto.CompactTextString(m) }
func (*Rpc) ProtoMessage()               {}
func (*Rpc) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

func (m *Rpc) GetJrpcBindAddr() string {
	if m != nil {
		return m.JrpcBindAddr
	}
	return ""
}

func (m *Rpc) GetGrpcBindAddr() string {
	if m != nil {
		return m.GrpcBindAddr
	}
	return ""
}

func (m *Rpc) GetWhitlist() []string {
	if m != nil {
		return m.Whitlist
	}
	return nil
}

func (m *Rpc) GetWhitelist() []string {
	if m != nil {
		return m.Whitelist
	}
	return nil
}

func (m *Rpc) GetJrpcFuncWhitelist() []string {
	if m != nil {
		return m.JrpcFuncWhitelist
	}
	return nil
}

func (m *Rpc) GetGrpcFuncWhitelist() []string {
	if m != nil {
		return m.GrpcFuncWhitelist
	}
	return nil
}

func (m *Rpc) GetJrpcFuncBlacklist() []string {
	if m != nil {
		return m.JrpcFuncBlacklist
	}
	return nil
}

func (m *Rpc) GetGrpcFuncBlacklist() []string {
	if m != nil {
		return m.GrpcFuncBlacklist
	}
	return nil
}

func (m *Rpc) GetMainnetJrpcAddr() string {
	if m != nil {
		return m.MainnetJrpcAddr
	}
	return ""
}

type Exec struct {
	MinExecFee int64     `protobuf:"varint,1,opt,name=minExecFee" json:"minExecFee,omitempty"`
	IsFree     bool      `protobuf:"varint,2,opt,name=isFree" json:"isFree,omitempty"`
	EnableStat bool      `protobuf:"varint,3,opt,name=enableStat" json:"enableStat,omitempty"`
	EnableMVCC bool      `protobuf:"varint,4,opt,name=enableMVCC" json:"enableMVCC,omitempty"`
	Game       *GameExec `protobuf:"bytes,5,opt,name=game" json:"game,omitempty"`
}

func (m *Exec) Reset()                    { *m = Exec{} }
func (m *Exec) String() string            { return proto.CompactTextString(m) }
func (*Exec) ProtoMessage()               {}
func (*Exec) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

func (m *Exec) GetMinExecFee() int64 {
	if m != nil {
		return m.MinExecFee
	}
	return 0
}

func (m *Exec) GetIsFree() bool {
	if m != nil {
		return m.IsFree
	}
	return false
}

func (m *Exec) GetEnableStat() bool {
	if m != nil {
		return m.EnableStat
	}
	return false
}

func (m *Exec) GetEnableMVCC() bool {
	if m != nil {
		return m.EnableMVCC
	}
	return false
}

func (m *Exec) GetGame() *GameExec {
	if m != nil {
		return m.Game
	}
	return nil
}

type Authority struct {
	Enable     bool   `protobuf:"varint,1,opt,name=enable" json:"enable,omitempty"`
	CryptoPath string `protobuf:"bytes,2,opt,name=cryptoPath" json:"cryptoPath,omitempty"`
	SignType   string `protobuf:"bytes,3,opt,name=signType" json:"signType,omitempty"`
}

func (m *Authority) Reset()                    { *m = Authority{} }
func (m *Authority) String() string            { return proto.CompactTextString(m) }
func (*Authority) ProtoMessage()               {}
func (*Authority) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{10} }

func (m *Authority) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *Authority) GetCryptoPath() string {
	if m != nil {
		return m.CryptoPath
	}
	return ""
}

func (m *Authority) GetSignType() string {
	if m != nil {
		return m.SignType
	}
	return ""
}

type Pprof struct {
	ListenAddr string `protobuf:"bytes,1,opt,name=listenAddr" json:"listenAddr,omitempty"`
}

func (m *Pprof) Reset()                    { *m = Pprof{} }
func (m *Pprof) String() string            { return proto.CompactTextString(m) }
func (*Pprof) ProtoMessage()               {}
func (*Pprof) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{11} }

func (m *Pprof) GetListenAddr() string {
	if m != nil {
		return m.ListenAddr
	}
	return ""
}

type GameExec struct {
	ActiveTime    int32 `protobuf:"varint,1,opt,name=activeTime" json:"activeTime,omitempty"`
	DefaultCount  int32 `protobuf:"varint,2,opt,name=defaultCount" json:"defaultCount,omitempty"`
	MaxCount      int32 `protobuf:"varint,3,opt,name=maxCount" json:"maxCount,omitempty"`
	MaxGameAmount int64 `protobuf:"varint,4,opt,name=maxGameAmount" json:"maxGameAmount,omitempty"`
	MinGameAmount int64 `protobuf:"varint,5,opt,name=minGameAmount" json:"minGameAmount,omitempty"`
}

func (m *GameExec) Reset()                    { *m = GameExec{} }
func (m *GameExec) String() string            { return proto.CompactTextString(m) }
func (*GameExec) ProtoMessage()               {}
func (*GameExec) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{12} }

func (m *GameExec) GetActiveTime() int32 {
	if m != nil {
		return m.ActiveTime
	}
	return 0
}

func (m *GameExec) GetDefaultCount() int32 {
	if m != nil {
		return m.DefaultCount
	}
	return 0
}

func (m *GameExec) GetMaxCount() int32 {
	if m != nil {
		return m.MaxCount
	}
	return 0
}

func (m *GameExec) GetMaxGameAmount() int64 {
	if m != nil {
		return m.MaxGameAmount
	}
	return 0
}

func (m *GameExec) GetMinGameAmount() int64 {
	if m != nil {
		return m.MinGameAmount
	}
	return 0
}

func init() {
	proto.RegisterType((*Config)(nil), "types.Config")
	proto.RegisterType((*Log)(nil), "types.Log")
	proto.RegisterType((*MemPool)(nil), "types.MemPool")
	proto.RegisterType((*Consensus)(nil), "types.Consensus")
	proto.RegisterType((*Wallet)(nil), "types.Wallet")
	proto.RegisterType((*Store)(nil), "types.Store")
	proto.RegisterType((*BlockChain)(nil), "types.BlockChain")
	proto.RegisterType((*P2P)(nil), "types.P2P")
	proto.RegisterType((*Rpc)(nil), "types.Rpc")
	proto.RegisterType((*Exec)(nil), "types.Exec")
	proto.RegisterType((*Authority)(nil), "types.Authority")
	proto.RegisterType((*Pprof)(nil), "types.Pprof")
	proto.RegisterType((*GameExec)(nil), "types.GameExec")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 1576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x57, 0xcd, 0x6e, 0x2b, 0x49,
	0x15, 0x96, 0xd3, 0xb6, 0x63, 0x57, 0xfe, 0xee, 0xed, 0xb9, 0x33, 0x34, 0xa3, 0x11, 0x44, 0x66,
	0x00, 0x0b, 0xa1, 0x08, 0xc2, 0xbc, 0x40, 0x62, 0xcd, 0xbd, 0xcc, 0x28, 0x09, 0xa6, 0x1c, 0xb8,
	0x4b, 0x54, 0xae, 0x3e, 0x69, 0x17, 0xe9, 0xee, 0x6a, 0xaa, 0xab, 0x13, 0x9b, 0x77, 0x60, 0xc7,
	0x96, 0x15, 0xef, 0xc0, 0x9e, 0x07, 0xe0, 0x41, 0x90, 0x58, 0xf2, 0x00, 0xe8, 0x9c, 0x2a, 0xf7,
	0x8f, 0xe3, 0x8b, 0x66, 0x71, 0x77, 0x3e, 0xdf, 0xf9, 0x5c, 0x55, 0xe7, 0xff, 0x34, 0x3b, 0x96,
	0x3a, 0x7f, 0x50, 0xc9, 0x45, 0x61, 0xb4, 0xd5, 0xe1, 0xc0, 0x6e, 0x0a, 0x28, 0x27, 0x7f, 0xed,
	0xb3, 0xe1, 0x8c, 0xf0, 0xf0, 0x0d, 0x1b, 0x58, 0x65, 0x53, 0x88, 0x7a, 0xe7, 0xbd, 0xe9, 0x98,
	0x3b, 0x21, 0xfc, 0x82, 0x05, 0xa9, 0x4e, 0xa2, 0x83, 0xf3, 0xde, 0xf4, 0xe8, 0x92, 0x5d, 0xd0,
	0xbf, 0x2e, 0x6e, 0x74, 0xc2, 0x11, 0x0e, 0x27, 0x6c, 0x50, 0x5a, 0x6d, 0x20, 0x0a, 0x48, 0x7f,
	0xec, 0xf5, 0x0b, 0xc4, 0xb8, 0x53, 0x85, 0x17, 0x6c, 0x2c, 0x75, 0x5e, 0x42, 0x5e, 0x56, 0x65,
	0x34, 0x20, 0xde, 0x2b, 0xcf, 0x9b, 0x6d, 0x71, 0xde, 0x50, 0xc2, 0x29, 0x3b, 0xcc, 0x20, 0x9b,
	0x6b, 0x9d, 0x46, 0x43, 0x62, 0x9f, 0x7a, 0xf6, 0xad, 0x43, 0xf9, 0x56, 0x1d, 0xfe, 0x92, 0xb1,
	0x65, 0xaa, 0xe5, 0xe3, 0x6c, 0x25, 0x54, 0x1e, 0x1d, 0x12, 0xf9, 0xb5, 0x27, 0x5f, 0xd7, 0x0a,
	0xde, 0x22, 0x85, 0x3f, 0x66, 0xc3, 0x67, 0x91, 0xa6, 0x60, 0xa3, 0x11, 0xd1, 0x4f, 0x3c, 0xfd,
	0x3d, 0x81, 0xdc, 0x2b, 0xd1, 0xea, 0xe2, 0xb2, 0x88, 0xc6, 0x1d, 0xab, 0xe7, 0x97, 0x73, 0x8e,
	0x30, 0x6a, 0x4d, 0x21, 0x23, 0xd6, 0xd1, 0xf2, 0x42, 0x72, 0x84, 0xc3, 0x1f, 0xb2, 0x3e, 0xac,
	0x41, 0x46, 0x47, 0xa4, 0x3e, 0xf2, 0xea, 0xaf, 0xd7, 0x20, 0x39, 0x29, 0xc2, 0x88, 0x1d, 0x5a,
	0x28, 0xed, 0x1d, 0xd8, 0xe8, 0xf8, 0xbc, 0x37, 0x1d, 0xf1, 0xad, 0x88, 0x9a, 0x07, 0xb5, 0xbe,
	0x57, 0x19, 0x44, 0x27, 0x4e, 0xe3, 0x45, 0x74, 0x74, 0x51, 0x18, 0xfd, 0x10, 0x9d, 0x76, 0x1c,
	0x3d, 0x47, 0x8c, 0x3b, 0x55, 0xf8, 0x25, 0xeb, 0x8b, 0xca, 0xae, 0xa2, 0xb3, 0x8e, 0x8f, 0xaf,
	0x2a, 0xbb, 0xd2, 0x46, 0xd9, 0x0d, 0x27, 0x6d, 0xf8, 0x23, 0xd6, 0x4f, 0x44, 0x06, 0xd1, 0x2b,
	0x62, 0x9d, 0x79, 0xd6, 0x3b, 0x91, 0x81, 0x7b, 0x22, 0x2a, 0x27, 0xff, 0x3c, 0x60, 0xc1, 0x8d,
	0x4e, 0xc2, 0xcf, 0xd9, 0x28, 0xd5, 0x49, 0x0a, 0x4f, 0x90, 0xfa, 0xb4, 0xa8, 0xe5, 0x70, 0xca,
	0xce, 0x52, 0x9d, 0x60, 0x08, 0x75, 0x0a, 0x37, 0x44, 0x39, 0x20, 0xca, 0x2e, 0x8c, 0x66, 0xa5,
	0x3a, 0x79, 0xab, 0x52, 0x97, 0x27, 0x63, 0xbe, 0x15, 0xc3, 0x73, 0x76, 0x94, 0x89, 0x35, 0xfe,
	0x5c, 0xa8, 0x3f, 0x43, 0xd4, 0x3f, 0xef, 0x4d, 0x4f, 0x78, 0x1b, 0x0a, 0x7f, 0xc0, 0x58, 0x26,
	0xd6, 0xd7, 0x42, 0x3e, 0x56, 0x85, 0x4b, 0x9f, 0x13, 0xde, 0x42, 0xc2, 0xcf, 0xd8, 0x30, 0x13,
	0xeb, 0xab, 0x04, 0x28, 0x59, 0x4e, 0xb8, 0x97, 0xc2, 0x2f, 0xd8, 0x38, 0xd5, 0x52, 0xa4, 0xe4,
	0xcc, 0x43, 0x72, 0x66, 0x03, 0xa0, 0x5d, 0x52, 0x67, 0x85, 0x81, 0xb2, 0xa4, 0x44, 0x18, 0xf1,
	0x5a, 0xc6, 0x1b, 0x25, 0x66, 0x81, 0xa1, 0x07, 0x8f, 0x49, 0xdb, 0x42, 0xc2, 0x9f, 0xb0, 0x53,
	0x2f, 0x55, 0xb9, 0xb4, 0x4a, 0xe7, 0x94, 0x08, 0x23, 0xbe, 0x83, 0x4e, 0xfe, 0xd6, 0x63, 0x87,
	0x3e, 0x65, 0xc3, 0x2f, 0xd9, 0x49, 0xa1, 0x75, 0x3a, 0x13, 0x72, 0xe5, 0x2c, 0x45, 0x67, 0x06,
	0xbc, 0x0b, 0xe2, 0xab, 0x32, 0x95, 0xdf, 0xaf, 0xdf, 0x02, 0x90, 0x2b, 0x03, 0x5e, 0xcb, 0xe8,
	0xa9, 0x07, 0x6d, 0x24, 0x5c, 0x49, 0x09, 0x85, 0x25, 0x3f, 0x8e, 0x78, 0x1b, 0x0a, 0x2f, 0x58,
	0x98, 0x89, 0xf5, 0xfd, 0xfa, 0xae, 0xca, 0xe6, 0x60, 0xae, 0xa4, 0xd4, 0x55, 0x6e, 0xc9, 0xa5,
	0x01, 0xdf, 0xa3, 0x99, 0xfc, 0x65, 0xc8, 0xc6, 0x75, 0x01, 0x86, 0x21, 0xeb, 0xe7, 0x98, 0x16,
	0x2e, 0xca, 0xf4, 0x1b, 0xe3, 0x96, 0x40, 0x0e, 0xa5, 0x2a, 0x7d, 0x64, 0xb7, 0x22, 0x45, 0x45,
	0xe5, 0x60, 0x4a, 0x2b, 0xcc, 0xf6, 0x31, 0x2d, 0x24, 0xfc, 0x19, 0x7b, 0xe5, 0xa9, 0x54, 0x87,
	0x14, 0x04, 0xf7, 0x92, 0x17, 0x38, 0x9e, 0xb5, 0xd2, 0xf6, 0x11, 0x36, 0x57, 0x71, 0x6c, 0x28,
	0xc2, 0x63, 0xde, 0x42, 0x6a, 0xcb, 0x6f, 0x55, 0xae, 0xf2, 0x84, 0xc2, 0xbc, 0xb5, 0xdc, 0x41,
	0x98, 0x03, 0x77, 0x3a, 0x86, 0x6f, 0x62, 0x0a, 0x74, 0xc0, 0xbd, 0x84, 0xfe, 0x9c, 0x03, 0x98,
	0xf2, 0x77, 0xfc, 0x86, 0xa2, 0x3c, 0xe6, 0xb5, 0x8c, 0xb7, 0xce, 0x52, 0x05, 0xb9, 0xa5, 0x5b,
	0xc7, 0xee, 0xd6, 0x06, 0xc1, 0x5b, 0x8d, 0x78, 0xb0, 0x57, 0x85, 0x9a, 0x6b, 0x63, 0xa9, 0xa6,
	0x02, 0xde, 0x86, 0x30, 0xa6, 0xaa, 0xbc, 0x83, 0xe7, 0x6f, 0xb5, 0xca, 0xf1, 0x42, 0xaa, 0xa8,
	0x11, 0xef, 0x82, 0xe8, 0x09, 0x03, 0x22, 0xfe, 0x4d, 0x9e, 0x6e, 0xea, 0xb7, 0xbc, 0xa6, 0xdb,
	0x5e, 0xe0, 0x78, 0xa7, 0x88, 0xe3, 0x9a, 0x16, 0x12, 0xad, 0x0d, 0xe1, 0x69, 0x31, 0x3c, 0x88,
	0x2a, 0xb5, 0x8b, 0x5c, 0x14, 0x33, 0x8a, 0xf0, 0x27, 0xce, 0xaf, 0xbb, 0x78, 0xf8, 0x73, 0xf6,
	0xfa, 0xd9, 0x28, 0x0b, 0xe4, 0xe9, 0x05, 0x48, 0x9d, 0xc7, 0x65, 0xf4, 0x86, 0xc8, 0x2f, 0x15,
	0x68, 0xcd, 0x0a, 0x84, 0xb1, 0x4b, 0x10, 0xf6, 0x5e, 0xc9, 0xc7, 0xe8, 0xd3, 0xf3, 0xde, 0x74,
	0xc0, 0xbb, 0x60, 0x78, 0xc9, 0xde, 0x14, 0xc2, 0x08, 0x0e, 0x99, 0xb6, 0xf0, 0xce, 0x14, 0xd2,
	0x79, 0x2c, 0xfa, 0x8c, 0x9e, 0xba, 0x57, 0x87, 0x56, 0x51, 0x52, 0xfc, 0x1a, 0x54, 0xb2, 0xb2,
	0xd1, 0xf7, 0x9c, 0x27, 0x5b, 0x10, 0x66, 0x2e, 0x64, 0x85, 0xdd, 0xd0, 0x83, 0xbe, 0xc9, 0x2d,
	0x98, 0x27, 0x91, 0x46, 0x91, 0xcb, 0xdc, 0x97, 0x1a, 0xf2, 0x53, 0x65, 0x57, 0xdb, 0x14, 0xff,
	0xbe, 0xf7, 0x53, 0x03, 0xe1, 0x3b, 0x9f, 0x85, 0xb2, 0xf4, 0xb7, 0xf2, 0xab, 0x99, 0xce, 0x32,
	0x65, 0x6f, 0xcb, 0x24, 0xfa, 0x9c, 0x8c, 0xda, 0xab, 0x9b, 0xfc, 0xb7, 0xc7, 0x86, 0x6e, 0x0c,
	0x50, 0x53, 0x51, 0x39, 0x96, 0xa1, 0xab, 0x53, 0x2f, 0x21, 0x1e, 0x1b, 0xf5, 0x04, 0xc6, 0xd7,
	0x83, 0x97, 0x08, 0x5f, 0xce, 0x85, 0x5d, 0xf9, 0xfe, 0xe6, 0x25, 0x2c, 0xa0, 0x78, 0x49, 0xf5,
	0x4d, 0xd9, 0x3f, 0xe0, 0x5b, 0x11, 0x53, 0xb3, 0x54, 0x49, 0x7e, 0xbf, 0x29, 0xc0, 0xa7, 0x7c,
	0x2d, 0x7f, 0x87, 0x84, 0x9f, 0xb0, 0x63, 0x2a, 0xb6, 0x58, 0x95, 0x62, 0x99, 0x6e, 0xfb, 0x5b,
	0x07, 0xc3, 0x36, 0x45, 0xf2, 0xf3, 0x4a, 0x59, 0x48, 0x55, 0x89, 0x13, 0x2f, 0x98, 0x8e, 0xf9,
	0x0e, 0x3a, 0x01, 0x36, 0xa0, 0x71, 0xbd, 0xb7, 0x03, 0x7c, 0x34, 0x83, 0x27, 0xff, 0x09, 0x18,
	0x6b, 0x66, 0x32, 0x5a, 0x10, 0xc3, 0xc3, 0x6e, 0x3f, 0xec, 0x60, 0x98, 0xec, 0x38, 0x09, 0xc0,
	0xca, 0x15, 0xfd, 0xf3, 0xae, 0xca, 0x7c, 0x5b, 0x7c, 0x81, 0xa3, 0xb5, 0x56, 0x65, 0xa0, 0x2b,
	0xbb, 0xcd, 0xf4, 0x80, 0x98, 0x3b, 0x28, 0xa6, 0xf9, 0x52, 0xb4, 0x0f, 0x74, 0x5d, 0xa9, 0x0b,
	0xb6, 0xcc, 0x1e, 0x7c, 0xc0, 0xec, 0xe1, 0x87, 0xcc, 0x3e, 0xec, 0xc6, 0xf9, 0x17, 0xec, 0x13,
	0x55, 0x2e, 0xac, 0xd1, 0x39, 0x8d, 0x44, 0x55, 0x5a, 0xc8, 0xe5, 0xc6, 0xcf, 0x9c, 0x7d, 0x2a,
	0x6c, 0x4c, 0xa5, 0xca, 0x93, 0x14, 0x6e, 0xb1, 0xa7, 0xf8, 0xf1, 0xd3, 0x20, 0x38, 0xd8, 0xe8,
	0xb1, 0xe5, 0x26, 0x97, 0x7e, 0xf2, 0x34, 0x40, 0xf8, 0x15, 0xfb, 0x54, 0x95, 0x1c, 0xa4, 0x36,
	0xb1, 0x2f, 0xef, 0x3f, 0x55, 0x90, 0x4b, 0xa0, 0x6d, 0x64, 0xc4, 0xf7, 0x2b, 0x31, 0xe3, 0x54,
	0x39, 0x17, 0x46, 0xb8, 0x4d, 0xca, 0x6d, 0x25, 0x6d, 0x88, 0x4a, 0x34, 0xc7, 0xbc, 0xba, 0x5f,
	0xff, 0xb6, 0x52, 0x58, 0x8b, 0x31, 0xac, 0xfd, 0x92, 0xb2, 0x47, 0x33, 0xf9, 0x57, 0xc0, 0x82,
	0xf9, 0xe5, 0x9c, 0xf2, 0x1c, 0x20, 0xa6, 0x1e, 0xda, 0x23, 0xd7, 0xd4, 0xf2, 0x47, 0xac, 0xa6,
	0x73, 0x76, 0x94, 0x98, 0x42, 0xde, 0xf8, 0x25, 0xc3, 0x05, 0xad, 0x0d, 0xe1, 0x99, 0xaa, 0x5c,
	0x00, 0xc4, 0xbe, 0x9c, 0xbc, 0x44, 0xcd, 0x09, 0xcc, 0x13, 0x98, 0x05, 0x4d, 0x32, 0x57, 0x48,
	0x6d, 0x08, 0xd7, 0x62, 0x7c, 0x71, 0xe9, 0xcb, 0xc7, 0x09, 0x78, 0x9e, 0xb3, 0xda, 0x47, 0xc8,
	0x4b, 0x54, 0x99, 0x65, 0xd2, 0xe4, 0x35, 0xa3, 0x87, 0x76, 0x30, 0xb4, 0xe3, 0x09, 0x4c, 0x89,
	0x9b, 0xc3, 0x91, 0xb3, 0xc3, 0x8b, 0x78, 0xea, 0x13, 0x98, 0x5b, 0xb5, 0xa6, 0x10, 0x0c, 0xb8,
	0x97, 0xb6, 0xb8, 0x70, 0x1e, 0xf7, 0xb8, 0x58, 0xe3, 0x0a, 0xa6, 0xf2, 0x1c, 0x0c, 0x9a, 0xf2,
	0xb5, 0x7b, 0xce, 0x29, 0x3d, 0x67, 0x17, 0xa6, 0x08, 0x23, 0x74, 0xad, 0x2b, 0x2c, 0x8e, 0x33,
	0x3a, 0xa6, 0x0d, 0x61, 0x5e, 0x55, 0x25, 0xbc, 0x53, 0x76, 0x55, 0x2d, 0xfd, 0x28, 0x6b, 0x80,
	0xc9, 0xbf, 0x0f, 0x58, 0xc0, 0x0b, 0x89, 0xf6, 0xfd, 0xd1, 0x14, 0xf2, 0x5a, 0xe5, 0x31, 0x0d,
	0x4e, 0xd7, 0x2c, 0x3a, 0x18, 0x72, 0x92, 0x36, 0xc7, 0x45, 0xb7, 0x83, 0x61, 0x5e, 0x60, 0x0b,
	0xa2, 0xbe, 0x14, 0x90, 0x63, 0x6b, 0x19, 0x5f, 0xd2, 0x34, 0xad, 0x3e, 0x29, 0x1b, 0x00, 0xc7,
	0x1a, 0xde, 0x86, 0x6b, 0xd6, 0xfb, 0x9a, 0x35, 0x20, 0xd6, 0x4b, 0x05, 0xb2, 0x93, 0x17, 0xec,
	0xa1, 0x63, 0x27, 0xfb, 0xd8, 0xdb, 0x23, 0xae, 0x53, 0x21, 0x1f, 0x89, 0x7d, 0xd8, 0x3d, 0xbb,
	0x56, 0xb4, 0xcf, 0x6e, 0xd8, 0xa3, 0xee, 0xd9, 0x0d, 0x7b, 0xca, 0xce, 0x32, 0x81, 0x0e, 0xb7,
	0xdf, 0x9a, 0x42, 0xb6, 0xb6, 0x8e, 0x5d, 0x78, 0xf2, 0xf7, 0x1e, 0xeb, 0xe3, 0x2e, 0xee, 0xb7,
	0x2c, 0xfc, 0xd9, 0x8c, 0xa2, 0x16, 0xe2, 0x92, 0xfa, 0xad, 0xf1, 0xdb, 0x22, 0x25, 0x35, 0x4a,
	0xf8, 0x3f, 0x97, 0x8e, 0x0b, 0x2b, 0xea, 0xed, 0xac, 0x41, 0x1a, 0xfd, 0xed, 0xef, 0x67, 0x33,
	0xaa, 0xa5, 0x5a, 0x8f, 0x48, 0xfd, 0x89, 0x30, 0xf8, 0x7f, 0x9f, 0x08, 0x7f, 0x60, 0xe3, 0xfa,
	0xd3, 0xa2, 0x55, 0x0e, 0xbd, 0x4e, 0x39, 0xe0, 0x2e, 0x6d, 0x36, 0x85, 0xd5, 0x54, 0xce, 0x2e,
	0x11, 0x5a, 0x48, 0x67, 0x0c, 0x06, 0xdd, 0x31, 0x38, 0xf9, 0x29, 0x1b, 0xd0, 0xe7, 0x0d, 0x1e,
	0x92, 0x52, 0x77, 0x6c, 0x65, 0x5c, 0x0b, 0x99, 0xfc, 0xa3, 0xc7, 0x46, 0xdb, 0xc7, 0x21, 0x59,
	0x48, 0xab, 0x9e, 0x80, 0x76, 0x4e, 0xd7, 0x72, 0x5a, 0x88, 0x1f, 0x3c, 0xb8, 0x29, 0xb9, 0xed,
	0xe9, 0xc0, 0x15, 0x68, 0x1b, 0xa3, 0x3d, 0x5c, 0xac, 0x9d, 0x3e, 0x70, 0x4d, 0x6b, 0x2b, 0xe3,
	0x00, 0xc9, 0xc4, 0x1a, 0xaf, 0xbb, 0xca, 0x5a, 0x0b, 0x76, 0x17, 0x24, 0x96, 0xca, 0x5b, 0xac,
	0x81, 0x67, 0xb5, 0xc1, 0xe5, 0x90, 0x3e, 0xc5, 0x7f, 0xf5, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xc4, 0x16, 0x14, 0xb5, 0x9a, 0x0f, 0x00, 0x00,
}
