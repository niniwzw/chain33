// Code generated by protoc-gen-go.
// source: pokerbull.proto
// DO NOT EDIT!

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	pokerbull.proto

It has these top-level messages:
	PokerBull
	PBHand
	PBPlayer
	PBResult
	PBPoker
	PBGameAction
	PBGameStart
	PBGameContinue
	PBGameQuit
	PBGameQuery
	QueryPBGameListByStatusAndPlayerNum
	PBGameRecord
	QueryPBGameInfo
	ReplyPBGame
	QueryPBGameInfos
	ReplyPBGameList
	ReceiptPBGame
	PBStartTxReq
	PBContinueTxReq
	PBQuitTxReq
	PBQueryReq
*/
package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import types2 "gitlab.33.cn/chain33/chain33/types"

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

// 斗牛游戏内容
type PokerBull struct {
	GameId      string      `protobuf:"bytes,1,opt,name=gameId" json:"gameId,omitempty"`
	Status      int32       `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
	StartTime   int64       `protobuf:"varint,3,opt,name=startTime" json:"startTime,omitempty"`
	StartTxHash string      `protobuf:"bytes,4,opt,name=startTxHash" json:"startTxHash,omitempty"`
	Value       int64       `protobuf:"varint,5,opt,name=value" json:"value,omitempty"`
	Poker       *PBPoker    `protobuf:"bytes,6,opt,name=poker" json:"poker,omitempty"`
	Players     []*PBPlayer `protobuf:"bytes,7,rep,name=players" json:"players,omitempty"`
	PlayerNum   int32       `protobuf:"varint,8,opt,name=playerNum" json:"playerNum,omitempty"`
	Results     []*PBResult `protobuf:"bytes,9,rep,name=results" json:"results,omitempty"`
	Index       int64       `protobuf:"varint,10,opt,name=index" json:"index,omitempty"`
	PrevIndex   int64       `protobuf:"varint,11,opt,name=prevIndex" json:"prevIndex,omitempty"`
	QuitTime    int64       `protobuf:"varint,12,opt,name=quitTime" json:"quitTime,omitempty"`
	QuitTxHash  string      `protobuf:"bytes,13,opt,name=quitTxHash" json:"quitTxHash,omitempty"`
	DealerAddr  string      `protobuf:"bytes,14,opt,name=dealerAddr" json:"dealerAddr,omitempty"`
	IsWaiting   bool        `protobuf:"varint,15,opt,name=isWaiting" json:"isWaiting,omitempty"`
}

func (m *PokerBull) Reset()                    { *m = PokerBull{} }
func (m *PokerBull) String() string            { return proto.CompactTextString(m) }
func (*PokerBull) ProtoMessage()               {}
func (*PokerBull) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PokerBull) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

func (m *PokerBull) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *PokerBull) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *PokerBull) GetStartTxHash() string {
	if m != nil {
		return m.StartTxHash
	}
	return ""
}

func (m *PokerBull) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *PokerBull) GetPoker() *PBPoker {
	if m != nil {
		return m.Poker
	}
	return nil
}

func (m *PokerBull) GetPlayers() []*PBPlayer {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *PokerBull) GetPlayerNum() int32 {
	if m != nil {
		return m.PlayerNum
	}
	return 0
}

func (m *PokerBull) GetResults() []*PBResult {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *PokerBull) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *PokerBull) GetPrevIndex() int64 {
	if m != nil {
		return m.PrevIndex
	}
	return 0
}

func (m *PokerBull) GetQuitTime() int64 {
	if m != nil {
		return m.QuitTime
	}
	return 0
}

func (m *PokerBull) GetQuitTxHash() string {
	if m != nil {
		return m.QuitTxHash
	}
	return ""
}

func (m *PokerBull) GetDealerAddr() string {
	if m != nil {
		return m.DealerAddr
	}
	return ""
}

func (m *PokerBull) GetIsWaiting() bool {
	if m != nil {
		return m.IsWaiting
	}
	return false
}

// 一把牌
type PBHand struct {
	Cards    []int32 `protobuf:"varint,1,rep,packed,name=cards" json:"cards,omitempty"`
	Result   int32   `protobuf:"varint,2,opt,name=result" json:"result,omitempty"`
	Address  string  `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
	IsWin    bool    `protobuf:"varint,4,opt,name=isWin" json:"isWin,omitempty"`
	Leverage int32   `protobuf:"varint,5,opt,name=leverage" json:"leverage,omitempty"`
}

func (m *PBHand) Reset()                    { *m = PBHand{} }
func (m *PBHand) String() string            { return proto.CompactTextString(m) }
func (*PBHand) ProtoMessage()               {}
func (*PBHand) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PBHand) GetCards() []int32 {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *PBHand) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *PBHand) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *PBHand) GetIsWin() bool {
	if m != nil {
		return m.IsWin
	}
	return false
}

func (m *PBHand) GetLeverage() int32 {
	if m != nil {
		return m.Leverage
	}
	return 0
}

// 玩家
type PBPlayer struct {
	Hands   []*PBHand `protobuf:"bytes,1,rep,name=hands" json:"hands,omitempty"`
	Address string    `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	TxHash  int64     `protobuf:"varint,3,opt,name=txHash" json:"txHash,omitempty"`
	Ready   bool      `protobuf:"varint,4,opt,name=ready" json:"ready,omitempty"`
}

func (m *PBPlayer) Reset()                    { *m = PBPlayer{} }
func (m *PBPlayer) String() string            { return proto.CompactTextString(m) }
func (*PBPlayer) ProtoMessage()               {}
func (*PBPlayer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PBPlayer) GetHands() []*PBHand {
	if m != nil {
		return m.Hands
	}
	return nil
}

func (m *PBPlayer) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *PBPlayer) GetTxHash() int64 {
	if m != nil {
		return m.TxHash
	}
	return 0
}

func (m *PBPlayer) GetReady() bool {
	if m != nil {
		return m.Ready
	}
	return false
}

// 本局游戏结果
type PBResult struct {
	Hands          []*PBHand `protobuf:"bytes,1,rep,name=hands" json:"hands,omitempty"`
	Winner         string    `protobuf:"bytes,2,opt,name=winner" json:"winner,omitempty"`
	Leverage       int32     `protobuf:"varint,3,opt,name=leverage" json:"leverage,omitempty"`
	Dealer         string    `protobuf:"bytes,4,opt,name=dealer" json:"dealer,omitempty"`
	DealerLeverage int32     `protobuf:"varint,5,opt,name=dealerLeverage" json:"dealerLeverage,omitempty"`
}

func (m *PBResult) Reset()                    { *m = PBResult{} }
func (m *PBResult) String() string            { return proto.CompactTextString(m) }
func (*PBResult) ProtoMessage()               {}
func (*PBResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PBResult) GetHands() []*PBHand {
	if m != nil {
		return m.Hands
	}
	return nil
}

func (m *PBResult) GetWinner() string {
	if m != nil {
		return m.Winner
	}
	return ""
}

func (m *PBResult) GetLeverage() int32 {
	if m != nil {
		return m.Leverage
	}
	return 0
}

func (m *PBResult) GetDealer() string {
	if m != nil {
		return m.Dealer
	}
	return ""
}

func (m *PBResult) GetDealerLeverage() int32 {
	if m != nil {
		return m.DealerLeverage
	}
	return 0
}

// 扑克牌
type PBPoker struct {
	Cards   []int32 `protobuf:"varint,1,rep,packed,name=cards" json:"cards,omitempty"`
	Pointer int32   `protobuf:"varint,2,opt,name=pointer" json:"pointer,omitempty"`
}

func (m *PBPoker) Reset()                    { *m = PBPoker{} }
func (m *PBPoker) String() string            { return proto.CompactTextString(m) }
func (*PBPoker) ProtoMessage()               {}
func (*PBPoker) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PBPoker) GetCards() []int32 {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *PBPoker) GetPointer() int32 {
	if m != nil {
		return m.Pointer
	}
	return 0
}

// 游戏状态
type PBGameAction struct {
	// Types that are valid to be assigned to Value:
	//	*PBGameAction_Start
	//	*PBGameAction_Continue
	//	*PBGameAction_Quit
	//	*PBGameAction_Query
	Value isPBGameAction_Value `protobuf_oneof:"value"`
	Ty    int32                `protobuf:"varint,10,opt,name=ty" json:"ty,omitempty"`
}

func (m *PBGameAction) Reset()                    { *m = PBGameAction{} }
func (m *PBGameAction) String() string            { return proto.CompactTextString(m) }
func (*PBGameAction) ProtoMessage()               {}
func (*PBGameAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type isPBGameAction_Value interface {
	isPBGameAction_Value()
}

type PBGameAction_Start struct {
	Start *PBGameStart `protobuf:"bytes,1,opt,name=start,oneof"`
}
type PBGameAction_Continue struct {
	Continue *PBGameContinue `protobuf:"bytes,2,opt,name=continue,oneof"`
}
type PBGameAction_Quit struct {
	Quit *PBGameQuit `protobuf:"bytes,3,opt,name=quit,oneof"`
}
type PBGameAction_Query struct {
	Query *PBGameQuery `protobuf:"bytes,4,opt,name=query,oneof"`
}

func (*PBGameAction_Start) isPBGameAction_Value()    {}
func (*PBGameAction_Continue) isPBGameAction_Value() {}
func (*PBGameAction_Quit) isPBGameAction_Value()     {}
func (*PBGameAction_Query) isPBGameAction_Value()    {}

func (m *PBGameAction) GetValue() isPBGameAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *PBGameAction) GetStart() *PBGameStart {
	if x, ok := m.GetValue().(*PBGameAction_Start); ok {
		return x.Start
	}
	return nil
}

func (m *PBGameAction) GetContinue() *PBGameContinue {
	if x, ok := m.GetValue().(*PBGameAction_Continue); ok {
		return x.Continue
	}
	return nil
}

func (m *PBGameAction) GetQuit() *PBGameQuit {
	if x, ok := m.GetValue().(*PBGameAction_Quit); ok {
		return x.Quit
	}
	return nil
}

func (m *PBGameAction) GetQuery() *PBGameQuery {
	if x, ok := m.GetValue().(*PBGameAction_Query); ok {
		return x.Query
	}
	return nil
}

func (m *PBGameAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PBGameAction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PBGameAction_OneofMarshaler, _PBGameAction_OneofUnmarshaler, _PBGameAction_OneofSizer, []interface{}{
		(*PBGameAction_Start)(nil),
		(*PBGameAction_Continue)(nil),
		(*PBGameAction_Quit)(nil),
		(*PBGameAction_Query)(nil),
	}
}

func _PBGameAction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PBGameAction)
	// value
	switch x := m.Value.(type) {
	case *PBGameAction_Start:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Start); err != nil {
			return err
		}
	case *PBGameAction_Continue:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Continue); err != nil {
			return err
		}
	case *PBGameAction_Quit:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Quit); err != nil {
			return err
		}
	case *PBGameAction_Query:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Query); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("PBGameAction.Value has unexpected type %T", x)
	}
	return nil
}

func _PBGameAction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PBGameAction)
	switch tag {
	case 1: // value.start
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PBGameStart)
		err := b.DecodeMessage(msg)
		m.Value = &PBGameAction_Start{msg}
		return true, err
	case 2: // value.continue
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PBGameContinue)
		err := b.DecodeMessage(msg)
		m.Value = &PBGameAction_Continue{msg}
		return true, err
	case 3: // value.quit
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PBGameQuit)
		err := b.DecodeMessage(msg)
		m.Value = &PBGameAction_Quit{msg}
		return true, err
	case 4: // value.query
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PBGameQuery)
		err := b.DecodeMessage(msg)
		m.Value = &PBGameAction_Query{msg}
		return true, err
	default:
		return false, nil
	}
}

func _PBGameAction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PBGameAction)
	// value
	switch x := m.Value.(type) {
	case *PBGameAction_Start:
		s := proto.Size(x.Start)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PBGameAction_Continue:
		s := proto.Size(x.Continue)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PBGameAction_Quit:
		s := proto.Size(x.Quit)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PBGameAction_Query:
		s := proto.Size(x.Query)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// 游戏启动
type PBGameStart struct {
	Value     int64 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	PlayerNum int32 `protobuf:"varint,2,opt,name=playerNum" json:"playerNum,omitempty"`
}

func (m *PBGameStart) Reset()                    { *m = PBGameStart{} }
func (m *PBGameStart) String() string            { return proto.CompactTextString(m) }
func (*PBGameStart) ProtoMessage()               {}
func (*PBGameStart) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PBGameStart) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *PBGameStart) GetPlayerNum() int32 {
	if m != nil {
		return m.PlayerNum
	}
	return 0
}

// 游戏继续
type PBGameContinue struct {
	GameId string `protobuf:"bytes,1,opt,name=gameId" json:"gameId,omitempty"`
}

func (m *PBGameContinue) Reset()                    { *m = PBGameContinue{} }
func (m *PBGameContinue) String() string            { return proto.CompactTextString(m) }
func (*PBGameContinue) ProtoMessage()               {}
func (*PBGameContinue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PBGameContinue) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

// 游戏结束
type PBGameQuit struct {
	GameId string `protobuf:"bytes,1,opt,name=gameId" json:"gameId,omitempty"`
}

func (m *PBGameQuit) Reset()                    { *m = PBGameQuit{} }
func (m *PBGameQuit) String() string            { return proto.CompactTextString(m) }
func (*PBGameQuit) ProtoMessage()               {}
func (*PBGameQuit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *PBGameQuit) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

// 查询游戏结果
type PBGameQuery struct {
	GameId string `protobuf:"bytes,1,opt,name=gameId" json:"gameId,omitempty"`
}

func (m *PBGameQuery) Reset()                    { *m = PBGameQuery{} }
func (m *PBGameQuery) String() string            { return proto.CompactTextString(m) }
func (*PBGameQuery) ProtoMessage()               {}
func (*PBGameQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *PBGameQuery) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

// 根据状态和游戏人数查找
type QueryPBGameListByStatusAndPlayerNum struct {
	Status    int32 `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	PlayerNum int32 `protobuf:"varint,2,opt,name=playerNum" json:"playerNum,omitempty"`
	Index     int64 `protobuf:"varint,3,opt,name=index" json:"index,omitempty"`
}

func (m *QueryPBGameListByStatusAndPlayerNum) Reset()         { *m = QueryPBGameListByStatusAndPlayerNum{} }
func (m *QueryPBGameListByStatusAndPlayerNum) String() string { return proto.CompactTextString(m) }
func (*QueryPBGameListByStatusAndPlayerNum) ProtoMessage()    {}
func (*QueryPBGameListByStatusAndPlayerNum) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{10}
}

func (m *QueryPBGameListByStatusAndPlayerNum) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *QueryPBGameListByStatusAndPlayerNum) GetPlayerNum() int32 {
	if m != nil {
		return m.PlayerNum
	}
	return 0
}

func (m *QueryPBGameListByStatusAndPlayerNum) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

// 索引value值
type PBGameRecord struct {
	GameId string `protobuf:"bytes,1,opt,name=gameId" json:"gameId,omitempty"`
}

func (m *PBGameRecord) Reset()                    { *m = PBGameRecord{} }
func (m *PBGameRecord) String() string            { return proto.CompactTextString(m) }
func (*PBGameRecord) ProtoMessage()               {}
func (*PBGameRecord) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *PBGameRecord) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

type QueryPBGameInfo struct {
	GameId string `protobuf:"bytes,1,opt,name=gameId" json:"gameId,omitempty"`
}

func (m *QueryPBGameInfo) Reset()                    { *m = QueryPBGameInfo{} }
func (m *QueryPBGameInfo) String() string            { return proto.CompactTextString(m) }
func (*QueryPBGameInfo) ProtoMessage()               {}
func (*QueryPBGameInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *QueryPBGameInfo) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

type ReplyPBGame struct {
	Game *PokerBull `protobuf:"bytes,1,opt,name=game" json:"game,omitempty"`
}

func (m *ReplyPBGame) Reset()                    { *m = ReplyPBGame{} }
func (m *ReplyPBGame) String() string            { return proto.CompactTextString(m) }
func (*ReplyPBGame) ProtoMessage()               {}
func (*ReplyPBGame) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *ReplyPBGame) GetGame() *PokerBull {
	if m != nil {
		return m.Game
	}
	return nil
}

type QueryPBGameInfos struct {
	GameIds []string `protobuf:"bytes,1,rep,name=gameIds" json:"gameIds,omitempty"`
}

func (m *QueryPBGameInfos) Reset()                    { *m = QueryPBGameInfos{} }
func (m *QueryPBGameInfos) String() string            { return proto.CompactTextString(m) }
func (*QueryPBGameInfos) ProtoMessage()               {}
func (*QueryPBGameInfos) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *QueryPBGameInfos) GetGameIds() []string {
	if m != nil {
		return m.GameIds
	}
	return nil
}

type ReplyPBGameList struct {
	Games []*PokerBull `protobuf:"bytes,1,rep,name=games" json:"games,omitempty"`
}

func (m *ReplyPBGameList) Reset()                    { *m = ReplyPBGameList{} }
func (m *ReplyPBGameList) String() string            { return proto.CompactTextString(m) }
func (*ReplyPBGameList) ProtoMessage()               {}
func (*ReplyPBGameList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *ReplyPBGameList) GetGames() []*PokerBull {
	if m != nil {
		return m.Games
	}
	return nil
}

type ReceiptPBGame struct {
	GameId    string `protobuf:"bytes,1,opt,name=gameId" json:"gameId,omitempty"`
	Status    int32  `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
	Addr      string `protobuf:"bytes,4,opt,name=addr" json:"addr,omitempty"`
	Index     int64  `protobuf:"varint,5,opt,name=index" json:"index,omitempty"`
	PrevIndex int64  `protobuf:"varint,6,opt,name=prevIndex" json:"prevIndex,omitempty"`
	PlayerNum int32  `protobuf:"varint,7,opt,name=playerNum" json:"playerNum,omitempty"`
}

func (m *ReceiptPBGame) Reset()                    { *m = ReceiptPBGame{} }
func (m *ReceiptPBGame) String() string            { return proto.CompactTextString(m) }
func (*ReceiptPBGame) ProtoMessage()               {}
func (*ReceiptPBGame) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *ReceiptPBGame) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

func (m *ReceiptPBGame) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ReceiptPBGame) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *ReceiptPBGame) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ReceiptPBGame) GetPrevIndex() int64 {
	if m != nil {
		return m.PrevIndex
	}
	return 0
}

func (m *ReceiptPBGame) GetPlayerNum() int32 {
	if m != nil {
		return m.PlayerNum
	}
	return 0
}

type PBStartTxReq struct {
	Value     int64 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	PlayerNum int32 `protobuf:"varint,2,opt,name=playerNum" json:"playerNum,omitempty"`
	Fee       int64 `protobuf:"varint,3,opt,name=fee" json:"fee,omitempty"`
}

func (m *PBStartTxReq) Reset()                    { *m = PBStartTxReq{} }
func (m *PBStartTxReq) String() string            { return proto.CompactTextString(m) }
func (*PBStartTxReq) ProtoMessage()               {}
func (*PBStartTxReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *PBStartTxReq) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *PBStartTxReq) GetPlayerNum() int32 {
	if m != nil {
		return m.PlayerNum
	}
	return 0
}

func (m *PBStartTxReq) GetFee() int64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

type PBContinueTxReq struct {
	GameId string `protobuf:"bytes,1,opt,name=gameId" json:"gameId,omitempty"`
	Fee    int64  `protobuf:"varint,2,opt,name=fee" json:"fee,omitempty"`
}

func (m *PBContinueTxReq) Reset()                    { *m = PBContinueTxReq{} }
func (m *PBContinueTxReq) String() string            { return proto.CompactTextString(m) }
func (*PBContinueTxReq) ProtoMessage()               {}
func (*PBContinueTxReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *PBContinueTxReq) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

func (m *PBContinueTxReq) GetFee() int64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

type PBQuitTxReq struct {
	GameId string `protobuf:"bytes,1,opt,name=gameId" json:"gameId,omitempty"`
	Fee    int64  `protobuf:"varint,2,opt,name=fee" json:"fee,omitempty"`
}

func (m *PBQuitTxReq) Reset()                    { *m = PBQuitTxReq{} }
func (m *PBQuitTxReq) String() string            { return proto.CompactTextString(m) }
func (*PBQuitTxReq) ProtoMessage()               {}
func (*PBQuitTxReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *PBQuitTxReq) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

func (m *PBQuitTxReq) GetFee() int64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

type PBQueryReq struct {
	GameId string `protobuf:"bytes,1,opt,name=gameId" json:"gameId,omitempty"`
	Fee    int64  `protobuf:"varint,2,opt,name=fee" json:"fee,omitempty"`
}

func (m *PBQueryReq) Reset()                    { *m = PBQueryReq{} }
func (m *PBQueryReq) String() string            { return proto.CompactTextString(m) }
func (*PBQueryReq) ProtoMessage()               {}
func (*PBQueryReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *PBQueryReq) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

func (m *PBQueryReq) GetFee() int64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func init() {
	proto.RegisterType((*PokerBull)(nil), "types.PokerBull")
	proto.RegisterType((*PBHand)(nil), "types.PBHand")
	proto.RegisterType((*PBPlayer)(nil), "types.PBPlayer")
	proto.RegisterType((*PBResult)(nil), "types.PBResult")
	proto.RegisterType((*PBPoker)(nil), "types.PBPoker")
	proto.RegisterType((*PBGameAction)(nil), "types.PBGameAction")
	proto.RegisterType((*PBGameStart)(nil), "types.PBGameStart")
	proto.RegisterType((*PBGameContinue)(nil), "types.PBGameContinue")
	proto.RegisterType((*PBGameQuit)(nil), "types.PBGameQuit")
	proto.RegisterType((*PBGameQuery)(nil), "types.PBGameQuery")
	proto.RegisterType((*QueryPBGameListByStatusAndPlayerNum)(nil), "types.QueryPBGameListByStatusAndPlayerNum")
	proto.RegisterType((*PBGameRecord)(nil), "types.PBGameRecord")
	proto.RegisterType((*QueryPBGameInfo)(nil), "types.QueryPBGameInfo")
	proto.RegisterType((*ReplyPBGame)(nil), "types.ReplyPBGame")
	proto.RegisterType((*QueryPBGameInfos)(nil), "types.QueryPBGameInfos")
	proto.RegisterType((*ReplyPBGameList)(nil), "types.ReplyPBGameList")
	proto.RegisterType((*ReceiptPBGame)(nil), "types.ReceiptPBGame")
	proto.RegisterType((*PBStartTxReq)(nil), "types.PBStartTxReq")
	proto.RegisterType((*PBContinueTxReq)(nil), "types.PBContinueTxReq")
	proto.RegisterType((*PBQuitTxReq)(nil), "types.PBQuitTxReq")
	proto.RegisterType((*PBQueryReq)(nil), "types.PBQueryReq")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Pokerbull service

type PokerbullClient interface {
	// 游戏开始
	Start(ctx context.Context, in *PBGameStart, opts ...grpc.CallOption) (*types2.UnsignTx, error)
	// 游戏继续
	Continue(ctx context.Context, in *PBGameContinue, opts ...grpc.CallOption) (*types2.UnsignTx, error)
	// 游戏结束
	Quit(ctx context.Context, in *PBGameQuit, opts ...grpc.CallOption) (*types2.UnsignTx, error)
}

type pokerbullClient struct {
	cc *grpc.ClientConn
}

func NewPokerbullClient(cc *grpc.ClientConn) PokerbullClient {
	return &pokerbullClient{cc}
}

func (c *pokerbullClient) Start(ctx context.Context, in *PBGameStart, opts ...grpc.CallOption) (*types2.UnsignTx, error) {
	out := new(types2.UnsignTx)
	err := grpc.Invoke(ctx, "/types.pokerbull/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokerbullClient) Continue(ctx context.Context, in *PBGameContinue, opts ...grpc.CallOption) (*types2.UnsignTx, error) {
	out := new(types2.UnsignTx)
	err := grpc.Invoke(ctx, "/types.pokerbull/Continue", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokerbullClient) Quit(ctx context.Context, in *PBGameQuit, opts ...grpc.CallOption) (*types2.UnsignTx, error) {
	out := new(types2.UnsignTx)
	err := grpc.Invoke(ctx, "/types.pokerbull/Quit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Pokerbull service

type PokerbullServer interface {
	// 游戏开始
	Start(context.Context, *PBGameStart) (*types2.UnsignTx, error)
	// 游戏继续
	Continue(context.Context, *PBGameContinue) (*types2.UnsignTx, error)
	// 游戏结束
	Quit(context.Context, *PBGameQuit) (*types2.UnsignTx, error)
}

func RegisterPokerbullServer(s *grpc.Server, srv PokerbullServer) {
	s.RegisterService(&_Pokerbull_serviceDesc, srv)
}

func _Pokerbull_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PBGameStart)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokerbullServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.pokerbull/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokerbullServer).Start(ctx, req.(*PBGameStart))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pokerbull_Continue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PBGameContinue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokerbullServer).Continue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.pokerbull/Continue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokerbullServer).Continue(ctx, req.(*PBGameContinue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pokerbull_Quit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PBGameQuit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokerbullServer).Quit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.pokerbull/Quit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokerbullServer).Quit(ctx, req.(*PBGameQuit))
	}
	return interceptor(ctx, in, info, handler)
}

var _Pokerbull_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.pokerbull",
	HandlerType: (*PokerbullServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _Pokerbull_Start_Handler,
		},
		{
			MethodName: "Continue",
			Handler:    _Pokerbull_Continue_Handler,
		},
		{
			MethodName: "Quit",
			Handler:    _Pokerbull_Quit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pokerbull.proto",
}

func init() { proto.RegisterFile("pokerbull.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 898 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0x4f, 0x6f, 0xdb, 0x36,
	0x14, 0x8f, 0x6c, 0xcb, 0x7f, 0x9e, 0x1a, 0x3b, 0x25, 0xb6, 0x40, 0x08, 0x86, 0xc1, 0x50, 0xb3,
	0xcc, 0x1d, 0x8a, 0x1c, 0x9c, 0x61, 0x43, 0xb1, 0x53, 0xbc, 0xc3, 0x12, 0xa0, 0x18, 0x1c, 0x26,
	0x43, 0xcf, 0xac, 0xc5, 0xa6, 0xc2, 0x14, 0xda, 0x21, 0xa9, 0x34, 0xbe, 0xee, 0x8b, 0xec, 0x30,
	0xec, 0x43, 0xed, 0xb0, 0xef, 0x32, 0xf0, 0x91, 0x94, 0x2c, 0xd7, 0x02, 0x96, 0xde, 0xf4, 0x7b,
	0xef, 0xf1, 0xf1, 0xf7, 0xfe, 0x52, 0x30, 0x5a, 0x2d, 0x7f, 0xe7, 0xf2, 0x5d, 0x91, 0xe7, 0xa7,
	0x2b, 0xb9, 0xd4, 0x4b, 0x12, 0xea, 0xf5, 0x8a, 0xab, 0xa3, 0xe7, 0x5a, 0x32, 0xa1, 0xd8, 0x42,
	0x67, 0x4b, 0x61, 0x35, 0xc9, 0x3f, 0x6d, 0x18, 0xcc, 0x8d, 0xf5, 0xac, 0xc8, 0x73, 0x72, 0x08,
	0xdd, 0x5b, 0x76, 0xc7, 0x2f, 0xd3, 0x38, 0x18, 0x07, 0x93, 0x01, 0x75, 0xc8, 0xc8, 0x95, 0x66,
	0xba, 0x50, 0x71, 0x6b, 0x1c, 0x4c, 0x42, 0xea, 0x10, 0xf9, 0x0a, 0x06, 0x4a, 0x33, 0xa9, 0x6f,
	0xb2, 0x3b, 0x1e, 0xb7, 0xc7, 0xc1, 0xa4, 0x4d, 0x2b, 0x01, 0x19, 0x43, 0x64, 0xc1, 0xe3, 0x05,
	0x53, 0x1f, 0xe2, 0x0e, 0xba, 0xdc, 0x14, 0x91, 0x2f, 0x20, 0x7c, 0x60, 0x79, 0xc1, 0xe3, 0x10,
	0xcf, 0x5a, 0x40, 0x8e, 0x21, 0xc4, 0x00, 0xe2, 0xee, 0x38, 0x98, 0x44, 0xd3, 0xe1, 0x29, 0xb2,
	0x3f, 0x9d, 0xcf, 0x90, 0x28, 0xb5, 0x4a, 0xf2, 0x12, 0x7a, 0xab, 0x9c, 0xad, 0xb9, 0x54, 0x71,
	0x6f, 0xdc, 0x9e, 0x44, 0xd3, 0x51, 0x65, 0x87, 0x72, 0xea, 0xf5, 0x86, 0xa6, 0xfd, 0xfc, 0xb5,
	0xb8, 0x8b, 0xfb, 0x18, 0x41, 0x25, 0x30, 0x8e, 0x24, 0x57, 0x45, 0xae, 0x55, 0x3c, 0xd8, 0x72,
	0x44, 0x51, 0x4e, 0xbd, 0xde, 0xf0, 0xcd, 0x44, 0xca, 0x1f, 0x63, 0xb0, 0x7c, 0x11, 0xa0, 0x7b,
	0xc9, 0x1f, 0x2e, 0x51, 0x13, 0xd9, 0x2c, 0x94, 0x02, 0x72, 0x04, 0xfd, 0xfb, 0x22, 0xb3, 0x29,
	0x7a, 0x86, 0xca, 0x12, 0x93, 0xaf, 0x01, 0xf0, 0xdb, 0x26, 0x68, 0x1f, 0x13, 0xb4, 0x21, 0x31,
	0xfa, 0x94, 0xb3, 0x9c, 0xcb, 0xf3, 0x34, 0x95, 0xf1, 0xd0, 0xea, 0x2b, 0x89, 0xb9, 0x39, 0x53,
	0x6f, 0x59, 0xa6, 0x33, 0x71, 0x1b, 0x8f, 0xc6, 0xc1, 0xa4, 0x4f, 0x2b, 0x41, 0xf2, 0x47, 0x00,
	0xdd, 0xf9, 0xec, 0x82, 0x89, 0xd4, 0x10, 0x5f, 0x30, 0x99, 0xaa, 0x38, 0x18, 0xb7, 0x27, 0x21,
	0xb5, 0xc0, 0x94, 0xd5, 0x46, 0xe6, 0xcb, 0x6a, 0x11, 0x89, 0xa1, 0xc7, 0xd2, 0x54, 0x72, 0xa5,
	0xb0, 0xa8, 0x03, 0xea, 0x21, 0x26, 0x40, 0xbd, 0xcd, 0x04, 0x16, 0xb3, 0x4f, 0x2d, 0x30, 0x21,
	0xe6, 0xfc, 0x81, 0x4b, 0x76, 0x6b, 0x2b, 0x19, 0xd2, 0x12, 0x27, 0x1f, 0xa1, 0xef, 0x0b, 0x42,
	0x5e, 0x40, 0xf8, 0x81, 0x09, 0xc7, 0x22, 0x9a, 0xee, 0x97, 0x79, 0x36, 0x1c, 0xa9, 0xd5, 0x6d,
	0x5e, 0xde, 0xaa, 0x5f, 0x7e, 0x08, 0x5d, 0x6d, 0x33, 0x65, 0x5b, 0xcd, 0x21, 0x43, 0x4a, 0x72,
	0x96, 0xae, 0x3d, 0x29, 0x04, 0xc9, 0x5f, 0x81, 0xb9, 0xd9, 0x56, 0xf0, 0xff, 0xdd, 0x7c, 0x08,
	0xdd, 0x8f, 0x99, 0x10, 0x5c, 0xba, 0x8b, 0x1d, 0xaa, 0x85, 0xd7, 0xae, 0x87, 0x67, 0xce, 0xd8,
	0x7a, 0xb8, 0xf6, 0x76, 0x88, 0x9c, 0xc0, 0xd0, 0x7e, 0xbd, 0xa9, 0x27, 0x66, 0x4b, 0x9a, 0xbc,
	0x86, 0x9e, 0xeb, 0xeb, 0x86, 0x1a, 0xc5, 0xd0, 0x5b, 0x2d, 0x33, 0xa1, 0x1d, 0xab, 0x90, 0x7a,
	0x98, 0xfc, 0x1b, 0xc0, 0xb3, 0xf9, 0xec, 0x17, 0x76, 0xc7, 0xcf, 0x71, 0xa2, 0xc9, 0x77, 0x10,
	0xe2, 0x70, 0xe1, 0xf0, 0x46, 0x53, 0x52, 0x06, 0x69, 0x6c, 0xae, 0x8d, 0xe6, 0x62, 0x8f, 0x5a,
	0x13, 0x72, 0x06, 0xfd, 0xc5, 0x52, 0xe8, 0x4c, 0x14, 0x1c, 0xfd, 0x46, 0xd3, 0x2f, 0x6b, 0xe6,
	0x3f, 0x3b, 0xe5, 0xc5, 0x1e, 0x2d, 0x0d, 0xc9, 0xb7, 0xd0, 0x31, 0xcd, 0x89, 0x49, 0x88, 0xa6,
	0xcf, 0x6b, 0x07, 0xae, 0x8a, 0xcc, 0xb8, 0x47, 0x03, 0xc3, 0xe4, 0xbe, 0xe0, 0xd2, 0x56, 0x64,
	0x9b, 0xc9, 0x95, 0xd1, 0x18, 0x26, 0x68, 0x42, 0x86, 0xd0, 0xd2, 0x6b, 0x1c, 0xa8, 0x90, 0xb6,
	0xf4, 0x7a, 0xd6, 0x73, 0x3b, 0x21, 0x39, 0x87, 0x68, 0x83, 0x7a, 0xb5, 0x2b, 0x82, 0xcd, 0x5d,
	0x51, 0x1b, 0xed, 0xd6, 0xd6, 0x68, 0x27, 0x13, 0x18, 0xd6, 0xc3, 0x69, 0xda, 0x70, 0xc9, 0x31,
	0x40, 0x15, 0x47, 0xa3, 0xd5, 0x37, 0x9e, 0x12, 0xc6, 0xd0, 0x68, 0x76, 0x0f, 0x2f, 0xd0, 0xc0,
	0xda, 0xbe, 0xc9, 0x94, 0x9e, 0xad, 0xaf, 0x71, 0x63, 0x9e, 0x8b, 0x74, 0x5e, 0x2e, 0x9e, 0x6a,
	0xab, 0x06, 0xdb, 0x5b, 0xb5, 0x39, 0xa6, 0x6a, 0x07, 0xb5, 0x37, 0x76, 0x50, 0x72, 0xe2, 0x7b,
	0x81, 0xf2, 0xc5, 0x52, 0xa6, 0x8d, 0xd4, 0x5e, 0xc2, 0x68, 0x83, 0xda, 0xa5, 0x78, 0xbf, 0x6c,
	0x34, 0x3d, 0x83, 0x88, 0xf2, 0x55, 0xee, 0x4c, 0xc9, 0x31, 0x74, 0x8c, 0xc2, 0x35, 0xd7, 0x81,
	0x2f, 0xa9, 0x7f, 0x3b, 0x28, 0x6a, 0x93, 0x57, 0x70, 0xb0, 0xe5, 0x1f, 0x5b, 0xd8, 0xba, 0xb4,
	0xad, 0x3d, 0xa0, 0x1e, 0x26, 0xaf, 0x61, 0xb4, 0x71, 0x85, 0x49, 0x14, 0x39, 0x81, 0xd0, 0x68,
	0xfd, 0xa4, 0x7e, 0x7a, 0x8f, 0x55, 0x27, 0x7f, 0x07, 0xb0, 0x4f, 0xf9, 0x82, 0x67, 0x2b, 0xed,
	0x08, 0x3e, 0xf5, 0xf1, 0x22, 0xd0, 0x31, 0x9b, 0xc5, 0x0d, 0x2e, 0x7e, 0x57, 0xc9, 0x0d, 0x1b,
	0x17, 0x7c, 0x77, 0x7b, 0xc1, 0xd7, 0xca, 0xd5, 0xdb, 0x6e, 0xc1, 0x1b, 0x53, 0x98, 0x6b, 0xfb,
	0xe6, 0x51, 0x7e, 0xff, 0x39, 0x6d, 0x4c, 0x0e, 0xa0, 0xfd, 0x9e, 0xfb, 0x07, 0xd6, 0x7c, 0x26,
	0x3f, 0xc1, 0x68, 0x3e, 0xf3, 0x4d, 0x6d, 0x1d, 0x37, 0x85, 0xef, 0x0e, 0xb7, 0xaa, 0xc3, 0x3f,
	0x9a, 0x2e, 0xbe, 0xc2, 0x57, 0xe6, 0x69, 0x07, 0x7f, 0x30, 0x43, 0x82, 0xe5, 0x7d, 0xd2, 0xb9,
	0xe9, 0x9f, 0x01, 0x0c, 0xca, 0x5f, 0x12, 0x72, 0x0a, 0xa1, 0x9d, 0xe8, 0x1d, 0x0b, 0xea, 0xc8,
	0xbf, 0xbd, 0xbf, 0x09, 0x95, 0xdd, 0x8a, 0x9b, 0xc7, 0x64, 0x8f, 0x7c, 0x0f, 0xfd, 0x72, 0x7c,
	0x77, 0x2f, 0xa9, 0x5d, 0xa7, 0x5e, 0x41, 0x07, 0x47, 0xf9, 0xd3, 0x2d, 0xb5, 0xc3, 0xfa, 0x5d,
	0x17, 0xff, 0x86, 0xce, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x4b, 0x0f, 0xda, 0x3a, 0x09,
	0x00, 0x00,
}
