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
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

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
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

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
	PoolCacheSize int64 `protobuf:"varint,1,opt,name=poolCacheSize" json:"poolCacheSize,omitempty"`
	MinTxFee      int64 `protobuf:"varint,2,opt,name=minTxFee" json:"minTxFee,omitempty"`
	ForceAccept   bool  `protobuf:"varint,3,opt,name=forceAccept" json:"forceAccept,omitempty"`
}

func (m *MemPool) Reset()                    { *m = MemPool{} }
func (m *MemPool) String() string            { return proto.CompactTextString(m) }
func (*MemPool) ProtoMessage()               {}
func (*MemPool) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

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

type Consensus struct {
	Name             string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Genesis          string `protobuf:"bytes,2,opt,name=genesis" json:"genesis,omitempty"`
	Minerstart       bool   `protobuf:"varint,3,opt,name=minerstart" json:"minerstart,omitempty"`
	GenesisBlockTime int64  `protobuf:"varint,4,opt,name=genesisBlockTime" json:"genesisBlockTime,omitempty"`
	HotkeyAddr       string `protobuf:"bytes,5,opt,name=hotkeyAddr" json:"hotkeyAddr,omitempty"`
	ForceMining      bool   `protobuf:"varint,6,opt,name=forceMining" json:"forceMining,omitempty"`
	NodeId           int64  `protobuf:"varint,7,opt,name=NodeId" json:"NodeId,omitempty"`
	PeersURL         string `protobuf:"bytes,8,opt,name=PeersURL" json:"PeersURL,omitempty"`
	ClientAddr       string `protobuf:"bytes,9,opt,name=ClientAddr" json:"ClientAddr,omitempty"`
	RaftApiPort      int64  `protobuf:"varint,15,opt,name=raftApiPort" json:"raftApiPort,omitempty"`
	IsNewJoinNode    bool   `protobuf:"varint,16,opt,name=isNewJoinNode" json:"isNewJoinNode,omitempty"`
	ReadOnlyPeersURL string `protobuf:"bytes,17,opt,name=readOnlyPeersURL" json:"readOnlyPeersURL,omitempty"`
	AddPeersURL      string `protobuf:"bytes,18,opt,name=addPeersURL" json:"addPeersURL,omitempty"`
	DefaultSnapCount int64  `protobuf:"varint,19,opt,name=defaultSnapCount" json:"defaultSnapCount,omitempty"`
}

func (m *Consensus) Reset()                    { *m = Consensus{} }
func (m *Consensus) String() string            { return proto.CompactTextString(m) }
func (*Consensus) ProtoMessage()               {}
func (*Consensus) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

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

type Wallet struct {
	MinFee       int64  `protobuf:"varint,1,opt,name=minFee" json:"minFee,omitempty"`
	Driver       string `protobuf:"bytes,2,opt,name=driver" json:"driver,omitempty"`
	DbPath       string `protobuf:"bytes,3,opt,name=dbPath" json:"dbPath,omitempty"`
	DbCache      int32  `protobuf:"varint,4,opt,name=dbCache" json:"dbCache,omitempty"`
	SignType     string `protobuf:"bytes,5,opt,name=signType" json:"signType,omitempty"`
	ForceMining  bool   `protobuf:"varint,6,opt,name=forceMining" json:"forceMining,omitempty"`
	Minerdisable bool   `protobuf:"varint,7,opt,name=minerdisable" json:"minerdisable,omitempty"`
}

func (m *Wallet) Reset()                    { *m = Wallet{} }
func (m *Wallet) String() string            { return proto.CompactTextString(m) }
func (*Wallet) ProtoMessage()               {}
func (*Wallet) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

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

type Store struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Driver  string `protobuf:"bytes,2,opt,name=driver" json:"driver,omitempty"`
	DbPath  string `protobuf:"bytes,3,opt,name=dbPath" json:"dbPath,omitempty"`
	DbCache int32  `protobuf:"varint,4,opt,name=dbCache" json:"dbCache,omitempty"`
}

func (m *Store) Reset()                    { *m = Store{} }
func (m *Store) String() string            { return proto.CompactTextString(m) }
func (*Store) ProtoMessage()               {}
func (*Store) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

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
	DefCacheSize        int64  `protobuf:"varint,1,opt,name=defCacheSize" json:"defCacheSize,omitempty"`
	MaxFetchBlockNum    int64  `protobuf:"varint,2,opt,name=maxFetchBlockNum" json:"maxFetchBlockNum,omitempty"`
	TimeoutSeconds      int64  `protobuf:"varint,3,opt,name=timeoutSeconds" json:"timeoutSeconds,omitempty"`
	BatchBlockNum       int64  `protobuf:"varint,4,opt,name=batchBlockNum" json:"batchBlockNum,omitempty"`
	Driver              string `protobuf:"bytes,5,opt,name=driver" json:"driver,omitempty"`
	DbPath              string `protobuf:"bytes,6,opt,name=dbPath" json:"dbPath,omitempty"`
	DbCache             int32  `protobuf:"varint,7,opt,name=dbCache" json:"dbCache,omitempty"`
	IsStrongConsistency bool   `protobuf:"varint,8,opt,name=isStrongConsistency" json:"isStrongConsistency,omitempty"`
	SingleMode          bool   `protobuf:"varint,9,opt,name=singleMode" json:"singleMode,omitempty"`
	Batchsync           bool   `protobuf:"varint,10,opt,name=batchsync" json:"batchsync,omitempty"`
}

func (m *BlockChain) Reset()                    { *m = BlockChain{} }
func (m *BlockChain) String() string            { return proto.CompactTextString(m) }
func (*BlockChain) ProtoMessage()               {}
func (*BlockChain) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

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
}

func (m *P2P) Reset()                    { *m = P2P{} }
func (m *P2P) String() string            { return proto.CompactTextString(m) }
func (*P2P) ProtoMessage()               {}
func (*P2P) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

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

type Rpc struct {
	JrpcBindAddr string   `protobuf:"bytes,1,opt,name=jrpcBindAddr" json:"jrpcBindAddr,omitempty"`
	GrpcBindAddr string   `protobuf:"bytes,2,opt,name=grpcBindAddr" json:"grpcBindAddr,omitempty"`
	Whitlist     []string `protobuf:"bytes,3,rep,name=whitlist" json:"whitlist,omitempty"`
}

func (m *Rpc) Reset()                    { *m = Rpc{} }
func (m *Rpc) String() string            { return proto.CompactTextString(m) }
func (*Rpc) ProtoMessage()               {}
func (*Rpc) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

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

type Exec struct {
	MinExecFee int64 `protobuf:"varint,1,opt,name=minExecFee" json:"minExecFee,omitempty"`
	IsFree     bool  `protobuf:"varint,2,opt,name=isFree" json:"isFree,omitempty"`
}

func (m *Exec) Reset()                    { *m = Exec{} }
func (m *Exec) String() string            { return proto.CompactTextString(m) }
func (*Exec) ProtoMessage()               {}
func (*Exec) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{9} }

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
}

func init() { proto.RegisterFile("config.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 1066 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xcd, 0x6e, 0x2b, 0x35,
	0x14, 0xc7, 0x95, 0x4e, 0x93, 0x26, 0x4e, 0x73, 0x3f, 0xe6, 0x22, 0x34, 0x42, 0x57, 0x50, 0x8d,
	0x00, 0x55, 0x2c, 0x2a, 0x28, 0x7b, 0xa4, 0x36, 0xba, 0x95, 0x40, 0x6d, 0x89, 0x9c, 0x22, 0xd6,
	0x13, 0xcf, 0xe9, 0xc4, 0xd4, 0x63, 0x8f, 0x6c, 0xa7, 0x4d, 0x78, 0x29, 0x9e, 0x81, 0x3d, 0xaf,
	0xc0, 0x8e, 0x17, 0xe0, 0x0d, 0xd0, 0x39, 0xe3, 0x4c, 0x66, 0xd2, 0x7b, 0x25, 0x16, 0x77, 0x97,
	0xff, 0xef, 0x9c, 0xd8, 0x3e, 0x1f, 0x3e, 0x1e, 0x76, 0x2c, 0x8c, 0xbe, 0x97, 0xc5, 0x59, 0x65,
	0x8d, 0x37, 0x71, 0xdf, 0x6f, 0x2a, 0x70, 0xe9, 0xbf, 0x07, 0x6c, 0x30, 0x25, 0x1e, 0x7f, 0xc2,
	0xfa, 0x5e, 0x7a, 0x05, 0x49, 0xef, 0xa4, 0x77, 0x3a, 0xe2, 0xb5, 0x88, 0xdf, 0xb2, 0x48, 0x99,
	0x22, 0x39, 0x38, 0xe9, 0x9d, 0x8e, 0xcf, 0xd9, 0x19, 0xfd, 0xeb, 0xec, 0xda, 0x14, 0x1c, 0x71,
	0x9c, 0xb2, 0xbe, 0xf3, 0xc6, 0x42, 0x12, 0x91, 0xfd, 0x38, 0xd8, 0xe7, 0xc8, 0x78, 0x6d, 0x8a,
	0xcf, 0xd8, 0x48, 0x18, 0xed, 0x40, 0xbb, 0x95, 0x4b, 0xfa, 0xe4, 0xf7, 0x2a, 0xf8, 0x4d, 0xb7,
	0x9c, 0xef, 0x5c, 0xe2, 0x53, 0x76, 0x54, 0x42, 0x39, 0x33, 0x46, 0x25, 0x03, 0xf2, 0x7e, 0x11,
	0xbc, 0x6f, 0x6a, 0xca, 0xb7, 0xe6, 0xf8, 0x3b, 0xc6, 0x16, 0xca, 0x88, 0x87, 0xe9, 0x32, 0x93,
	0x3a, 0x39, 0x22, 0xe7, 0xd7, 0xc1, 0xf9, 0xb2, 0x31, 0xf0, 0x96, 0x53, 0xfc, 0x15, 0x1b, 0x3c,
	0x65, 0x4a, 0x81, 0x4f, 0x86, 0xe4, 0x3e, 0x09, 0xee, 0xbf, 0x12, 0xe4, 0xc1, 0x88, 0x51, 0x57,
	0xe7, 0x55, 0x32, 0xea, 0x44, 0x3d, 0x3b, 0x9f, 0x71, 0xc4, 0x68, 0xb5, 0x95, 0x48, 0x58, 0xc7,
	0xca, 0x2b, 0xc1, 0x11, 0xc7, 0x5f, 0xb0, 0x43, 0x58, 0x83, 0x48, 0xc6, 0x64, 0x1e, 0x07, 0xf3,
	0xbb, 0x35, 0x08, 0x4e, 0x86, 0xf4, 0xcf, 0x03, 0x16, 0x5d, 0x9b, 0x22, 0xfe, 0x8c, 0x0d, 0x95,
	0x29, 0x14, 0x3c, 0x82, 0x0a, 0x39, 0x6f, 0x74, 0x7c, 0xca, 0x5e, 0x2a, 0x53, 0x60, 0x7e, 0x8c,
	0x82, 0x6b, 0x72, 0x39, 0x20, 0x97, 0x7d, 0x1c, 0x27, 0xec, 0x48, 0x99, 0xe2, 0x4a, 0xaa, 0xba,
	0x08, 0x23, 0xbe, 0x95, 0xf1, 0x09, 0x1b, 0x97, 0xd9, 0x1a, 0x7f, 0xce, 0xe5, 0xef, 0x90, 0x1c,
	0x9e, 0xf4, 0x4e, 0x27, 0xbc, 0x8d, 0xe2, 0xcf, 0x19, 0x2b, 0xb3, 0xf5, 0x65, 0x26, 0x1e, 0x56,
	0x55, 0x5d, 0x9b, 0x09, 0x6f, 0x91, 0xf8, 0x53, 0x36, 0x28, 0xb3, 0xf5, 0x45, 0x01, 0x54, 0x89,
	0x09, 0x0f, 0x2a, 0x7e, 0xcb, 0x46, 0xca, 0x88, 0x4c, 0xdd, 0xc9, 0x12, 0x28, 0xef, 0x43, 0xbe,
	0x03, 0x18, 0x97, 0x30, 0x65, 0x65, 0xc1, 0x39, 0xca, 0xf2, 0x90, 0x37, 0x1a, 0x77, 0x14, 0x98,
	0x62, 0x4b, 0x07, 0x1e, 0x91, 0xb5, 0x45, 0xe2, 0xaf, 0xd9, 0x8b, 0xa0, 0x56, 0x5a, 0x78, 0x69,
	0x34, 0x65, 0x79, 0xc8, 0xf7, 0x68, 0x5a, 0xb2, 0xa3, 0xd0, 0x0e, 0xf1, 0x97, 0x6c, 0x52, 0x19,
	0xa3, 0xa6, 0x99, 0x58, 0xd6, 0x81, 0x62, 0x2e, 0x23, 0xde, 0x85, 0x78, 0xa8, 0x52, 0xea, 0xbb,
	0xf5, 0x15, 0x00, 0x65, 0x32, 0xe2, 0x8d, 0xc6, 0x44, 0xdd, 0x1b, 0x2b, 0xe0, 0x42, 0x08, 0xa8,
	0x3c, 0xa5, 0x71, 0xc8, 0xdb, 0x28, 0xfd, 0x3b, 0x62, 0xa3, 0xa6, 0x59, 0xe3, 0x98, 0x1d, 0xea,
	0xac, 0xdc, 0x5e, 0x14, 0xfa, 0x8d, 0x65, 0x28, 0x40, 0x83, 0x93, 0x2e, 0x14, 0x6a, 0x2b, 0x29,
	0xc9, 0x52, 0x83, 0x75, 0x3e, 0xb3, 0xdb, 0xc5, 0x5b, 0x24, 0xfe, 0x86, 0xbd, 0x0a, 0xae, 0xd4,
	0xb3, 0x94, 0xd3, 0x43, 0x3a, 0xe1, 0x33, 0x8e, 0x6b, 0x2d, 0x8d, 0x7f, 0x80, 0xcd, 0x45, 0x9e,
	0x5b, 0x2a, 0xd8, 0x88, 0xb7, 0x48, 0x13, 0xc9, 0x8d, 0xd4, 0x52, 0x17, 0x54, 0xb5, 0x6d, 0x24,
	0x35, 0xc2, 0x92, 0xde, 0x9a, 0x1c, 0x7e, 0xcc, 0xa9, 0x6e, 0x11, 0x0f, 0x0a, 0xf3, 0x33, 0x03,
	0xb0, 0xee, 0x17, 0x7e, 0x4d, 0x45, 0x1b, 0xf1, 0x46, 0xe3, 0xae, 0x53, 0x25, 0x41, 0x7b, 0xda,
	0x75, 0x54, 0xef, 0xba, 0x23, 0xb8, 0xab, 0xcd, 0xee, 0xfd, 0x45, 0x25, 0x67, 0xc6, 0xfa, 0xe4,
	0x25, 0x2d, 0xdc, 0x46, 0x58, 0x23, 0xe9, 0x6e, 0xe1, 0xe9, 0x27, 0x23, 0x35, 0x6e, 0x98, 0xbc,
	0xa2, 0x93, 0x75, 0x21, 0x66, 0xc2, 0x42, 0x96, 0xff, 0xac, 0xd5, 0xa6, 0x39, 0xcb, 0x6b, 0xda,
	0xed, 0x19, 0xc7, 0x3d, 0xb3, 0x3c, 0x6f, 0xdc, 0x62, 0x72, 0x6b, 0x23, 0x5c, 0x2d, 0x87, 0xfb,
	0x6c, 0xa5, 0xfc, 0x5c, 0x67, 0xd5, 0xd4, 0xac, 0xb4, 0x4f, 0xde, 0xd4, 0x79, 0xdd, 0xe7, 0xe9,
	0x5f, 0x3d, 0x36, 0xa8, 0x47, 0x00, 0xf5, 0xbc, 0xd4, 0xd8, 0x26, 0x75, 0x1f, 0x05, 0x85, 0x3c,
	0xb7, 0xf2, 0x11, 0x6c, 0xa8, 0x6f, 0x50, 0xc4, 0x17, 0xb3, 0xcc, 0x2f, 0xc3, 0xf5, 0x0b, 0x0a,
	0x1b, 0x22, 0x5f, 0x50, 0xff, 0x51, 0x35, 0xfb, 0x7c, 0x2b, 0x31, 0xd5, 0x4e, 0x16, 0xfa, 0x6e,
	0x53, 0x41, 0x28, 0x61, 0xa3, 0xff, 0x47, 0x01, 0x53, 0x76, 0x4c, 0xcd, 0x93, 0x4b, 0x97, 0x2d,
	0xd4, 0xf6, 0xfa, 0x75, 0x58, 0x0a, 0xac, 0x4f, 0x23, 0xf8, 0xbd, 0x9d, 0xfa, 0xd1, 0x02, 0x49,
	0xff, 0x39, 0x60, 0x6c, 0x37, 0x67, 0xf1, 0x64, 0x39, 0xdc, 0xef, 0xdf, 0xc3, 0x0e, 0xc3, 0xa2,
	0xe0, 0x00, 0x02, 0x2f, 0x96, 0xf4, 0xcf, 0xdb, 0x55, 0x19, 0xae, 0xe3, 0x33, 0x8e, 0xb3, 0xc0,
	0xcb, 0x12, 0xcc, 0xca, 0xcf, 0x41, 0x18, 0x9d, 0x3b, 0x3a, 0x58, 0xc4, 0xf7, 0x28, 0x36, 0xd7,
	0x22, 0x6b, 0x2f, 0x58, 0xdf, 0x9e, 0x2e, 0x6c, 0x85, 0xdd, 0xff, 0x40, 0xd8, 0x83, 0x0f, 0x85,
	0x7d, 0xd4, 0xad, 0xdf, 0xb7, 0xec, 0x8d, 0x74, 0x73, 0x6f, 0x8d, 0xa6, 0x49, 0x2c, 0x9d, 0x07,
	0x2d, 0x36, 0x61, 0xd4, 0xbd, 0xcf, 0x84, 0x17, 0xc8, 0x49, 0x5d, 0x28, 0xb8, 0xc1, 0xde, 0x0f,
	0x53, 0x6f, 0x47, 0x70, 0x9e, 0xd2, 0x61, 0xdd, 0x46, 0x8b, 0x30, 0xf0, 0x76, 0x20, 0xfd, 0x23,
	0x62, 0xd1, 0xec, 0x7c, 0x46, 0x7d, 0x03, 0x90, 0xd3, 0x1d, 0xeb, 0xd1, 0x91, 0x1a, 0xfd, 0x11,
	0xbb, 0xf3, 0x84, 0x8d, 0x0b, 0x5b, 0x89, 0xeb, 0xf0, 0xa6, 0xd4, 0xc9, 0x6a, 0x23, 0x5c, 0x53,
	0xba, 0x39, 0x40, 0x1e, 0xda, 0x33, 0x28, 0xfc, 0xa7, 0x03, 0xfb, 0x08, 0x76, 0x4e, 0x93, 0xae,
	0x6e, 0xcc, 0x36, 0xc2, 0x4f, 0x0c, 0x3c, 0x31, 0x3e, 0x0b, 0x11, 0x7e, 0x62, 0x90, 0xc0, 0xf5,
	0x40, 0x53, 0x2f, 0xd7, 0x99, 0x09, 0x8a, 0x3a, 0xdd, 0x15, 0xbb, 0x7e, 0x62, 0x74, 0xd0, 0x0e,
	0xc3, 0x38, 0x1e, 0xc1, 0x3a, 0x7c, 0x28, 0xc6, 0x75, 0x1c, 0x41, 0xe2, 0xaa, 0x8f, 0x60, 0x6f,
	0xe4, 0x3a, 0x39, 0x26, 0x43, 0x50, 0x5b, 0x9e, 0xad, 0x93, 0xc9, 0x8e, 0x67, 0x6b, 0x7c, 0x71,
	0xa5, 0xd6, 0x60, 0x31, 0x94, 0x77, 0xf5, 0x71, 0x5e, 0xd0, 0x71, 0xf6, 0x31, 0xc6, 0x49, 0xe8,
	0xd2, 0xac, 0xb0, 0x29, 0x5f, 0xd2, 0x32, 0x6d, 0x94, 0x4a, 0x16, 0xf1, 0x4a, 0x60, 0x00, 0xbf,
	0xd9, 0x4a, 0x5c, 0x4a, 0x9d, 0xd3, 0xe4, 0xac, 0x6f, 0x61, 0x87, 0xa1, 0x4f, 0xd1, 0xf6, 0xa9,
	0xcb, 0xd7, 0x61, 0x58, 0xf8, 0xa7, 0xa5, 0xf4, 0x4a, 0x3a, 0x7c, 0x3f, 0x30, 0x73, 0x8d, 0x4e,
	0x7f, 0x60, 0x87, 0xf8, 0x69, 0x11, 0x5e, 0x19, 0xfc, 0xb9, 0x1b, 0x5d, 0x2d, 0x52, 0x17, 0xed,
	0xca, 0x86, 0xd7, 0x8f, 0x8a, 0x86, 0x6a, 0x31, 0xa0, 0xcf, 0xc1, 0xef, 0xff, 0x0b, 0x00, 0x00,
	0xff, 0xff, 0x98, 0x7f, 0xd8, 0xd3, 0x1e, 0x0a, 0x00, 0x00,
}
