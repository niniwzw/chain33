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
	Title           string      `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Loglevel        string      `protobuf:"bytes,2,opt,name=loglevel" json:"loglevel,omitempty"`
	LogConsoleLevel string      `protobuf:"bytes,10,opt,name=logConsoleLevel" json:"logConsoleLevel,omitempty"`
	LogFile         string      `protobuf:"bytes,9,opt,name=logFile" json:"logFile,omitempty"`
	Store           *Store      `protobuf:"bytes,3,opt,name=store" json:"store,omitempty"`
	LocalStore      *LocalStore `protobuf:"bytes,11,opt,name=localStore" json:"localStore,omitempty"`
	Consensus       *Consensus  `protobuf:"bytes,4,opt,name=consensus" json:"consensus,omitempty"`
	MemPool         *MemPool    `protobuf:"bytes,5,opt,name=memPool" json:"memPool,omitempty"`
	BlockChain      *BlockChain `protobuf:"bytes,6,opt,name=blockChain" json:"blockChain,omitempty"`
	Wallet          *Wallet     `protobuf:"bytes,7,opt,name=wallet" json:"wallet,omitempty"`
	P2P             *P2P        `protobuf:"bytes,8,opt,name=p2p" json:"p2p,omitempty"`
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

func (m *Config) GetLoglevel() string {
	if m != nil {
		return m.Loglevel
	}
	return ""
}

func (m *Config) GetLogConsoleLevel() string {
	if m != nil {
		return m.LogConsoleLevel
	}
	return ""
}

func (m *Config) GetLogFile() string {
	if m != nil {
		return m.LogFile
	}
	return ""
}

func (m *Config) GetStore() *Store {
	if m != nil {
		return m.Store
	}
	return nil
}

func (m *Config) GetLocalStore() *LocalStore {
	if m != nil {
		return m.LocalStore
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

type MemPool struct {
	PoolCacheSize int64 `protobuf:"varint,1,opt,name=poolCacheSize" json:"poolCacheSize,omitempty"`
	MinTxFee      int64 `protobuf:"varint,2,opt,name=minTxFee" json:"minTxFee,omitempty"`
}

func (m *MemPool) Reset()                    { *m = MemPool{} }
func (m *MemPool) String() string            { return proto.CompactTextString(m) }
func (*MemPool) ProtoMessage()               {}
func (*MemPool) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

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

type Consensus struct {
	Name             string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Genesis          string `protobuf:"bytes,2,opt,name=genesis" json:"genesis,omitempty"`
	Minerstart       bool   `protobuf:"varint,3,opt,name=minerstart" json:"minerstart,omitempty"`
	NodeId           int64  `protobuf:"varint,4,opt,name=nodeId" json:"nodeId,omitempty"`
	RaftApiPort      int64  `protobuf:"varint,5,opt,name=raftApiPort" json:"raftApiPort,omitempty"`
	IsNewJoinNode    bool   `protobuf:"varint,6,opt,name=isNewJoinNode" json:"isNewJoinNode,omitempty"`
	PeersURL         string `protobuf:"bytes,7,opt,name=peersURL" json:"peersURL,omitempty"`
	ReadOnlyPeersURL string `protobuf:"bytes,8,opt,name=readOnlyPeersURL" json:"readOnlyPeersURL,omitempty"`
	AddPeersURL      string `protobuf:"bytes,9,opt,name=addPeersURL" json:"addPeersURL,omitempty"`
	GenesisBlockTime int64  `protobuf:"varint,10,opt,name=genesisBlockTime" json:"genesisBlockTime,omitempty"`
	HotkeyAddr       string `protobuf:"bytes,11,opt,name=hotkeyAddr" json:"hotkeyAddr,omitempty"`
}

func (m *Consensus) Reset()                    { *m = Consensus{} }
func (m *Consensus) String() string            { return proto.CompactTextString(m) }
func (*Consensus) ProtoMessage()               {}
func (*Consensus) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

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

func (m *Consensus) GetNodeId() int64 {
	if m != nil {
		return m.NodeId
	}
	return 0
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

func (m *Consensus) GetPeersURL() string {
	if m != nil {
		return m.PeersURL
	}
	return ""
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

type Wallet struct {
	MinFee int64  `protobuf:"varint,1,opt,name=minFee" json:"minFee,omitempty"`
	DbPath string `protobuf:"bytes,2,opt,name=dbPath" json:"dbPath,omitempty"`
}

func (m *Wallet) Reset()                    { *m = Wallet{} }
func (m *Wallet) String() string            { return proto.CompactTextString(m) }
func (*Wallet) ProtoMessage()               {}
func (*Wallet) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *Wallet) GetMinFee() int64 {
	if m != nil {
		return m.MinFee
	}
	return 0
}

func (m *Wallet) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
}

type Store struct {
	Driver string `protobuf:"bytes,1,opt,name=driver" json:"driver,omitempty"`
	DbPath string `protobuf:"bytes,2,opt,name=dbPath" json:"dbPath,omitempty"`
}

func (m *Store) Reset()                    { *m = Store{} }
func (m *Store) String() string            { return proto.CompactTextString(m) }
func (*Store) ProtoMessage()               {}
func (*Store) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

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

type LocalStore struct {
	Driver string `protobuf:"bytes,1,opt,name=driver" json:"driver,omitempty"`
	DbPath string `protobuf:"bytes,2,opt,name=dbPath" json:"dbPath,omitempty"`
}

func (m *LocalStore) Reset()                    { *m = LocalStore{} }
func (m *LocalStore) String() string            { return proto.CompactTextString(m) }
func (*LocalStore) ProtoMessage()               {}
func (*LocalStore) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *LocalStore) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *LocalStore) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
}

type BlockChain struct {
	DefCacheSize     int64  `protobuf:"varint,1,opt,name=defCacheSize" json:"defCacheSize,omitempty"`
	MaxFetchBlockNum int64  `protobuf:"varint,2,opt,name=maxFetchBlockNum" json:"maxFetchBlockNum,omitempty"`
	TimeoutSeconds   int64  `protobuf:"varint,3,opt,name=timeoutSeconds" json:"timeoutSeconds,omitempty"`
	BatchBlockNum    int64  `protobuf:"varint,4,opt,name=batchBlockNum" json:"batchBlockNum,omitempty"`
	Driver           string `protobuf:"bytes,5,opt,name=driver" json:"driver,omitempty"`
	DbPath           string `protobuf:"bytes,6,opt,name=dbPath" json:"dbPath,omitempty"`
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

type P2P struct {
	SeedPort     int32    `protobuf:"varint,1,opt,name=seedPort" json:"seedPort,omitempty"`
	DbPath       string   `protobuf:"bytes,2,opt,name=dbPath" json:"dbPath,omitempty"`
	GrpcLogFile  string   `protobuf:"bytes,3,opt,name=grpcLogFile" json:"grpcLogFile,omitempty"`
	IsSeed       bool     `protobuf:"varint,4,opt,name=isSeed" json:"isSeed,omitempty"`
	Seeds        []string `protobuf:"bytes,5,rep,name=seeds" json:"seeds,omitempty"`
	Enable       bool     `protobuf:"varint,6,opt,name=enable" json:"enable,omitempty"`
	MsgCacheSize int32    `protobuf:"varint,7,opt,name=msgCacheSize" json:"msgCacheSize,omitempty"`
	Version      int32    `protobuf:"varint,8,opt,name=version" json:"version,omitempty"`
	VerMix       int32    `protobuf:"varint,9,opt,name=verMix" json:"verMix,omitempty"`
	VerMax       int32    `protobuf:"varint,10,opt,name=verMax" json:"verMax,omitempty"`
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

func (m *P2P) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
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

func init() {
	proto.RegisterType((*Config)(nil), "types.Config")
	proto.RegisterType((*MemPool)(nil), "types.MemPool")
	proto.RegisterType((*Consensus)(nil), "types.Consensus")
	proto.RegisterType((*Wallet)(nil), "types.Wallet")
	proto.RegisterType((*Store)(nil), "types.Store")
	proto.RegisterType((*LocalStore)(nil), "types.LocalStore")
	proto.RegisterType((*BlockChain)(nil), "types.BlockChain")
	proto.RegisterType((*P2P)(nil), "types.P2P")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 712 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0xcf, 0x6e, 0xdb, 0x38,
	0x10, 0xc6, 0xe1, 0x28, 0xf2, 0x9f, 0xc9, 0x9f, 0xcd, 0x12, 0x8b, 0x85, 0xb0, 0x58, 0x2c, 0x0c,
	0x63, 0x77, 0x61, 0xec, 0xc1, 0xc0, 0xba, 0x87, 0xf6, 0xd0, 0x4b, 0x6a, 0x20, 0x40, 0x5b, 0x27,
	0x15, 0xe8, 0x14, 0x3d, 0xcb, 0xd2, 0xc4, 0x26, 0x42, 0x91, 0x82, 0xc8, 0x38, 0x4e, 0x5f, 0xa2,
	0x8f, 0xd1, 0x77, 0xea, 0x4b, 0xf4, 0x15, 0x0a, 0x8e, 0x28, 0x59, 0x4e, 0x90, 0x43, 0x6f, 0xfe,
	0x7e, 0xfc, 0x48, 0x70, 0x66, 0x3e, 0xd1, 0x70, 0x9c, 0x6a, 0x75, 0x23, 0x56, 0x93, 0xa2, 0xd4,
	0x56, 0xb3, 0xd0, 0x3e, 0x14, 0x68, 0x46, 0x5f, 0x03, 0xe8, 0xce, 0x88, 0xb3, 0xdf, 0x20, 0xb4,
	0xc2, 0x4a, 0x8c, 0x3a, 0xc3, 0xce, 0x78, 0xc0, 0x2b, 0xc1, 0xfe, 0x80, 0xbe, 0xd4, 0x2b, 0x89,
	0x1b, 0x94, 0xd1, 0x01, 0x2d, 0x34, 0x9a, 0x8d, 0xe1, 0x17, 0xa9, 0x57, 0x33, 0xad, 0x8c, 0x96,
	0x38, 0x27, 0x0b, 0x90, 0xe5, 0x31, 0x66, 0x11, 0xf4, 0xa4, 0x5e, 0x5d, 0x08, 0x89, 0xd1, 0x80,
	0x1c, 0xb5, 0x64, 0x23, 0x08, 0x8d, 0xd5, 0x25, 0x46, 0xc1, 0xb0, 0x33, 0x3e, 0x9a, 0x1e, 0x4f,
	0xe8, 0x5e, 0x93, 0x85, 0x63, 0xbc, 0x5a, 0x62, 0xff, 0x03, 0x48, 0x9d, 0x26, 0x92, 0x60, 0x74,
	0x44, 0xc6, 0x5f, 0xbd, 0x71, 0xde, 0x2c, 0xf0, 0x96, 0x89, 0x4d, 0x60, 0x90, 0x6a, 0x65, 0x50,
	0x99, 0x3b, 0x13, 0x1d, 0xd2, 0x8e, 0x33, 0xbf, 0x63, 0x56, 0x73, 0xbe, 0xb3, 0xb0, 0x31, 0xf4,
	0x72, 0xcc, 0x63, 0xad, 0x65, 0x14, 0x92, 0xfb, 0xd4, 0xbb, 0x2f, 0x2b, 0xca, 0xeb, 0x65, 0x77,
	0x99, 0xa5, 0xd4, 0xe9, 0xed, 0x6c, 0x9d, 0x08, 0x15, 0x75, 0xf7, 0x2e, 0xf3, 0xa6, 0x59, 0xe0,
	0x2d, 0x13, 0xfb, 0x07, 0xba, 0xf7, 0x89, 0x94, 0x68, 0xa3, 0x1e, 0xd9, 0x4f, 0xbc, 0xfd, 0x13,
	0x41, 0xee, 0x17, 0xd9, 0x9f, 0x10, 0x14, 0xd3, 0x22, 0xea, 0x93, 0x07, 0xbc, 0x27, 0x9e, 0xc6,
	0xdc, 0xe1, 0xd1, 0x7b, 0xe8, 0xf9, 0xbb, 0xb0, 0xbf, 0xe1, 0xa4, 0xd0, 0x5a, 0xce, 0x92, 0x74,
	0x8d, 0x0b, 0xf1, 0xb9, 0x9a, 0x58, 0xc0, 0xf7, 0xa1, 0x9b, 0x5c, 0x2e, 0xd4, 0xf5, 0xf6, 0x02,
	0x91, 0x26, 0x17, 0xf0, 0x46, 0x8f, 0xbe, 0x1f, 0xc0, 0xa0, 0xe9, 0x03, 0x63, 0x70, 0xa8, 0x92,
	0xbc, 0x1e, 0x3c, 0xfd, 0x76, 0x13, 0x5b, 0xa1, 0x42, 0x23, 0x8c, 0x1f, 0x7b, 0x2d, 0xd9, 0x5f,
	0x00, 0xb9, 0x50, 0x58, 0x1a, 0x9b, 0x94, 0x96, 0xc6, 0xd6, 0xe7, 0x2d, 0xc2, 0x7e, 0x87, 0xae,
	0xd2, 0x19, 0xbe, 0xcd, 0xa8, 0xef, 0x01, 0xf7, 0x8a, 0x0d, 0xe1, 0xa8, 0x4c, 0x6e, 0xec, 0x79,
	0x21, 0x62, 0x5d, 0x5a, 0x6a, 0x73, 0xc0, 0xdb, 0xc8, 0xd5, 0x25, 0xcc, 0x15, 0xde, 0xbf, 0xd3,
	0x42, 0x5d, 0xe9, 0x0c, 0xa9, 0xbb, 0x7d, 0xbe, 0x0f, 0x5d, 0x5d, 0x05, 0x62, 0x69, 0x3e, 0xf2,
	0x39, 0xf5, 0x73, 0xc0, 0x1b, 0xcd, 0xfe, 0x83, 0xb3, 0x12, 0x93, 0xec, 0x83, 0x92, 0x0f, 0x71,
	0xed, 0xe9, 0x93, 0xe7, 0x09, 0x77, 0xf7, 0x49, 0xb2, 0xac, 0xb1, 0x55, 0xb9, 0x6c, 0x23, 0x77,
	0x9a, 0x2f, 0x9a, 0x06, 0x7b, 0x2d, 0x72, 0xa4, 0x80, 0x07, 0xfc, 0x09, 0x77, 0x5d, 0x59, 0x6b,
	0x7b, 0x8b, 0x0f, 0xe7, 0x59, 0x56, 0x52, 0x46, 0x07, 0xbc, 0x45, 0x46, 0xaf, 0xa0, 0x5b, 0x8d,
	0xdb, 0xf5, 0x27, 0x17, 0xca, 0x4d, 0xa5, 0x1a, 0x9b, 0x57, 0x8e, 0x67, 0xcb, 0x38, 0xb1, 0x6b,
	0xdf, 0x70, 0xaf, 0x46, 0x2f, 0x21, 0xac, 0x32, 0xed, 0x0c, 0xa5, 0xd8, 0x60, 0xe9, 0x07, 0xe5,
	0xd5, 0xb3, 0x1b, 0x5f, 0x03, 0xec, 0xbe, 0x8e, 0x9f, 0xde, 0xfd, 0xad, 0x03, 0xb0, 0xcb, 0x33,
	0x1b, 0xc1, 0x71, 0x86, 0x37, 0x8f, 0x23, 0xb7, 0xc7, 0x5c, 0xbf, 0xf2, 0x64, 0x7b, 0x81, 0x36,
	0x5d, 0xd3, 0xce, 0xab, 0xbb, 0xdc, 0x27, 0xef, 0x09, 0x67, 0xff, 0xc2, 0xa9, 0x15, 0x39, 0xea,
	0x3b, 0xbb, 0xc0, 0x54, 0xab, 0xcc, 0x50, 0x92, 0x02, 0xfe, 0x88, 0xba, 0x4c, 0x2c, 0x93, 0xf6,
	0x81, 0x55, 0xa8, 0xf6, 0x61, 0xab, 0xb8, 0xf0, 0x99, 0xe2, 0xba, 0x7b, 0xc5, 0x7d, 0x39, 0x80,
	0x20, 0x9e, 0xc6, 0x2e, 0x4b, 0x06, 0x31, 0xa3, 0x40, 0xba, 0x8a, 0x42, 0xde, 0xe8, 0xe7, 0x1a,
	0xe3, 0x72, 0xb3, 0x2a, 0x8b, 0x74, 0xee, 0xdf, 0xb3, 0xa0, 0xca, 0x4d, 0x0b, 0xb9, 0x9d, 0xc2,
	0x2c, 0x10, 0xab, 0x2f, 0xa0, 0xcf, 0xbd, 0x72, 0x2f, 0xac, 0x3b, 0xdd, 0x44, 0xe1, 0x30, 0x70,
	0x2f, 0x2c, 0x09, 0xe7, 0x46, 0x95, 0x2c, 0x65, 0x1d, 0x77, 0xaf, 0x5c, 0xc7, 0x73, 0xb3, 0xda,
	0x75, 0xbc, 0x47, 0xf7, 0xdb, 0x63, 0xee, 0x2b, 0xdd, 0x60, 0x69, 0x84, 0x56, 0x14, 0xf3, 0x90,
	0xd7, 0xd2, 0x9d, 0xba, 0xc1, 0xf2, 0x52, 0x6c, 0x29, 0xd8, 0x21, 0xf7, 0xaa, 0xe6, 0xc9, 0x96,
	0x92, 0xec, 0x79, 0xb2, 0x5d, 0x76, 0xe9, 0x6f, 0xe1, 0xc5, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x49, 0x50, 0x3d, 0xac, 0x26, 0x06, 0x00, 0x00,
}
