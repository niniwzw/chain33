// Code generated by protoc-gen-go.
// source: paracross.proto
// DO NOT EDIT!

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	paracross.proto

It has these top-level messages:
	ParacrossStatusDetails
	ParacrossHeightStatus
	ParacrossStatus
	ParacrossNodeStatus
	ParacrossCommitAction
	ParacrossMinerAction
	ParacrossAction
	ReceiptParacrossCommit
	ReceiptParacrossMiner
	ReceiptParacrossDone
	ReceiptParacrossRecord
	ParacrossTx
	ReqParacrossTitleHeight
	RespParacrossTitles
	ParacrossAsset
*/
package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import types2 "gitlab.33.cn/chain33/chain33/types"
import types1 "gitlab.33.cn/chain33/chain33/types"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// stateDB
type ParacrossStatusDetails struct {
	Addrs     []string `protobuf:"bytes,1,rep,name=addrs" json:"addrs,omitempty"`
	BlockHash [][]byte `protobuf:"bytes,2,rep,name=blockHash,proto3" json:"blockHash,omitempty"`
}

func (m *ParacrossStatusDetails) Reset()                    { *m = ParacrossStatusDetails{} }
func (m *ParacrossStatusDetails) String() string            { return proto.CompactTextString(m) }
func (*ParacrossStatusDetails) ProtoMessage()               {}
func (*ParacrossStatusDetails) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ParacrossStatusDetails) GetAddrs() []string {
	if m != nil {
		return m.Addrs
	}
	return nil
}

func (m *ParacrossStatusDetails) GetBlockHash() [][]byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

type ParacrossHeightStatus struct {
	// ing, done
	Status  int32                   `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	Title   string                  `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Height  int64                   `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	Details *ParacrossStatusDetails `protobuf:"bytes,4,opt,name=details" json:"details,omitempty"`
}

func (m *ParacrossHeightStatus) Reset()                    { *m = ParacrossHeightStatus{} }
func (m *ParacrossHeightStatus) String() string            { return proto.CompactTextString(m) }
func (*ParacrossHeightStatus) ProtoMessage()               {}
func (*ParacrossHeightStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ParacrossHeightStatus) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ParacrossHeightStatus) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ParacrossHeightStatus) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ParacrossHeightStatus) GetDetails() *ParacrossStatusDetails {
	if m != nil {
		return m.Details
	}
	return nil
}

type ParacrossStatus struct {
	Title     string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Height    int64  `protobuf:"varint,2,opt,name=height" json:"height,omitempty"`
	BlockHash []byte `protobuf:"bytes,3,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
}

func (m *ParacrossStatus) Reset()                    { *m = ParacrossStatus{} }
func (m *ParacrossStatus) String() string            { return proto.CompactTextString(m) }
func (*ParacrossStatus) ProtoMessage()               {}
func (*ParacrossStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ParacrossStatus) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ParacrossStatus) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ParacrossStatus) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

// action
type ParacrossNodeStatus struct {
	MainBlockHash   []byte   `protobuf:"bytes,1,opt,name=mainBlockHash,proto3" json:"mainBlockHash,omitempty"`
	MainBlockHeight int64    `protobuf:"varint,2,opt,name=mainBlockHeight" json:"mainBlockHeight,omitempty"`
	Title           string   `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Height          int64    `protobuf:"varint,4,opt,name=height" json:"height,omitempty"`
	PreBlockHash    []byte   `protobuf:"bytes,5,opt,name=preBlockHash,proto3" json:"preBlockHash,omitempty"`
	BlockHash       []byte   `protobuf:"bytes,6,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	PreStateHash    []byte   `protobuf:"bytes,7,opt,name=preStateHash,proto3" json:"preStateHash,omitempty"`
	StateHash       []byte   `protobuf:"bytes,8,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	TxCounts        uint32   `protobuf:"varint,9,opt,name=txCounts" json:"txCounts,omitempty"`
	TxResult        []byte   `protobuf:"bytes,10,opt,name=txResult,proto3" json:"txResult,omitempty"`
	TxHashs         [][]byte `protobuf:"bytes,11,rep,name=txHashs,proto3" json:"txHashs,omitempty"`
	CrossTxResult   []byte   `protobuf:"bytes,12,opt,name=crossTxResult,proto3" json:"crossTxResult,omitempty"`
	CrossTxHashs    [][]byte `protobuf:"bytes,13,rep,name=crossTxHashs,proto3" json:"crossTxHashs,omitempty"`
}

func (m *ParacrossNodeStatus) Reset()                    { *m = ParacrossNodeStatus{} }
func (m *ParacrossNodeStatus) String() string            { return proto.CompactTextString(m) }
func (*ParacrossNodeStatus) ProtoMessage()               {}
func (*ParacrossNodeStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ParacrossNodeStatus) GetMainBlockHash() []byte {
	if m != nil {
		return m.MainBlockHash
	}
	return nil
}

func (m *ParacrossNodeStatus) GetMainBlockHeight() int64 {
	if m != nil {
		return m.MainBlockHeight
	}
	return 0
}

func (m *ParacrossNodeStatus) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ParacrossNodeStatus) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ParacrossNodeStatus) GetPreBlockHash() []byte {
	if m != nil {
		return m.PreBlockHash
	}
	return nil
}

func (m *ParacrossNodeStatus) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *ParacrossNodeStatus) GetPreStateHash() []byte {
	if m != nil {
		return m.PreStateHash
	}
	return nil
}

func (m *ParacrossNodeStatus) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *ParacrossNodeStatus) GetTxCounts() uint32 {
	if m != nil {
		return m.TxCounts
	}
	return 0
}

func (m *ParacrossNodeStatus) GetTxResult() []byte {
	if m != nil {
		return m.TxResult
	}
	return nil
}

func (m *ParacrossNodeStatus) GetTxHashs() [][]byte {
	if m != nil {
		return m.TxHashs
	}
	return nil
}

func (m *ParacrossNodeStatus) GetCrossTxResult() []byte {
	if m != nil {
		return m.CrossTxResult
	}
	return nil
}

func (m *ParacrossNodeStatus) GetCrossTxHashs() [][]byte {
	if m != nil {
		return m.CrossTxHashs
	}
	return nil
}

type ParacrossCommitAction struct {
	Status *ParacrossNodeStatus `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *ParacrossCommitAction) Reset()                    { *m = ParacrossCommitAction{} }
func (m *ParacrossCommitAction) String() string            { return proto.CompactTextString(m) }
func (*ParacrossCommitAction) ProtoMessage()               {}
func (*ParacrossCommitAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ParacrossCommitAction) GetStatus() *ParacrossNodeStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ParacrossMinerAction struct {
	Status *ParacrossNodeStatus `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *ParacrossMinerAction) Reset()                    { *m = ParacrossMinerAction{} }
func (m *ParacrossMinerAction) String() string            { return proto.CompactTextString(m) }
func (*ParacrossMinerAction) ProtoMessage()               {}
func (*ParacrossMinerAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ParacrossMinerAction) GetStatus() *ParacrossNodeStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ParacrossAction struct {
	// Types that are valid to be assigned to Value:
	//	*ParacrossAction_Commit
	//	*ParacrossAction_Miner
	//	*ParacrossAction_AssetTransfer
	//	*ParacrossAction_AssetWithdraw
	//	*ParacrossAction_Transfer
	//	*ParacrossAction_Withdraw
	//	*ParacrossAction_TransferToExec
	Value isParacrossAction_Value `protobuf_oneof:"value"`
	Ty    int32                   `protobuf:"varint,2,opt,name=ty" json:"ty,omitempty"`
}

func (m *ParacrossAction) Reset()                    { *m = ParacrossAction{} }
func (m *ParacrossAction) String() string            { return proto.CompactTextString(m) }
func (*ParacrossAction) ProtoMessage()               {}
func (*ParacrossAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type isParacrossAction_Value interface {
	isParacrossAction_Value()
}

type ParacrossAction_Commit struct {
	Commit *ParacrossCommitAction `protobuf:"bytes,1,opt,name=commit,oneof"`
}
type ParacrossAction_Miner struct {
	Miner *ParacrossMinerAction `protobuf:"bytes,3,opt,name=miner,oneof"`
}
type ParacrossAction_AssetTransfer struct {
	AssetTransfer *types2.AssetsTransfer `protobuf:"bytes,4,opt,name=assetTransfer,oneof"`
}
type ParacrossAction_AssetWithdraw struct {
	AssetWithdraw *types2.AssetsWithdraw `protobuf:"bytes,5,opt,name=assetWithdraw,oneof"`
}
type ParacrossAction_Transfer struct {
	Transfer *types2.AssetsTransfer `protobuf:"bytes,6,opt,name=transfer,oneof"`
}
type ParacrossAction_Withdraw struct {
	Withdraw *types2.AssetsWithdraw `protobuf:"bytes,7,opt,name=withdraw,oneof"`
}
type ParacrossAction_TransferToExec struct {
	TransferToExec *types2.AssetsTransferToExec `protobuf:"bytes,8,opt,name=transferToExec,oneof"`
}

func (*ParacrossAction_Commit) isParacrossAction_Value()         {}
func (*ParacrossAction_Miner) isParacrossAction_Value()          {}
func (*ParacrossAction_AssetTransfer) isParacrossAction_Value()  {}
func (*ParacrossAction_AssetWithdraw) isParacrossAction_Value()  {}
func (*ParacrossAction_Transfer) isParacrossAction_Value()       {}
func (*ParacrossAction_Withdraw) isParacrossAction_Value()       {}
func (*ParacrossAction_TransferToExec) isParacrossAction_Value() {}

func (m *ParacrossAction) GetValue() isParacrossAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ParacrossAction) GetCommit() *ParacrossCommitAction {
	if x, ok := m.GetValue().(*ParacrossAction_Commit); ok {
		return x.Commit
	}
	return nil
}

func (m *ParacrossAction) GetMiner() *ParacrossMinerAction {
	if x, ok := m.GetValue().(*ParacrossAction_Miner); ok {
		return x.Miner
	}
	return nil
}

func (m *ParacrossAction) GetAssetTransfer() *types2.AssetsTransfer {
	if x, ok := m.GetValue().(*ParacrossAction_AssetTransfer); ok {
		return x.AssetTransfer
	}
	return nil
}

func (m *ParacrossAction) GetAssetWithdraw() *types2.AssetsWithdraw {
	if x, ok := m.GetValue().(*ParacrossAction_AssetWithdraw); ok {
		return x.AssetWithdraw
	}
	return nil
}

func (m *ParacrossAction) GetTransfer() *types2.AssetsTransfer {
	if x, ok := m.GetValue().(*ParacrossAction_Transfer); ok {
		return x.Transfer
	}
	return nil
}

func (m *ParacrossAction) GetWithdraw() *types2.AssetsWithdraw {
	if x, ok := m.GetValue().(*ParacrossAction_Withdraw); ok {
		return x.Withdraw
	}
	return nil
}

func (m *ParacrossAction) GetTransferToExec() *types2.AssetsTransferToExec {
	if x, ok := m.GetValue().(*ParacrossAction_TransferToExec); ok {
		return x.TransferToExec
	}
	return nil
}

func (m *ParacrossAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ParacrossAction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ParacrossAction_OneofMarshaler, _ParacrossAction_OneofUnmarshaler, _ParacrossAction_OneofSizer, []interface{}{
		(*ParacrossAction_Commit)(nil),
		(*ParacrossAction_Miner)(nil),
		(*ParacrossAction_AssetTransfer)(nil),
		(*ParacrossAction_AssetWithdraw)(nil),
		(*ParacrossAction_Transfer)(nil),
		(*ParacrossAction_Withdraw)(nil),
		(*ParacrossAction_TransferToExec)(nil),
	}
}

func _ParacrossAction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ParacrossAction)
	// value
	switch x := m.Value.(type) {
	case *ParacrossAction_Commit:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Commit); err != nil {
			return err
		}
	case *ParacrossAction_Miner:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Miner); err != nil {
			return err
		}
	case *ParacrossAction_AssetTransfer:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AssetTransfer); err != nil {
			return err
		}
	case *ParacrossAction_AssetWithdraw:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AssetWithdraw); err != nil {
			return err
		}
	case *ParacrossAction_Transfer:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Transfer); err != nil {
			return err
		}
	case *ParacrossAction_Withdraw:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Withdraw); err != nil {
			return err
		}
	case *ParacrossAction_TransferToExec:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TransferToExec); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ParacrossAction.Value has unexpected type %T", x)
	}
	return nil
}

func _ParacrossAction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ParacrossAction)
	switch tag {
	case 1: // value.commit
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ParacrossCommitAction)
		err := b.DecodeMessage(msg)
		m.Value = &ParacrossAction_Commit{msg}
		return true, err
	case 3: // value.miner
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ParacrossMinerAction)
		err := b.DecodeMessage(msg)
		m.Value = &ParacrossAction_Miner{msg}
		return true, err
	case 4: // value.assetTransfer
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(types2.AssetsTransfer)
		err := b.DecodeMessage(msg)
		m.Value = &ParacrossAction_AssetTransfer{msg}
		return true, err
	case 5: // value.assetWithdraw
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(types2.AssetsWithdraw)
		err := b.DecodeMessage(msg)
		m.Value = &ParacrossAction_AssetWithdraw{msg}
		return true, err
	case 6: // value.transfer
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(types2.AssetsTransfer)
		err := b.DecodeMessage(msg)
		m.Value = &ParacrossAction_Transfer{msg}
		return true, err
	case 7: // value.withdraw
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(types2.AssetsWithdraw)
		err := b.DecodeMessage(msg)
		m.Value = &ParacrossAction_Withdraw{msg}
		return true, err
	case 8: // value.transferToExec
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(types2.AssetsTransferToExec)
		err := b.DecodeMessage(msg)
		m.Value = &ParacrossAction_TransferToExec{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ParacrossAction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ParacrossAction)
	// value
	switch x := m.Value.(type) {
	case *ParacrossAction_Commit:
		s := proto.Size(x.Commit)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ParacrossAction_Miner:
		s := proto.Size(x.Miner)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ParacrossAction_AssetTransfer:
		s := proto.Size(x.AssetTransfer)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ParacrossAction_AssetWithdraw:
		s := proto.Size(x.AssetWithdraw)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ParacrossAction_Transfer:
		s := proto.Size(x.Transfer)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ParacrossAction_Withdraw:
		s := proto.Size(x.Withdraw)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ParacrossAction_TransferToExec:
		s := proto.Size(x.TransferToExec)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// logs
type ReceiptParacrossCommit struct {
	Addr    string                 `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Status  *ParacrossNodeStatus   `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	Prev    *ParacrossHeightStatus `protobuf:"bytes,3,opt,name=prev" json:"prev,omitempty"`
	Current *ParacrossHeightStatus `protobuf:"bytes,4,opt,name=current" json:"current,omitempty"`
}

func (m *ReceiptParacrossCommit) Reset()                    { *m = ReceiptParacrossCommit{} }
func (m *ReceiptParacrossCommit) String() string            { return proto.CompactTextString(m) }
func (*ReceiptParacrossCommit) ProtoMessage()               {}
func (*ReceiptParacrossCommit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ReceiptParacrossCommit) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *ReceiptParacrossCommit) GetStatus() *ParacrossNodeStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReceiptParacrossCommit) GetPrev() *ParacrossHeightStatus {
	if m != nil {
		return m.Prev
	}
	return nil
}

func (m *ReceiptParacrossCommit) GetCurrent() *ParacrossHeightStatus {
	if m != nil {
		return m.Current
	}
	return nil
}

type ReceiptParacrossMiner struct {
	Status *ParacrossNodeStatus `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *ReceiptParacrossMiner) Reset()                    { *m = ReceiptParacrossMiner{} }
func (m *ReceiptParacrossMiner) String() string            { return proto.CompactTextString(m) }
func (*ReceiptParacrossMiner) ProtoMessage()               {}
func (*ReceiptParacrossMiner) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ReceiptParacrossMiner) GetStatus() *ParacrossNodeStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReceiptParacrossDone struct {
	TotalNodes     int32  `protobuf:"varint,1,opt,name=totalNodes" json:"totalNodes,omitempty"`
	TotalCommit    int32  `protobuf:"varint,2,opt,name=totalCommit" json:"totalCommit,omitempty"`
	MostSameCommit int32  `protobuf:"varint,3,opt,name=mostSameCommit" json:"mostSameCommit,omitempty"`
	Title          string `protobuf:"bytes,4,opt,name=title" json:"title,omitempty"`
	Height         int64  `protobuf:"varint,5,opt,name=height" json:"height,omitempty"`
	StateHash      []byte `protobuf:"bytes,6,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	TxCounts       uint32 `protobuf:"varint,7,opt,name=txCounts" json:"txCounts,omitempty"`
	TxResult       []byte `protobuf:"bytes,8,opt,name=txResult,proto3" json:"txResult,omitempty"`
}

func (m *ReceiptParacrossDone) Reset()                    { *m = ReceiptParacrossDone{} }
func (m *ReceiptParacrossDone) String() string            { return proto.CompactTextString(m) }
func (*ReceiptParacrossDone) ProtoMessage()               {}
func (*ReceiptParacrossDone) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ReceiptParacrossDone) GetTotalNodes() int32 {
	if m != nil {
		return m.TotalNodes
	}
	return 0
}

func (m *ReceiptParacrossDone) GetTotalCommit() int32 {
	if m != nil {
		return m.TotalCommit
	}
	return 0
}

func (m *ReceiptParacrossDone) GetMostSameCommit() int32 {
	if m != nil {
		return m.MostSameCommit
	}
	return 0
}

func (m *ReceiptParacrossDone) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ReceiptParacrossDone) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ReceiptParacrossDone) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *ReceiptParacrossDone) GetTxCounts() uint32 {
	if m != nil {
		return m.TxCounts
	}
	return 0
}

func (m *ReceiptParacrossDone) GetTxResult() []byte {
	if m != nil {
		return m.TxResult
	}
	return nil
}

type ReceiptParacrossRecord struct {
	Addr   string               `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Status *ParacrossNodeStatus `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
}

func (m *ReceiptParacrossRecord) Reset()                    { *m = ReceiptParacrossRecord{} }
func (m *ReceiptParacrossRecord) String() string            { return proto.CompactTextString(m) }
func (*ReceiptParacrossRecord) ProtoMessage()               {}
func (*ReceiptParacrossRecord) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ReceiptParacrossRecord) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *ReceiptParacrossRecord) GetStatus() *ParacrossNodeStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// LocalDB
// title-height-addr : txHash
type ParacrossTx struct {
	TxHash string `protobuf:"bytes,1,opt,name=txHash" json:"txHash,omitempty"`
}

func (m *ParacrossTx) Reset()                    { *m = ParacrossTx{} }
func (m *ParacrossTx) String() string            { return proto.CompactTextString(m) }
func (*ParacrossTx) ProtoMessage()               {}
func (*ParacrossTx) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ParacrossTx) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

// query
type ReqParacrossTitleHeight struct {
	Title  string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Height int64  `protobuf:"varint,2,opt,name=height" json:"height,omitempty"`
}

func (m *ReqParacrossTitleHeight) Reset()                    { *m = ReqParacrossTitleHeight{} }
func (m *ReqParacrossTitleHeight) String() string            { return proto.CompactTextString(m) }
func (*ReqParacrossTitleHeight) ProtoMessage()               {}
func (*ReqParacrossTitleHeight) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *ReqParacrossTitleHeight) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ReqParacrossTitleHeight) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type RespParacrossTitles struct {
	Titles []*ReceiptParacrossDone `protobuf:"bytes,1,rep,name=titles" json:"titles,omitempty"`
}

func (m *RespParacrossTitles) Reset()                    { *m = RespParacrossTitles{} }
func (m *RespParacrossTitles) String() string            { return proto.CompactTextString(m) }
func (*RespParacrossTitles) ProtoMessage()               {}
func (*RespParacrossTitles) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *RespParacrossTitles) GetTitles() []*ReceiptParacrossDone {
	if m != nil {
		return m.Titles
	}
	return nil
}

// 跨链转账相关
type ParacrossAsset struct {
	// input
	From       string `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	To         string `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
	IsWithdraw bool   `protobuf:"varint,3,opt,name=isWithdraw" json:"isWithdraw,omitempty"`
	TxHash     []byte `protobuf:"bytes,4,opt,name=txHash,proto3" json:"txHash,omitempty"`
	Amount     int64  `protobuf:"varint,5,opt,name=amount" json:"amount,omitempty"`
	Exec       string `protobuf:"bytes,6,opt,name=exec" json:"exec,omitempty"`
	Symbol     string `protobuf:"bytes,7,opt,name=symbol" json:"symbol,omitempty"`
	// 主链部分
	Height int64 `protobuf:"varint,10,opt,name=height" json:"height,omitempty"`
	// 平行链部分
	CommitDoneHeight int64 `protobuf:"varint,21,opt,name=commitDoneHeight" json:"commitDoneHeight,omitempty"`
	ParaHeight       int64 `protobuf:"varint,22,opt,name=paraHeight" json:"paraHeight,omitempty"`
	Success          bool  `protobuf:"varint,23,opt,name=success" json:"success,omitempty"`
}

func (m *ParacrossAsset) Reset()                    { *m = ParacrossAsset{} }
func (m *ParacrossAsset) String() string            { return proto.CompactTextString(m) }
func (*ParacrossAsset) ProtoMessage()               {}
func (*ParacrossAsset) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *ParacrossAsset) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *ParacrossAsset) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *ParacrossAsset) GetIsWithdraw() bool {
	if m != nil {
		return m.IsWithdraw
	}
	return false
}

func (m *ParacrossAsset) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *ParacrossAsset) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ParacrossAsset) GetExec() string {
	if m != nil {
		return m.Exec
	}
	return ""
}

func (m *ParacrossAsset) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ParacrossAsset) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ParacrossAsset) GetCommitDoneHeight() int64 {
	if m != nil {
		return m.CommitDoneHeight
	}
	return 0
}

func (m *ParacrossAsset) GetParaHeight() int64 {
	if m != nil {
		return m.ParaHeight
	}
	return 0
}

func (m *ParacrossAsset) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*ParacrossStatusDetails)(nil), "types.ParacrossStatusDetails")
	proto.RegisterType((*ParacrossHeightStatus)(nil), "types.ParacrossHeightStatus")
	proto.RegisterType((*ParacrossStatus)(nil), "types.ParacrossStatus")
	proto.RegisterType((*ParacrossNodeStatus)(nil), "types.ParacrossNodeStatus")
	proto.RegisterType((*ParacrossCommitAction)(nil), "types.ParacrossCommitAction")
	proto.RegisterType((*ParacrossMinerAction)(nil), "types.ParacrossMinerAction")
	proto.RegisterType((*ParacrossAction)(nil), "types.ParacrossAction")
	proto.RegisterType((*ReceiptParacrossCommit)(nil), "types.ReceiptParacrossCommit")
	proto.RegisterType((*ReceiptParacrossMiner)(nil), "types.ReceiptParacrossMiner")
	proto.RegisterType((*ReceiptParacrossDone)(nil), "types.ReceiptParacrossDone")
	proto.RegisterType((*ReceiptParacrossRecord)(nil), "types.ReceiptParacrossRecord")
	proto.RegisterType((*ParacrossTx)(nil), "types.ParacrossTx")
	proto.RegisterType((*ReqParacrossTitleHeight)(nil), "types.ReqParacrossTitleHeight")
	proto.RegisterType((*RespParacrossTitles)(nil), "types.RespParacrossTitles")
	proto.RegisterType((*ParacrossAsset)(nil), "types.ParacrossAsset")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Paracross service

type ParacrossClient interface {
	GetTitle(ctx context.Context, in *types1.ReqString, opts ...grpc.CallOption) (*ParacrossStatus, error)
	ListTitles(ctx context.Context, in *types1.ReqNil, opts ...grpc.CallOption) (*RespParacrossTitles, error)
	GetTitleHeight(ctx context.Context, in *ReqParacrossTitleHeight, opts ...grpc.CallOption) (*ReceiptParacrossDone, error)
	GetAssetTxResult(ctx context.Context, in *types1.ReqHash, opts ...grpc.CallOption) (*ParacrossAsset, error)
}

type paracrossClient struct {
	cc *grpc.ClientConn
}

func NewParacrossClient(cc *grpc.ClientConn) ParacrossClient {
	return &paracrossClient{cc}
}

func (c *paracrossClient) GetTitle(ctx context.Context, in *types1.ReqString, opts ...grpc.CallOption) (*ParacrossStatus, error) {
	out := new(ParacrossStatus)
	err := grpc.Invoke(ctx, "/types.paracross/GetTitle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paracrossClient) ListTitles(ctx context.Context, in *types1.ReqNil, opts ...grpc.CallOption) (*RespParacrossTitles, error) {
	out := new(RespParacrossTitles)
	err := grpc.Invoke(ctx, "/types.paracross/ListTitles", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paracrossClient) GetTitleHeight(ctx context.Context, in *ReqParacrossTitleHeight, opts ...grpc.CallOption) (*ReceiptParacrossDone, error) {
	out := new(ReceiptParacrossDone)
	err := grpc.Invoke(ctx, "/types.paracross/GetTitleHeight", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paracrossClient) GetAssetTxResult(ctx context.Context, in *types1.ReqHash, opts ...grpc.CallOption) (*ParacrossAsset, error) {
	out := new(ParacrossAsset)
	err := grpc.Invoke(ctx, "/types.paracross/GetAssetTxResult", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Paracross service

type ParacrossServer interface {
	GetTitle(context.Context, *types1.ReqString) (*ParacrossStatus, error)
	ListTitles(context.Context, *types1.ReqNil) (*RespParacrossTitles, error)
	GetTitleHeight(context.Context, *ReqParacrossTitleHeight) (*ReceiptParacrossDone, error)
	GetAssetTxResult(context.Context, *types1.ReqHash) (*ParacrossAsset, error)
}

func RegisterParacrossServer(s *grpc.Server, srv ParacrossServer) {
	s.RegisterService(&_Paracross_serviceDesc, srv)
}

func _Paracross_GetTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types1.ReqString)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParacrossServer).GetTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.paracross/GetTitle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParacrossServer).GetTitle(ctx, req.(*types1.ReqString))
	}
	return interceptor(ctx, in, info, handler)
}

func _Paracross_ListTitles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types1.ReqNil)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParacrossServer).ListTitles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.paracross/ListTitles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParacrossServer).ListTitles(ctx, req.(*types1.ReqNil))
	}
	return interceptor(ctx, in, info, handler)
}

func _Paracross_GetTitleHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqParacrossTitleHeight)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParacrossServer).GetTitleHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.paracross/GetTitleHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParacrossServer).GetTitleHeight(ctx, req.(*ReqParacrossTitleHeight))
	}
	return interceptor(ctx, in, info, handler)
}

func _Paracross_GetAssetTxResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types1.ReqHash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParacrossServer).GetAssetTxResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.paracross/GetAssetTxResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParacrossServer).GetAssetTxResult(ctx, req.(*types1.ReqHash))
	}
	return interceptor(ctx, in, info, handler)
}

var _Paracross_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.paracross",
	HandlerType: (*ParacrossServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTitle",
			Handler:    _Paracross_GetTitle_Handler,
		},
		{
			MethodName: "ListTitles",
			Handler:    _Paracross_ListTitles_Handler,
		},
		{
			MethodName: "GetTitleHeight",
			Handler:    _Paracross_GetTitleHeight_Handler,
		},
		{
			MethodName: "GetAssetTxResult",
			Handler:    _Paracross_GetAssetTxResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "paracross.proto",
}

func init() { proto.RegisterFile("paracross.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 985 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x56, 0x4b, 0x6f, 0x23, 0x45,
	0x10, 0xf6, 0x8c, 0xdf, 0xe5, 0xd8, 0x09, 0xbd, 0x89, 0x77, 0x34, 0x2c, 0xc8, 0x1a, 0x01, 0xb2,
	0x38, 0x44, 0xc8, 0x91, 0x82, 0x10, 0xe2, 0xb0, 0x2f, 0xc5, 0x5a, 0x96, 0x05, 0x75, 0x22, 0x71,
	0x42, 0x62, 0x32, 0xee, 0xdd, 0x8c, 0xf0, 0xb8, 0xbd, 0xd3, 0xed, 0xdd, 0xe4, 0xce, 0x91, 0x33,
	0x3f, 0x83, 0x0b, 0xff, 0x81, 0x03, 0xbf, 0x0a, 0x75, 0xf5, 0x63, 0x1e, 0xb6, 0xac, 0x08, 0xb8,
	0x75, 0x55, 0x7f, 0x55, 0xd5, 0x55, 0xf5, 0x75, 0x75, 0xc3, 0xe1, 0x3a, 0xce, 0xe3, 0x24, 0xe7,
	0x42, 0x9c, 0xae, 0x73, 0x2e, 0x39, 0x69, 0xcb, 0xbb, 0x35, 0x13, 0xe1, 0x07, 0x32, 0x8f, 0x57,
	0x22, 0x4e, 0x64, 0xca, 0x57, 0x7a, 0x27, 0x3c, 0x48, 0x78, 0x96, 0x59, 0x29, 0x7a, 0x09, 0xe3,
	0x1f, 0xac, 0xe9, 0xa5, 0x8c, 0xe5, 0x46, 0x3c, 0x63, 0x32, 0x4e, 0x97, 0x82, 0x1c, 0x43, 0x3b,
	0x5e, 0x2c, 0x72, 0x11, 0x78, 0x93, 0xe6, 0xb4, 0x4f, 0xb5, 0x40, 0x1e, 0x41, 0xff, 0x7a, 0xc9,
	0x93, 0x5f, 0xe6, 0xb1, 0xb8, 0x09, 0xfc, 0x49, 0x73, 0x7a, 0x40, 0x0b, 0x45, 0xf4, 0xbb, 0x07,
	0x27, 0xce, 0xdd, 0x9c, 0xa5, 0x6f, 0x6e, 0xa4, 0x76, 0x4a, 0xc6, 0xd0, 0x11, 0xb8, 0x0a, 0xbc,
	0x89, 0x37, 0x6d, 0x53, 0x23, 0xa9, 0x28, 0x32, 0x95, 0x4b, 0x16, 0xf8, 0x13, 0x4f, 0x45, 0x41,
	0x41, 0xa1, 0x6f, 0xd0, 0x3a, 0x68, 0x4e, 0xbc, 0x69, 0x93, 0x1a, 0x89, 0x7c, 0x09, 0xdd, 0x85,
	0x3e, 0x5e, 0xd0, 0x9a, 0x78, 0xd3, 0xc1, 0xec, 0xa3, 0x53, 0xcc, 0xf3, 0x74, 0x77, 0x0e, 0xd4,
	0xa2, 0xa3, 0x9f, 0xe0, 0xb0, 0x06, 0x29, 0x22, 0x7b, 0xbb, 0x23, 0xfb, 0x95, 0xc8, 0x95, 0xbc,
	0xd5, 0xa1, 0x2a, 0x79, 0xff, 0xd9, 0x84, 0x07, 0xce, 0xff, 0x2b, 0xbe, 0x60, 0x26, 0xc6, 0x27,
	0x30, 0xcc, 0xe2, 0x74, 0xf5, 0xc4, 0x59, 0x7a, 0x68, 0x59, 0x55, 0x92, 0x29, 0x1c, 0x16, 0x8a,
	0x72, 0xf0, 0xba, 0xba, 0x38, 0x73, 0x73, 0xf7, 0x99, 0x5b, 0x95, 0x33, 0x47, 0x70, 0xb0, 0xce,
	0x59, 0x11, 0xbc, 0x8d, 0xc1, 0x2b, 0xba, 0x6a, 0x5e, 0x9d, 0x5a, 0x5e, 0xc6, 0x83, 0x4a, 0x86,
	0x21, 0xa0, 0xeb, 0x3c, 0x38, 0x9d, 0xf2, 0x20, 0x1c, 0xa0, 0xa7, 0x3d, 0x38, 0x05, 0x09, 0xa1,
	0x27, 0x6f, 0x9f, 0xf2, 0xcd, 0x4a, 0x8a, 0xa0, 0x3f, 0xf1, 0xa6, 0x43, 0xea, 0x64, 0xbd, 0x47,
	0x99, 0xd8, 0x2c, 0x65, 0x00, 0x68, 0xe8, 0x64, 0x12, 0x40, 0x57, 0xde, 0x2a, 0x0f, 0x22, 0x18,
	0x20, 0xcb, 0xac, 0xa8, 0x6a, 0x8a, 0x65, 0xbe, 0xb2, 0xa6, 0x07, 0xba, 0xa6, 0x15, 0xa5, 0x3a,
	0xb9, 0x51, 0x68, 0x27, 0x43, 0x74, 0x52, 0xd1, 0x45, 0xdf, 0x96, 0xc8, 0xfa, 0x94, 0x67, 0x59,
	0x2a, 0x1f, 0xe3, 0x45, 0x21, 0xb3, 0x0a, 0x59, 0x07, 0xb3, 0xb0, 0xce, 0xb2, 0xa2, 0xc5, 0x96,
	0xc8, 0xd1, 0x0b, 0x38, 0x76, 0xdb, 0xdf, 0xa5, 0x2b, 0x96, 0xff, 0x07, 0x5f, 0x7f, 0x35, 0x4b,
	0x74, 0x35, 0x7e, 0xce, 0xa1, 0x93, 0xe0, 0x19, 0x8d, 0x9f, 0x47, 0x75, 0x3f, 0xe5, 0x0c, 0xe6,
	0x0d, 0x6a, 0xd0, 0xe4, 0x0c, 0xda, 0x99, 0x3a, 0x0e, 0x52, 0x66, 0x30, 0xfb, 0xb0, 0x6e, 0x56,
	0x3a, 0xeb, 0xbc, 0x41, 0x35, 0x96, 0x7c, 0x03, 0xc3, 0x58, 0x08, 0x26, 0xaf, 0xd4, 0xf4, 0x78,
	0xcd, 0x72, 0x73, 0xdb, 0x4e, 0x8c, 0xf1, 0x63, 0xb5, 0x27, 0xec, 0xe6, 0xbc, 0x41, 0xab, 0x68,
	0x67, 0xfe, 0x63, 0x2a, 0x6f, 0x16, 0x79, 0xfc, 0x1e, 0x99, 0x57, 0x37, 0xb7, 0x9b, 0xce, 0xdc,
	0x2a, 0xc8, 0x19, 0xf4, 0xa4, 0x0d, 0xdc, 0xd9, 0x1f, 0xd8, 0x01, 0x95, 0xd1, 0x7b, 0x1b, 0xae,
	0xbb, 0x3f, 0x9c, 0x03, 0x92, 0xe7, 0x30, 0xb2, 0x0e, 0xae, 0xf8, 0xf3, 0x5b, 0x96, 0x20, 0x81,
	0x8b, 0x2a, 0x55, 0xe3, 0x69, 0xc8, 0xbc, 0x41, 0x6b, 0x46, 0x64, 0x04, 0xbe, 0xbc, 0xc3, 0x3b,
	0xdb, 0xa6, 0xbe, 0xbc, 0x7b, 0xd2, 0x85, 0xf6, 0xbb, 0x78, 0xb9, 0x61, 0xd1, 0xdf, 0x1e, 0x8c,
	0x29, 0x4b, 0x58, 0xba, 0x96, 0xb5, 0x3e, 0x11, 0x02, 0x2d, 0x35, 0x51, 0xcd, 0xf4, 0xc1, 0x75,
	0x89, 0x2b, 0xfe, 0x7d, 0xb9, 0x42, 0xbe, 0x80, 0xd6, 0x3a, 0x67, 0xef, 0x4c, 0x7b, 0xb7, 0x58,
	0x51, 0x1e, 0xc2, 0x14, 0x91, 0xe4, 0x1c, 0xba, 0xc9, 0x26, 0xcf, 0xd9, 0x4a, 0x9a, 0xb6, 0xee,
	0x37, 0xb2, 0x60, 0x75, 0x5d, 0xea, 0xb9, 0x20, 0x79, 0xfe, 0x15, 0xc5, 0x7f, 0xf5, 0xe1, 0xb8,
	0xee, 0xed, 0x19, 0x5f, 0x31, 0xf2, 0x31, 0x80, 0xe4, 0x32, 0x5e, 0x2a, 0x1b, 0xfb, 0x58, 0x94,
	0x34, 0x64, 0x02, 0x03, 0x94, 0x74, 0x19, 0x4d, 0xd1, 0xcb, 0x2a, 0xf2, 0x19, 0x8c, 0x32, 0x2e,
	0xe4, 0x65, 0x9c, 0x31, 0x03, 0x6a, 0x22, 0xa8, 0xa6, 0x2d, 0x86, 0x69, 0x6b, 0xf7, 0x30, 0x6d,
	0xd7, 0x1f, 0x80, 0x62, 0xcc, 0x75, 0xf6, 0x8d, 0xb9, 0xee, 0x9e, 0x31, 0xd7, 0xab, 0x8e, 0xb9,
	0xe8, 0xe7, 0x6d, 0x7e, 0x50, 0x96, 0xf0, 0x7c, 0xf1, 0x7f, 0xf1, 0x23, 0xfa, 0x14, 0x06, 0x6e,
	0xfb, 0xea, 0x56, 0xa5, 0xa7, 0x07, 0xa9, 0x71, 0x6c, 0xa4, 0xe8, 0x02, 0x1e, 0x52, 0xf6, 0xb6,
	0x40, 0xaa, 0x5a, 0xd4, 0x1f, 0x9d, 0xfb, 0x3c, 0x94, 0xd1, 0x0b, 0x78, 0x40, 0x99, 0x58, 0x57,
	0x3d, 0x09, 0x72, 0x06, 0x1d, 0xb4, 0xd3, 0xdf, 0x89, 0xe2, 0x86, 0xed, 0xe2, 0x00, 0x35, 0xd0,
	0xe8, 0x0f, 0x1f, 0x46, 0xc5, 0x1c, 0x54, 0x77, 0x51, 0x95, 0xe5, 0x75, 0xce, 0x33, 0x5b, 0x16,
	0xb5, 0xc6, 0xeb, 0xc7, 0xcd, 0x07, 0xc2, 0x97, 0x5c, 0x51, 0x28, 0x75, 0xf7, 0x1d, 0x9b, 0xdf,
	0xa3, 0x25, 0x4d, 0xa9, 0x06, 0x2d, 0x6c, 0x87, 0x91, 0x94, 0x3e, 0xce, 0x54, 0xcf, 0x6c, 0xeb,
	0xb5, 0xa4, 0x62, 0x32, 0x35, 0x1b, 0x3a, 0x3a, 0xa6, 0x5a, 0xe3, 0x7f, 0xe6, 0x2e, 0xbb, 0xe6,
	0x4b, 0x6c, 0x77, 0x9f, 0x1a, 0xa9, 0x54, 0x16, 0xa8, 0xd0, 0xe7, 0x73, 0x38, 0xd2, 0x03, 0x59,
	0x25, 0x68, 0x1e, 0xf9, 0x13, 0x44, 0x6c, 0xe9, 0xd5, 0xf9, 0xd5, 0x77, 0xce, 0xa0, 0xc6, 0x88,
	0x2a, 0x69, 0xd4, 0xdb, 0x28, 0x36, 0x49, 0xc2, 0x84, 0x08, 0x1e, 0x62, 0x72, 0x56, 0x9c, 0xfd,
	0xe6, 0x43, 0xdf, 0xfd, 0x04, 0xc9, 0x39, 0xf4, 0x2e, 0x98, 0xc4, 0x06, 0x90, 0x23, 0x57, 0xef,
	0xb7, 0x97, 0x32, 0x4f, 0x57, 0x6f, 0xc2, 0xf1, 0xee, 0xaf, 0x53, 0xd4, 0x20, 0x5f, 0x01, 0xbc,
	0x4c, 0x85, 0x34, 0x9d, 0x1b, 0x16, 0x96, 0xaf, 0xd2, 0x65, 0x18, 0x3a, 0x71, 0xab, 0xc9, 0x51,
	0x83, 0x7c, 0x0f, 0x23, 0x1b, 0xd2, 0x26, 0x53, 0x98, 0xef, 0x62, 0x57, 0xb8, 0x8f, 0x08, 0x51,
	0x83, 0x7c, 0x0d, 0x47, 0x17, 0x4c, 0x62, 0xef, 0xdd, 0xdb, 0x3e, 0x2a, 0x5c, 0xaa, 0xbe, 0x85,
	0x27, 0xf5, 0x4c, 0x10, 0x1e, 0x35, 0xae, 0x3b, 0xf8, 0xc7, 0x3d, 0xfb, 0x27, 0x00, 0x00, 0xff,
	0xff, 0x6c, 0x00, 0x31, 0x06, 0x1e, 0x0b, 0x00, 0x00,
}
