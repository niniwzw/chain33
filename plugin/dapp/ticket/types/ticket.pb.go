// Code generated by protoc-gen-go.
// source: ticket.proto
// DO NOT EDIT!

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	ticket.proto

It has these top-level messages:
	Ticket
	TicketAction
	TicketMiner
	MinerFlag
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
	ReqBindMiner
	ReplyBindMiner
*/
package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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

type Ticket struct {
	TicketId string `protobuf:"bytes,1,opt,name=ticketId" json:"ticketId,omitempty"`
	// 0 -> 未成熟 1 -> 可挖矿 2 -> 已挖成功 3-> 已关闭
	Status int32 `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
	// genesis 创建的私钥比较特殊
	IsGenesis bool `protobuf:"varint,3,opt,name=isGenesis" json:"isGenesis,omitempty"`
	// 创建时间
	CreateTime int64 `protobuf:"varint,4,opt,name=createTime" json:"createTime,omitempty"`
	// 挖矿时间
	MinerTime int64 `protobuf:"varint,5,opt,name=minerTime" json:"minerTime,omitempty"`
	// 挖到的币的数目
	MinerValue   int64  `protobuf:"varint,8,opt,name=minerValue" json:"minerValue,omitempty"`
	MinerAddress string `protobuf:"bytes,6,opt,name=minerAddress" json:"minerAddress,omitempty"`
	// return wallet
	ReturnAddress string `protobuf:"bytes,7,opt,name=returnAddress" json:"returnAddress,omitempty"`
}

func (m *Ticket) Reset()                    { *m = Ticket{} }
func (m *Ticket) String() string            { return proto.CompactTextString(m) }
func (*Ticket) ProtoMessage()               {}
func (*Ticket) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ticket) GetTicketId() string {
	if m != nil {
		return m.TicketId
	}
	return ""
}

func (m *Ticket) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Ticket) GetIsGenesis() bool {
	if m != nil {
		return m.IsGenesis
	}
	return false
}

func (m *Ticket) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Ticket) GetMinerTime() int64 {
	if m != nil {
		return m.MinerTime
	}
	return 0
}

func (m *Ticket) GetMinerValue() int64 {
	if m != nil {
		return m.MinerValue
	}
	return 0
}

func (m *Ticket) GetMinerAddress() string {
	if m != nil {
		return m.MinerAddress
	}
	return ""
}

func (m *Ticket) GetReturnAddress() string {
	if m != nil {
		return m.ReturnAddress
	}
	return ""
}

// message for execs.ticket
type TicketAction struct {
	// Types that are valid to be assigned to Value:
	//	*TicketAction_Tbind
	//	*TicketAction_Topen
	//	*TicketAction_Genesis
	//	*TicketAction_Tclose
	//	*TicketAction_Miner
	Value isTicketAction_Value `protobuf_oneof:"value"`
	Ty    int32                `protobuf:"varint,10,opt,name=ty" json:"ty,omitempty"`
}

func (m *TicketAction) Reset()                    { *m = TicketAction{} }
func (m *TicketAction) String() string            { return proto.CompactTextString(m) }
func (*TicketAction) ProtoMessage()               {}
func (*TicketAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isTicketAction_Value interface {
	isTicketAction_Value()
}

type TicketAction_Tbind struct {
	Tbind *TicketBind `protobuf:"bytes,5,opt,name=tbind,oneof"`
}
type TicketAction_Topen struct {
	Topen *TicketOpen `protobuf:"bytes,1,opt,name=topen,oneof"`
}
type TicketAction_Genesis struct {
	Genesis *TicketGenesis `protobuf:"bytes,2,opt,name=genesis,oneof"`
}
type TicketAction_Tclose struct {
	Tclose *TicketClose `protobuf:"bytes,3,opt,name=tclose,oneof"`
}
type TicketAction_Miner struct {
	Miner *TicketMiner `protobuf:"bytes,4,opt,name=miner,oneof"`
}

func (*TicketAction_Tbind) isTicketAction_Value()   {}
func (*TicketAction_Topen) isTicketAction_Value()   {}
func (*TicketAction_Genesis) isTicketAction_Value() {}
func (*TicketAction_Tclose) isTicketAction_Value()  {}
func (*TicketAction_Miner) isTicketAction_Value()   {}

func (m *TicketAction) GetValue() isTicketAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *TicketAction) GetTbind() *TicketBind {
	if x, ok := m.GetValue().(*TicketAction_Tbind); ok {
		return x.Tbind
	}
	return nil
}

func (m *TicketAction) GetTopen() *TicketOpen {
	if x, ok := m.GetValue().(*TicketAction_Topen); ok {
		return x.Topen
	}
	return nil
}

func (m *TicketAction) GetGenesis() *TicketGenesis {
	if x, ok := m.GetValue().(*TicketAction_Genesis); ok {
		return x.Genesis
	}
	return nil
}

func (m *TicketAction) GetTclose() *TicketClose {
	if x, ok := m.GetValue().(*TicketAction_Tclose); ok {
		return x.Tclose
	}
	return nil
}

func (m *TicketAction) GetMiner() *TicketMiner {
	if x, ok := m.GetValue().(*TicketAction_Miner); ok {
		return x.Miner
	}
	return nil
}

func (m *TicketAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TicketAction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TicketAction_OneofMarshaler, _TicketAction_OneofUnmarshaler, _TicketAction_OneofSizer, []interface{}{
		(*TicketAction_Tbind)(nil),
		(*TicketAction_Topen)(nil),
		(*TicketAction_Genesis)(nil),
		(*TicketAction_Tclose)(nil),
		(*TicketAction_Miner)(nil),
	}
}

func _TicketAction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TicketAction)
	// value
	switch x := m.Value.(type) {
	case *TicketAction_Tbind:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Tbind); err != nil {
			return err
		}
	case *TicketAction_Topen:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Topen); err != nil {
			return err
		}
	case *TicketAction_Genesis:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Genesis); err != nil {
			return err
		}
	case *TicketAction_Tclose:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Tclose); err != nil {
			return err
		}
	case *TicketAction_Miner:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Miner); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TicketAction.Value has unexpected type %T", x)
	}
	return nil
}

func _TicketAction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TicketAction)
	switch tag {
	case 5: // value.tbind
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TicketBind)
		err := b.DecodeMessage(msg)
		m.Value = &TicketAction_Tbind{msg}
		return true, err
	case 1: // value.topen
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TicketOpen)
		err := b.DecodeMessage(msg)
		m.Value = &TicketAction_Topen{msg}
		return true, err
	case 2: // value.genesis
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TicketGenesis)
		err := b.DecodeMessage(msg)
		m.Value = &TicketAction_Genesis{msg}
		return true, err
	case 3: // value.tclose
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TicketClose)
		err := b.DecodeMessage(msg)
		m.Value = &TicketAction_Tclose{msg}
		return true, err
	case 4: // value.miner
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TicketMiner)
		err := b.DecodeMessage(msg)
		m.Value = &TicketAction_Miner{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TicketAction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TicketAction)
	// value
	switch x := m.Value.(type) {
	case *TicketAction_Tbind:
		s := proto.Size(x.Tbind)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TicketAction_Topen:
		s := proto.Size(x.Topen)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TicketAction_Genesis:
		s := proto.Size(x.Genesis)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TicketAction_Tclose:
		s := proto.Size(x.Tclose)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TicketAction_Miner:
		s := proto.Size(x.Miner)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type TicketMiner struct {
	Bits     uint32 `protobuf:"varint,1,opt,name=bits" json:"bits,omitempty"`
	Reward   int64  `protobuf:"varint,2,opt,name=reward" json:"reward,omitempty"`
	TicketId string `protobuf:"bytes,3,opt,name=ticketId" json:"ticketId,omitempty"`
	Modify   []byte `protobuf:"bytes,4,opt,name=modify,proto3" json:"modify,omitempty"`
}

func (m *TicketMiner) Reset()                    { *m = TicketMiner{} }
func (m *TicketMiner) String() string            { return proto.CompactTextString(m) }
func (*TicketMiner) ProtoMessage()               {}
func (*TicketMiner) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TicketMiner) GetBits() uint32 {
	if m != nil {
		return m.Bits
	}
	return 0
}

func (m *TicketMiner) GetReward() int64 {
	if m != nil {
		return m.Reward
	}
	return 0
}

func (m *TicketMiner) GetTicketId() string {
	if m != nil {
		return m.TicketId
	}
	return ""
}

func (m *TicketMiner) GetModify() []byte {
	if m != nil {
		return m.Modify
	}
	return nil
}

type MinerFlag struct {
	Flag    int32 `protobuf:"varint,1,opt,name=flag" json:"flag,omitempty"`
	Reserve int64 `protobuf:"varint,2,opt,name=reserve" json:"reserve,omitempty"`
}

func (m *MinerFlag) Reset()                    { *m = MinerFlag{} }
func (m *MinerFlag) String() string            { return proto.CompactTextString(m) }
func (*MinerFlag) ProtoMessage()               {}
func (*MinerFlag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MinerFlag) GetFlag() int32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

func (m *MinerFlag) GetReserve() int64 {
	if m != nil {
		return m.Reserve
	}
	return 0
}

type TicketBind struct {
	MinerAddress  string `protobuf:"bytes,1,opt,name=minerAddress" json:"minerAddress,omitempty"`
	ReturnAddress string `protobuf:"bytes,2,opt,name=returnAddress" json:"returnAddress,omitempty"`
}

func (m *TicketBind) Reset()                    { *m = TicketBind{} }
func (m *TicketBind) String() string            { return proto.CompactTextString(m) }
func (*TicketBind) ProtoMessage()               {}
func (*TicketBind) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TicketBind) GetMinerAddress() string {
	if m != nil {
		return m.MinerAddress
	}
	return ""
}

func (m *TicketBind) GetReturnAddress() string {
	if m != nil {
		return m.ReturnAddress
	}
	return ""
}

type TicketOpen struct {
	// 用户挖矿的ticket 地址
	MinerAddress string `protobuf:"bytes,1,opt,name=minerAddress" json:"minerAddress,omitempty"`
	// 购买ticket的数目
	Count int32 `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	// 币实际存储的地址
	ReturnAddress string `protobuf:"bytes,3,opt,name=returnAddress" json:"returnAddress,omitempty"`
}

func (m *TicketOpen) Reset()                    { *m = TicketOpen{} }
func (m *TicketOpen) String() string            { return proto.CompactTextString(m) }
func (*TicketOpen) ProtoMessage()               {}
func (*TicketOpen) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TicketOpen) GetMinerAddress() string {
	if m != nil {
		return m.MinerAddress
	}
	return ""
}

func (m *TicketOpen) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *TicketOpen) GetReturnAddress() string {
	if m != nil {
		return m.ReturnAddress
	}
	return ""
}

type TicketGenesis struct {
	MinerAddress  string `protobuf:"bytes,1,opt,name=minerAddress" json:"minerAddress,omitempty"`
	ReturnAddress string `protobuf:"bytes,2,opt,name=returnAddress" json:"returnAddress,omitempty"`
	Count         int32  `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
}

func (m *TicketGenesis) Reset()                    { *m = TicketGenesis{} }
func (m *TicketGenesis) String() string            { return proto.CompactTextString(m) }
func (*TicketGenesis) ProtoMessage()               {}
func (*TicketGenesis) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *TicketGenesis) GetMinerAddress() string {
	if m != nil {
		return m.MinerAddress
	}
	return ""
}

func (m *TicketGenesis) GetReturnAddress() string {
	if m != nil {
		return m.ReturnAddress
	}
	return ""
}

func (m *TicketGenesis) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type TicketClose struct {
	TicketId []string `protobuf:"bytes,1,rep,name=ticketId" json:"ticketId,omitempty"`
}

func (m *TicketClose) Reset()                    { *m = TicketClose{} }
func (m *TicketClose) String() string            { return proto.CompactTextString(m) }
func (*TicketClose) ProtoMessage()               {}
func (*TicketClose) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *TicketClose) GetTicketId() []string {
	if m != nil {
		return m.TicketId
	}
	return nil
}

type TicketList struct {
	Addr   string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Status int32  `protobuf:"varint,3,opt,name=status" json:"status,omitempty"`
}

func (m *TicketList) Reset()                    { *m = TicketList{} }
func (m *TicketList) String() string            { return proto.CompactTextString(m) }
func (*TicketList) ProtoMessage()               {}
func (*TicketList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *TicketList) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *TicketList) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type TicketInfos struct {
	TicketIds []string `protobuf:"bytes,1,rep,name=ticketIds" json:"ticketIds,omitempty"`
}

func (m *TicketInfos) Reset()                    { *m = TicketInfos{} }
func (m *TicketInfos) String() string            { return proto.CompactTextString(m) }
func (*TicketInfos) ProtoMessage()               {}
func (*TicketInfos) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *TicketInfos) GetTicketIds() []string {
	if m != nil {
		return m.TicketIds
	}
	return nil
}

type ReplyTicketList struct {
	Tickets []*Ticket `protobuf:"bytes,1,rep,name=tickets" json:"tickets,omitempty"`
}

func (m *ReplyTicketList) Reset()                    { *m = ReplyTicketList{} }
func (m *ReplyTicketList) String() string            { return proto.CompactTextString(m) }
func (*ReplyTicketList) ProtoMessage()               {}
func (*ReplyTicketList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ReplyTicketList) GetTickets() []*Ticket {
	if m != nil {
		return m.Tickets
	}
	return nil
}

type ReplyWalletTickets struct {
	Tickets  []*Ticket `protobuf:"bytes,1,rep,name=tickets" json:"tickets,omitempty"`
	Privkeys [][]byte  `protobuf:"bytes,2,rep,name=privkeys,proto3" json:"privkeys,omitempty"`
}

func (m *ReplyWalletTickets) Reset()                    { *m = ReplyWalletTickets{} }
func (m *ReplyWalletTickets) String() string            { return proto.CompactTextString(m) }
func (*ReplyWalletTickets) ProtoMessage()               {}
func (*ReplyWalletTickets) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ReplyWalletTickets) GetTickets() []*Ticket {
	if m != nil {
		return m.Tickets
	}
	return nil
}

func (m *ReplyWalletTickets) GetPrivkeys() [][]byte {
	if m != nil {
		return m.Privkeys
	}
	return nil
}

type ReceiptTicket struct {
	TicketId   string `protobuf:"bytes,1,opt,name=ticketId" json:"ticketId,omitempty"`
	Status     int32  `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
	PrevStatus int32  `protobuf:"varint,3,opt,name=prevStatus" json:"prevStatus,omitempty"`
	Addr       string `protobuf:"bytes,4,opt,name=addr" json:"addr,omitempty"`
}

func (m *ReceiptTicket) Reset()                    { *m = ReceiptTicket{} }
func (m *ReceiptTicket) String() string            { return proto.CompactTextString(m) }
func (*ReceiptTicket) ProtoMessage()               {}
func (*ReceiptTicket) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *ReceiptTicket) GetTicketId() string {
	if m != nil {
		return m.TicketId
	}
	return ""
}

func (m *ReceiptTicket) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ReceiptTicket) GetPrevStatus() int32 {
	if m != nil {
		return m.PrevStatus
	}
	return 0
}

func (m *ReceiptTicket) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type ReceiptTicketBind struct {
	OldMinerAddress string `protobuf:"bytes,1,opt,name=oldMinerAddress" json:"oldMinerAddress,omitempty"`
	NewMinerAddress string `protobuf:"bytes,2,opt,name=newMinerAddress" json:"newMinerAddress,omitempty"`
	ReturnAddress   string `protobuf:"bytes,3,opt,name=returnAddress" json:"returnAddress,omitempty"`
}

func (m *ReceiptTicketBind) Reset()                    { *m = ReceiptTicketBind{} }
func (m *ReceiptTicketBind) String() string            { return proto.CompactTextString(m) }
func (*ReceiptTicketBind) ProtoMessage()               {}
func (*ReceiptTicketBind) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *ReceiptTicketBind) GetOldMinerAddress() string {
	if m != nil {
		return m.OldMinerAddress
	}
	return ""
}

func (m *ReceiptTicketBind) GetNewMinerAddress() string {
	if m != nil {
		return m.NewMinerAddress
	}
	return ""
}

func (m *ReceiptTicketBind) GetReturnAddress() string {
	if m != nil {
		return m.ReturnAddress
	}
	return ""
}

type ReqBindMiner struct {
	BindAddr     string `protobuf:"bytes,1,opt,name=bindAddr" json:"bindAddr,omitempty"`
	OriginAddr   string `protobuf:"bytes,2,opt,name=originAddr" json:"originAddr,omitempty"`
	Amount       int64  `protobuf:"varint,3,opt,name=amount" json:"amount,omitempty"`
	CheckBalance bool   `protobuf:"varint,4,opt,name=checkBalance" json:"checkBalance,omitempty"`
}

func (m *ReqBindMiner) Reset()                    { *m = ReqBindMiner{} }
func (m *ReqBindMiner) String() string            { return proto.CompactTextString(m) }
func (*ReqBindMiner) ProtoMessage()               {}
func (*ReqBindMiner) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *ReqBindMiner) GetBindAddr() string {
	if m != nil {
		return m.BindAddr
	}
	return ""
}

func (m *ReqBindMiner) GetOriginAddr() string {
	if m != nil {
		return m.OriginAddr
	}
	return ""
}

func (m *ReqBindMiner) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ReqBindMiner) GetCheckBalance() bool {
	if m != nil {
		return m.CheckBalance
	}
	return false
}

type ReplyBindMiner struct {
	TxHex string `protobuf:"bytes,1,opt,name=txHex" json:"txHex,omitempty"`
}

func (m *ReplyBindMiner) Reset()                    { *m = ReplyBindMiner{} }
func (m *ReplyBindMiner) String() string            { return proto.CompactTextString(m) }
func (*ReplyBindMiner) ProtoMessage()               {}
func (*ReplyBindMiner) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *ReplyBindMiner) GetTxHex() string {
	if m != nil {
		return m.TxHex
	}
	return ""
}

func init() {
	proto.RegisterType((*Ticket)(nil), "types.Ticket")
	proto.RegisterType((*TicketAction)(nil), "types.TicketAction")
	proto.RegisterType((*TicketMiner)(nil), "types.TicketMiner")
	proto.RegisterType((*MinerFlag)(nil), "types.MinerFlag")
	proto.RegisterType((*TicketBind)(nil), "types.TicketBind")
	proto.RegisterType((*TicketOpen)(nil), "types.TicketOpen")
	proto.RegisterType((*TicketGenesis)(nil), "types.TicketGenesis")
	proto.RegisterType((*TicketClose)(nil), "types.TicketClose")
	proto.RegisterType((*TicketList)(nil), "types.TicketList")
	proto.RegisterType((*TicketInfos)(nil), "types.TicketInfos")
	proto.RegisterType((*ReplyTicketList)(nil), "types.ReplyTicketList")
	proto.RegisterType((*ReplyWalletTickets)(nil), "types.ReplyWalletTickets")
	proto.RegisterType((*ReceiptTicket)(nil), "types.ReceiptTicket")
	proto.RegisterType((*ReceiptTicketBind)(nil), "types.ReceiptTicketBind")
	proto.RegisterType((*ReqBindMiner)(nil), "types.ReqBindMiner")
	proto.RegisterType((*ReplyBindMiner)(nil), "types.ReplyBindMiner")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Ticket service

type TicketClient interface {
	// 创建绑定挖矿
	CreateBindMiner(ctx context.Context, in *ReqBindMiner, opts ...grpc.CallOption) (*ReplyBindMiner, error)
	// 查询钱包票数
	GetTicketCount(ctx context.Context, in *types1.ReqNil, opts ...grpc.CallOption) (*types1.Int64, error)
	// Miner
	// 设置自动挖矿
	SetAutoMining(ctx context.Context, in *MinerFlag, opts ...grpc.CallOption) (*types1.Reply, error)
}

type ticketClient struct {
	cc *grpc.ClientConn
}

func NewTicketClient(cc *grpc.ClientConn) TicketClient {
	return &ticketClient{cc}
}

func (c *ticketClient) CreateBindMiner(ctx context.Context, in *ReqBindMiner, opts ...grpc.CallOption) (*ReplyBindMiner, error) {
	out := new(ReplyBindMiner)
	err := grpc.Invoke(ctx, "/types.ticket/CreateBindMiner", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketClient) GetTicketCount(ctx context.Context, in *types1.ReqNil, opts ...grpc.CallOption) (*types1.Int64, error) {
	out := new(types1.Int64)
	err := grpc.Invoke(ctx, "/types.ticket/GetTicketCount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketClient) SetAutoMining(ctx context.Context, in *MinerFlag, opts ...grpc.CallOption) (*types1.Reply, error) {
	out := new(types1.Reply)
	err := grpc.Invoke(ctx, "/types.ticket/SetAutoMining", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Ticket service

type TicketServer interface {
	// 创建绑定挖矿
	CreateBindMiner(context.Context, *ReqBindMiner) (*ReplyBindMiner, error)
	// 查询钱包票数
	GetTicketCount(context.Context, *types1.ReqNil) (*types1.Int64, error)
	// Miner
	// 设置自动挖矿
	SetAutoMining(context.Context, *MinerFlag) (*types1.Reply, error)
}

func RegisterTicketServer(s *grpc.Server, srv TicketServer) {
	s.RegisterService(&_Ticket_serviceDesc, srv)
}

func _Ticket_CreateBindMiner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqBindMiner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServer).CreateBindMiner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.ticket/CreateBindMiner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServer).CreateBindMiner(ctx, req.(*ReqBindMiner))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ticket_GetTicketCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types1.ReqNil)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServer).GetTicketCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.ticket/GetTicketCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServer).GetTicketCount(ctx, req.(*types1.ReqNil))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ticket_SetAutoMining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinerFlag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServer).SetAutoMining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.ticket/SetAutoMining",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServer).SetAutoMining(ctx, req.(*MinerFlag))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ticket_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.ticket",
	HandlerType: (*TicketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBindMiner",
			Handler:    _Ticket_CreateBindMiner_Handler,
		},
		{
			MethodName: "GetTicketCount",
			Handler:    _Ticket_GetTicketCount_Handler,
		},
		{
			MethodName: "SetAutoMining",
			Handler:    _Ticket_SetAutoMining_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ticket.proto",
}

func init() { proto.RegisterFile("ticket.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 779 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x55, 0xcf, 0x6f, 0xeb, 0x44,
	0x10, 0x8e, 0xe3, 0x3a, 0x49, 0x27, 0x4e, 0x4b, 0x97, 0x82, 0xac, 0x08, 0x55, 0xd1, 0x0a, 0x41,
	0x0a, 0xa8, 0x40, 0x40, 0x08, 0xb8, 0xa0, 0xb4, 0x12, 0x4d, 0x25, 0x0a, 0xd2, 0xb6, 0x2a, 0xe2,
	0xe8, 0xda, 0xdb, 0xb0, 0xaa, 0xb3, 0x36, 0xf6, 0x26, 0x6d, 0xae, 0x1c, 0x9e, 0xf4, 0x0e, 0xef,
	0x4f, 0x79, 0xff, 0xe3, 0xd3, 0x8e, 0xd7, 0xbf, 0x92, 0x1e, 0x22, 0xbd, 0x77, 0xf3, 0x7c, 0xfb,
	0x8d, 0x67, 0xf6, 0x9b, 0xcf, 0x63, 0x70, 0x95, 0x08, 0x1e, 0xb9, 0x3a, 0x4b, 0xd2, 0x58, 0xc5,
	0xc4, 0x51, 0xeb, 0x84, 0x67, 0x43, 0x37, 0x88, 0x17, 0x8b, 0x58, 0xe6, 0x20, 0xfd, 0xbf, 0x0d,
	0x9d, 0x5b, 0x64, 0x91, 0x21, 0xf4, 0x72, 0xfe, 0x55, 0xe8, 0x59, 0x23, 0x6b, 0xbc, 0xcf, 0xca,
	0x98, 0x7c, 0x0a, 0x9d, 0x4c, 0xf9, 0x6a, 0x99, 0x79, 0xed, 0x91, 0x35, 0x76, 0x98, 0x89, 0xc8,
	0x67, 0xb0, 0x2f, 0xb2, 0x4b, 0x2e, 0x79, 0x26, 0x32, 0xcf, 0x1e, 0x59, 0xe3, 0x1e, 0xab, 0x00,
	0x72, 0x02, 0x10, 0xa4, 0xdc, 0x57, 0xfc, 0x56, 0x2c, 0xb8, 0xb7, 0x37, 0xb2, 0xc6, 0x36, 0xab,
	0x21, 0x3a, 0x7b, 0x21, 0x24, 0x4f, 0xf1, 0xd8, 0xc1, 0xe3, 0x0a, 0xd0, 0xd9, 0x18, 0xdc, 0xf9,
	0xd1, 0x92, 0x7b, 0xbd, 0x3c, 0xbb, 0x42, 0x08, 0x05, 0x17, 0xa3, 0x69, 0x18, 0xa6, 0x3c, 0xcb,
	0xbc, 0x0e, 0xf6, 0xdc, 0xc0, 0xc8, 0xe7, 0x30, 0x48, 0xb9, 0x5a, 0xa6, 0xb2, 0x20, 0x75, 0x91,
	0xd4, 0x04, 0xe9, 0xeb, 0x36, 0xb8, 0xb9, 0x08, 0xd3, 0x40, 0x89, 0x58, 0x92, 0x53, 0x70, 0xd4,
	0xbd, 0x90, 0x21, 0x36, 0xd5, 0x9f, 0x1c, 0x9d, 0xa1, 0x74, 0x67, 0x39, 0xe7, 0x5c, 0xc8, 0x70,
	0xd6, 0x62, 0x39, 0x03, 0xa9, 0x71, 0xc2, 0x25, 0x4a, 0xb6, 0x49, 0xfd, 0x2b, 0xe1, 0x12, 0xa9,
	0x9a, 0x41, 0xbe, 0x83, 0xee, 0xdc, 0x48, 0xd5, 0x46, 0xf2, 0x71, 0x83, 0x6c, 0x54, 0x9b, 0xb5,
	0x58, 0x41, 0x23, 0xdf, 0x40, 0x47, 0x05, 0x51, 0x9c, 0x71, 0xd4, 0xb6, 0x3f, 0x21, 0x8d, 0x84,
	0x0b, 0x7d, 0x32, 0x6b, 0x31, 0xc3, 0x21, 0x5f, 0x81, 0x83, 0x97, 0x47, 0xa5, 0x37, 0xc9, 0xd7,
	0xfa, 0x44, 0xf7, 0x82, 0x14, 0x72, 0x00, 0x6d, 0xb5, 0xf6, 0x00, 0x87, 0xd9, 0x56, 0xeb, 0xf3,
	0x2e, 0x38, 0x2b, 0xad, 0x2a, 0x5d, 0x40, 0xbf, 0x96, 0x40, 0x08, 0xec, 0xdd, 0x0b, 0x95, 0xe1,
	0xed, 0x06, 0x0c, 0x9f, 0xb5, 0x19, 0x52, 0xfe, 0xe4, 0xa7, 0x21, 0x5e, 0xc3, 0x66, 0x26, 0x6a,
	0x18, 0xc8, 0xde, 0x36, 0xd0, 0x22, 0x0e, 0xc5, 0xc3, 0x1a, 0x9b, 0x73, 0x99, 0x89, 0xe8, 0x2f,
	0xb0, 0x8f, 0x85, 0x7e, 0x8f, 0xfc, 0xb9, 0x2e, 0xf6, 0x10, 0xf9, 0x73, 0x2c, 0xe6, 0x30, 0x7c,
	0x26, 0x1e, 0x74, 0x53, 0x9e, 0xf1, 0x74, 0xc5, 0x4d, 0xb5, 0x22, 0xa4, 0x77, 0x00, 0xd5, 0x40,
	0xb6, 0xdc, 0x60, 0xed, 0xe2, 0x86, 0xf6, 0x4b, 0x6e, 0x88, 0x8a, 0xf7, 0xea, 0xe9, 0xed, 0xf4,
	0xde, 0x63, 0x70, 0x82, 0x78, 0x29, 0x95, 0xf9, 0x38, 0xf2, 0x60, 0xbb, 0x9a, 0xfd, 0x52, 0xb5,
	0x18, 0x06, 0x8d, 0xf1, 0x7f, 0xb8, 0x8b, 0x54, 0x6d, 0xd9, 0xb5, 0xb6, 0xe8, 0x69, 0x31, 0x60,
	0xb4, 0xcf, 0xc6, 0x57, 0x6f, 0xd7, 0x87, 0x46, 0x7f, 0x2e, 0x94, 0xf8, 0x43, 0x64, 0x4a, 0x4f,
	0xc7, 0x0f, 0xc3, 0xd4, 0x34, 0x84, 0xcf, 0xb5, 0xbd, 0x60, 0xd7, 0xf7, 0x02, 0xfd, 0xba, 0x28,
	0x72, 0x25, 0x1f, 0x62, 0x5c, 0x13, 0xc5, 0x4b, 0x33, 0x53, 0xa5, 0x02, 0xe8, 0xaf, 0x70, 0xc8,
	0x78, 0x12, 0xad, 0x6b, 0xb5, 0xbe, 0x84, 0x6e, 0x7e, 0x9e, 0xd3, 0xfb, 0x93, 0x41, 0xc3, 0xcc,
	0xac, 0x38, 0xa5, 0xff, 0x00, 0xc1, 0xdc, 0xbf, 0xfd, 0x28, 0xe2, 0x2a, 0x3f, 0xcd, 0x76, 0x4e,
	0xd7, 0xb7, 0x4f, 0x52, 0xb1, 0x7a, 0xe4, 0x6b, 0xad, 0xa1, 0x3d, 0x76, 0x59, 0x19, 0xd3, 0x27,
	0x18, 0x30, 0x1e, 0x70, 0x91, 0xa8, 0xf7, 0x58, 0x90, 0x27, 0x00, 0x49, 0xca, 0x57, 0x37, 0x75,
	0x91, 0x6a, 0x48, 0x29, 0xea, 0x5e, 0x25, 0x2a, 0x7d, 0x63, 0xc1, 0x51, 0xa3, 0x32, 0x1a, 0x7c,
	0x0c, 0x87, 0x71, 0x14, 0x5e, 0x6f, 0x5b, 0x63, 0x13, 0xd6, 0x4c, 0xc9, 0x9f, 0x1a, 0xcc, 0xdc,
	0x1f, 0x9b, 0xf0, 0x8e, 0x16, 0x7d, 0x65, 0x81, 0xcb, 0xf8, 0x7f, 0xba, 0x8b, 0x7c, 0x29, 0x0c,
	0xa1, 0xa7, 0x77, 0xdf, 0xb4, 0x72, 0x43, 0x19, 0xeb, 0x0b, 0xc7, 0xa9, 0x98, 0x0b, 0xcc, 0x36,
	0x75, 0x6b, 0x88, 0x16, 0xca, 0x5f, 0x94, 0xae, 0xb4, 0x99, 0x89, 0xb4, 0xed, 0x83, 0x7f, 0x79,
	0xf0, 0x78, 0xee, 0x47, 0xbe, 0x0c, 0xf2, 0xbf, 0x45, 0x8f, 0x35, 0x30, 0xfa, 0x05, 0x1c, 0xe0,
	0xb0, 0xab, 0x4e, 0x8e, 0xc1, 0x51, 0xcf, 0x33, 0xfe, 0x6c, 0xda, 0xc8, 0x83, 0xc9, 0x5b, 0x0b,
	0x3a, 0xf9, 0x64, 0xc8, 0x6f, 0x70, 0x78, 0x81, 0x3f, 0x9c, 0x2a, 0xe7, 0x63, 0xe3, 0x85, 0xfa,
	0x95, 0x86, 0x9f, 0x94, 0x60, 0xfd, 0xfd, 0xb4, 0x45, 0xbe, 0x85, 0x83, 0xcb, 0xc2, 0x58, 0x17,
	0xd8, 0xe9, 0xa0, 0xca, 0xff, 0x53, 0x44, 0x43, 0xd7, 0x84, 0x57, 0x52, 0xfd, 0xf4, 0x23, 0x6d,
	0x91, 0xef, 0x61, 0x70, 0xc3, 0xd5, 0x74, 0xa9, 0xe2, 0x6b, 0x21, 0x85, 0x9c, 0x93, 0x8f, 0x0c,
	0xa1, 0xdc, 0x73, 0x65, 0x0a, 0x16, 0xa3, 0xad, 0xfb, 0x0e, 0xfe, 0x8b, 0x7f, 0x78, 0x17, 0x00,
	0x00, 0xff, 0xff, 0x0e, 0xb0, 0xda, 0x90, 0xb0, 0x07, 0x00, 0x00,
}
