// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pbft.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Operation struct {
	Value *Block `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *Operation) Reset()                    { *m = Operation{} }
func (m *Operation) String() string            { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()               {}
func (*Operation) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *Operation) GetValue() *Block {
	if m != nil {
		return m.Value
	}
	return nil
}

type Checkpoint struct {
	Sequence uint32 `protobuf:"varint,1,opt,name=sequence" json:"sequence,omitempty"`
	Digest   []byte `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (m *Checkpoint) Reset()                    { *m = Checkpoint{} }
func (m *Checkpoint) String() string            { return proto.CompactTextString(m) }
func (*Checkpoint) ProtoMessage()               {}
func (*Checkpoint) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

func (m *Checkpoint) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Checkpoint) GetDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

type Entry struct {
	Sequence uint32 `protobuf:"varint,1,opt,name=sequence" json:"sequence,omitempty"`
	Digest   []byte `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
	View     uint32 `protobuf:"varint,3,opt,name=view" json:"view,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{2} }

func (m *Entry) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Entry) GetDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *Entry) GetView() uint32 {
	if m != nil {
		return m.View
	}
	return 0
}

type ViewChange struct {
	Viewchanger uint32 `protobuf:"varint,1,opt,name=viewchanger" json:"viewchanger,omitempty"`
	Digest      []byte `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (m *ViewChange) Reset()                    { *m = ViewChange{} }
func (m *ViewChange) String() string            { return proto.CompactTextString(m) }
func (*ViewChange) ProtoMessage()               {}
func (*ViewChange) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{3} }

func (m *ViewChange) GetViewchanger() uint32 {
	if m != nil {
		return m.Viewchanger
	}
	return 0
}

func (m *ViewChange) GetDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

type Summary struct {
	Sequence uint32 `protobuf:"varint,1,opt,name=sequence" json:"sequence,omitempty"`
	Digest   []byte `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (m *Summary) Reset()                    { *m = Summary{} }
func (m *Summary) String() string            { return proto.CompactTextString(m) }
func (*Summary) ProtoMessage()               {}
func (*Summary) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{4} }

func (m *Summary) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Summary) GetDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

type Result struct {
	Value *Block `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{5} }

func (m *Result) GetValue() *Block {
	if m != nil {
		return m.Value
	}
	return nil
}

type Request struct {
	// Types that are valid to be assigned to Value:
	//	*Request_Client
	//	*Request_Preprepare
	//	*Request_Prepare
	//	*Request_Commit
	//	*Request_Checkpoint
	//	*Request_Viewchange
	//	*Request_Ack
	//	*Request_Newview
	Value isRequest_Value `protobuf_oneof:"value"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{6} }

type isRequest_Value interface {
	isRequest_Value()
}

type Request_Client struct {
	Client *RequestClient `protobuf:"bytes,1,opt,name=client,oneof"`
}
type Request_Preprepare struct {
	Preprepare *RequestPrePrepare `protobuf:"bytes,2,opt,name=preprepare,oneof"`
}
type Request_Prepare struct {
	Prepare *RequestPrepare `protobuf:"bytes,3,opt,name=prepare,oneof"`
}
type Request_Commit struct {
	Commit *RequestCommit `protobuf:"bytes,4,opt,name=commit,oneof"`
}
type Request_Checkpoint struct {
	Checkpoint *RequestCheckpoint `protobuf:"bytes,5,opt,name=checkpoint,oneof"`
}
type Request_Viewchange struct {
	Viewchange *RequestViewChange `protobuf:"bytes,6,opt,name=viewchange,oneof"`
}
type Request_Ack struct {
	Ack *RequestAck `protobuf:"bytes,7,opt,name=ack,oneof"`
}
type Request_Newview struct {
	Newview *RequestNewView `protobuf:"bytes,8,opt,name=newview,oneof"`
}

func (*Request_Client) isRequest_Value()     {}
func (*Request_Preprepare) isRequest_Value() {}
func (*Request_Prepare) isRequest_Value()    {}
func (*Request_Commit) isRequest_Value()     {}
func (*Request_Checkpoint) isRequest_Value() {}
func (*Request_Viewchange) isRequest_Value() {}
func (*Request_Ack) isRequest_Value()        {}
func (*Request_Newview) isRequest_Value()    {}

func (m *Request) GetValue() isRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Request) GetClient() *RequestClient {
	if x, ok := m.GetValue().(*Request_Client); ok {
		return x.Client
	}
	return nil
}

func (m *Request) GetPreprepare() *RequestPrePrepare {
	if x, ok := m.GetValue().(*Request_Preprepare); ok {
		return x.Preprepare
	}
	return nil
}

func (m *Request) GetPrepare() *RequestPrepare {
	if x, ok := m.GetValue().(*Request_Prepare); ok {
		return x.Prepare
	}
	return nil
}

func (m *Request) GetCommit() *RequestCommit {
	if x, ok := m.GetValue().(*Request_Commit); ok {
		return x.Commit
	}
	return nil
}

func (m *Request) GetCheckpoint() *RequestCheckpoint {
	if x, ok := m.GetValue().(*Request_Checkpoint); ok {
		return x.Checkpoint
	}
	return nil
}

func (m *Request) GetViewchange() *RequestViewChange {
	if x, ok := m.GetValue().(*Request_Viewchange); ok {
		return x.Viewchange
	}
	return nil
}

func (m *Request) GetAck() *RequestAck {
	if x, ok := m.GetValue().(*Request_Ack); ok {
		return x.Ack
	}
	return nil
}

func (m *Request) GetNewview() *RequestNewView {
	if x, ok := m.GetValue().(*Request_Newview); ok {
		return x.Newview
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Request) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Request_OneofMarshaler, _Request_OneofUnmarshaler, _Request_OneofSizer, []interface{}{
		(*Request_Client)(nil),
		(*Request_Preprepare)(nil),
		(*Request_Prepare)(nil),
		(*Request_Commit)(nil),
		(*Request_Checkpoint)(nil),
		(*Request_Viewchange)(nil),
		(*Request_Ack)(nil),
		(*Request_Newview)(nil),
	}
}

func _Request_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Request)
	// value
	switch x := m.Value.(type) {
	case *Request_Client:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Client); err != nil {
			return err
		}
	case *Request_Preprepare:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Preprepare); err != nil {
			return err
		}
	case *Request_Prepare:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Prepare); err != nil {
			return err
		}
	case *Request_Commit:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Commit); err != nil {
			return err
		}
	case *Request_Checkpoint:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Checkpoint); err != nil {
			return err
		}
	case *Request_Viewchange:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Viewchange); err != nil {
			return err
		}
	case *Request_Ack:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Ack); err != nil {
			return err
		}
	case *Request_Newview:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Newview); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Request.Value has unexpected type %T", x)
	}
	return nil
}

func _Request_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Request)
	switch tag {
	case 1: // value.client
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestClient)
		err := b.DecodeMessage(msg)
		m.Value = &Request_Client{msg}
		return true, err
	case 2: // value.preprepare
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestPrePrepare)
		err := b.DecodeMessage(msg)
		m.Value = &Request_Preprepare{msg}
		return true, err
	case 3: // value.prepare
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestPrepare)
		err := b.DecodeMessage(msg)
		m.Value = &Request_Prepare{msg}
		return true, err
	case 4: // value.commit
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestCommit)
		err := b.DecodeMessage(msg)
		m.Value = &Request_Commit{msg}
		return true, err
	case 5: // value.checkpoint
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestCheckpoint)
		err := b.DecodeMessage(msg)
		m.Value = &Request_Checkpoint{msg}
		return true, err
	case 6: // value.viewchange
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestViewChange)
		err := b.DecodeMessage(msg)
		m.Value = &Request_Viewchange{msg}
		return true, err
	case 7: // value.ack
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestAck)
		err := b.DecodeMessage(msg)
		m.Value = &Request_Ack{msg}
		return true, err
	case 8: // value.newview
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestNewView)
		err := b.DecodeMessage(msg)
		m.Value = &Request_Newview{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Request_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Request)
	// value
	switch x := m.Value.(type) {
	case *Request_Client:
		s := proto.Size(x.Client)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Request_Preprepare:
		s := proto.Size(x.Preprepare)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Request_Prepare:
		s := proto.Size(x.Prepare)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Request_Commit:
		s := proto.Size(x.Commit)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Request_Checkpoint:
		s := proto.Size(x.Checkpoint)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Request_Viewchange:
		s := proto.Size(x.Viewchange)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Request_Ack:
		s := proto.Size(x.Ack)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Request_Newview:
		s := proto.Size(x.Newview)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type RequestClient struct {
	Op        *Operation `protobuf:"bytes,1,opt,name=op" json:"op,omitempty"`
	Timestamp string     `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Client    string     `protobuf:"bytes,3,opt,name=client" json:"client,omitempty"`
}

func (m *RequestClient) Reset()                    { *m = RequestClient{} }
func (m *RequestClient) String() string            { return proto.CompactTextString(m) }
func (*RequestClient) ProtoMessage()               {}
func (*RequestClient) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{7} }

func (m *RequestClient) GetOp() *Operation {
	if m != nil {
		return m.Op
	}
	return nil
}

func (m *RequestClient) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *RequestClient) GetClient() string {
	if m != nil {
		return m.Client
	}
	return ""
}

type RequestPrePrepare struct {
	View     uint32 `protobuf:"varint,1,opt,name=view" json:"view,omitempty"`
	Sequence uint32 `protobuf:"varint,2,opt,name=sequence" json:"sequence,omitempty"`
	Digest   []byte `protobuf:"bytes,3,opt,name=digest,proto3" json:"digest,omitempty"`
	Replica  uint32 `protobuf:"varint,4,opt,name=replica" json:"replica,omitempty"`
}

func (m *RequestPrePrepare) Reset()                    { *m = RequestPrePrepare{} }
func (m *RequestPrePrepare) String() string            { return proto.CompactTextString(m) }
func (*RequestPrePrepare) ProtoMessage()               {}
func (*RequestPrePrepare) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{8} }

func (m *RequestPrePrepare) GetView() uint32 {
	if m != nil {
		return m.View
	}
	return 0
}

func (m *RequestPrePrepare) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *RequestPrePrepare) GetDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *RequestPrePrepare) GetReplica() uint32 {
	if m != nil {
		return m.Replica
	}
	return 0
}

type RequestPrepare struct {
	View     uint32 `protobuf:"varint,1,opt,name=view" json:"view,omitempty"`
	Sequence uint32 `protobuf:"varint,2,opt,name=sequence" json:"sequence,omitempty"`
	Digest   []byte `protobuf:"bytes,3,opt,name=digest,proto3" json:"digest,omitempty"`
	Replica  uint32 `protobuf:"varint,4,opt,name=replica" json:"replica,omitempty"`
}

func (m *RequestPrepare) Reset()                    { *m = RequestPrepare{} }
func (m *RequestPrepare) String() string            { return proto.CompactTextString(m) }
func (*RequestPrepare) ProtoMessage()               {}
func (*RequestPrepare) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{9} }

func (m *RequestPrepare) GetView() uint32 {
	if m != nil {
		return m.View
	}
	return 0
}

func (m *RequestPrepare) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *RequestPrepare) GetDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *RequestPrepare) GetReplica() uint32 {
	if m != nil {
		return m.Replica
	}
	return 0
}

type RequestCommit struct {
	View     uint32 `protobuf:"varint,1,opt,name=view" json:"view,omitempty"`
	Sequence uint32 `protobuf:"varint,2,opt,name=sequence" json:"sequence,omitempty"`
	Replica  uint32 `protobuf:"varint,3,opt,name=replica" json:"replica,omitempty"`
}

func (m *RequestCommit) Reset()                    { *m = RequestCommit{} }
func (m *RequestCommit) String() string            { return proto.CompactTextString(m) }
func (*RequestCommit) ProtoMessage()               {}
func (*RequestCommit) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{10} }

func (m *RequestCommit) GetView() uint32 {
	if m != nil {
		return m.View
	}
	return 0
}

func (m *RequestCommit) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *RequestCommit) GetReplica() uint32 {
	if m != nil {
		return m.Replica
	}
	return 0
}

type RequestCheckpoint struct {
	Sequence uint32 `protobuf:"varint,1,opt,name=sequence" json:"sequence,omitempty"`
	Digest   []byte `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
	Replica  uint32 `protobuf:"varint,3,opt,name=replica" json:"replica,omitempty"`
}

func (m *RequestCheckpoint) Reset()                    { *m = RequestCheckpoint{} }
func (m *RequestCheckpoint) String() string            { return proto.CompactTextString(m) }
func (*RequestCheckpoint) ProtoMessage()               {}
func (*RequestCheckpoint) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{11} }

func (m *RequestCheckpoint) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *RequestCheckpoint) GetDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *RequestCheckpoint) GetReplica() uint32 {
	if m != nil {
		return m.Replica
	}
	return 0
}

type RequestViewChange struct {
	View        uint32        `protobuf:"varint,1,opt,name=view" json:"view,omitempty"`
	Sequence    uint32        `protobuf:"varint,2,opt,name=sequence" json:"sequence,omitempty"`
	Checkpoints []*Checkpoint `protobuf:"bytes,3,rep,name=checkpoints" json:"checkpoints,omitempty"`
	Preps       []*Entry      `protobuf:"bytes,4,rep,name=preps" json:"preps,omitempty"`
	Prepreps    []*Entry      `protobuf:"bytes,5,rep,name=prepreps" json:"prepreps,omitempty"`
	Replica     uint32        `protobuf:"varint,6,opt,name=replica" json:"replica,omitempty"`
}

func (m *RequestViewChange) Reset()                    { *m = RequestViewChange{} }
func (m *RequestViewChange) String() string            { return proto.CompactTextString(m) }
func (*RequestViewChange) ProtoMessage()               {}
func (*RequestViewChange) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{12} }

func (m *RequestViewChange) GetView() uint32 {
	if m != nil {
		return m.View
	}
	return 0
}

func (m *RequestViewChange) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *RequestViewChange) GetCheckpoints() []*Checkpoint {
	if m != nil {
		return m.Checkpoints
	}
	return nil
}

func (m *RequestViewChange) GetPreps() []*Entry {
	if m != nil {
		return m.Preps
	}
	return nil
}

func (m *RequestViewChange) GetPrepreps() []*Entry {
	if m != nil {
		return m.Prepreps
	}
	return nil
}

func (m *RequestViewChange) GetReplica() uint32 {
	if m != nil {
		return m.Replica
	}
	return 0
}

type RequestAck struct {
	View        uint32 `protobuf:"varint,1,opt,name=view" json:"view,omitempty"`
	Replica     uint32 `protobuf:"varint,2,opt,name=replica" json:"replica,omitempty"`
	Viewchanger uint32 `protobuf:"varint,3,opt,name=viewchanger" json:"viewchanger,omitempty"`
	Digest      []byte `protobuf:"bytes,4,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (m *RequestAck) Reset()                    { *m = RequestAck{} }
func (m *RequestAck) String() string            { return proto.CompactTextString(m) }
func (*RequestAck) ProtoMessage()               {}
func (*RequestAck) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{13} }

func (m *RequestAck) GetView() uint32 {
	if m != nil {
		return m.View
	}
	return 0
}

func (m *RequestAck) GetReplica() uint32 {
	if m != nil {
		return m.Replica
	}
	return 0
}

func (m *RequestAck) GetViewchanger() uint32 {
	if m != nil {
		return m.Viewchanger
	}
	return 0
}

func (m *RequestAck) GetDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

type RequestNewView struct {
	View        uint32        `protobuf:"varint,1,opt,name=view" json:"view,omitempty"`
	Viewchanges []*ViewChange `protobuf:"bytes,2,rep,name=viewchanges" json:"viewchanges,omitempty"`
	Summaries   []*Summary    `protobuf:"bytes,4,rep,name=summaries" json:"summaries,omitempty"`
	Replica     uint32        `protobuf:"varint,5,opt,name=replica" json:"replica,omitempty"`
}

func (m *RequestNewView) Reset()                    { *m = RequestNewView{} }
func (m *RequestNewView) String() string            { return proto.CompactTextString(m) }
func (*RequestNewView) ProtoMessage()               {}
func (*RequestNewView) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{14} }

func (m *RequestNewView) GetView() uint32 {
	if m != nil {
		return m.View
	}
	return 0
}

func (m *RequestNewView) GetViewchanges() []*ViewChange {
	if m != nil {
		return m.Viewchanges
	}
	return nil
}

func (m *RequestNewView) GetSummaries() []*Summary {
	if m != nil {
		return m.Summaries
	}
	return nil
}

func (m *RequestNewView) GetReplica() uint32 {
	if m != nil {
		return m.Replica
	}
	return 0
}

type ClientReply struct {
	View      uint32  `protobuf:"varint,1,opt,name=view" json:"view,omitempty"`
	Timestamp string  `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Client    string  `protobuf:"bytes,3,opt,name=client" json:"client,omitempty"`
	Replica   uint32  `protobuf:"varint,4,opt,name=replica" json:"replica,omitempty"`
	Result    *Result `protobuf:"bytes,5,opt,name=result" json:"result,omitempty"`
}

func (m *ClientReply) Reset()                    { *m = ClientReply{} }
func (m *ClientReply) String() string            { return proto.CompactTextString(m) }
func (*ClientReply) ProtoMessage()               {}
func (*ClientReply) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{15} }

func (m *ClientReply) GetView() uint32 {
	if m != nil {
		return m.View
	}
	return 0
}

func (m *ClientReply) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *ClientReply) GetClient() string {
	if m != nil {
		return m.Client
	}
	return ""
}

func (m *ClientReply) GetReplica() uint32 {
	if m != nil {
		return m.Replica
	}
	return 0
}

func (m *ClientReply) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*Operation)(nil), "types.Operation")
	proto.RegisterType((*Checkpoint)(nil), "types.Checkpoint")
	proto.RegisterType((*Entry)(nil), "types.Entry")
	proto.RegisterType((*ViewChange)(nil), "types.ViewChange")
	proto.RegisterType((*Summary)(nil), "types.Summary")
	proto.RegisterType((*Result)(nil), "types.Result")
	proto.RegisterType((*Request)(nil), "types.Request")
	proto.RegisterType((*RequestClient)(nil), "types.RequestClient")
	proto.RegisterType((*RequestPrePrepare)(nil), "types.RequestPrePrepare")
	proto.RegisterType((*RequestPrepare)(nil), "types.RequestPrepare")
	proto.RegisterType((*RequestCommit)(nil), "types.RequestCommit")
	proto.RegisterType((*RequestCheckpoint)(nil), "types.RequestCheckpoint")
	proto.RegisterType((*RequestViewChange)(nil), "types.RequestViewChange")
	proto.RegisterType((*RequestAck)(nil), "types.RequestAck")
	proto.RegisterType((*RequestNewView)(nil), "types.RequestNewView")
	proto.RegisterType((*ClientReply)(nil), "types.ClientReply")
}

func init() { proto.RegisterFile("pbft.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 651 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x5f, 0x6b, 0xdb, 0x30,
	0x10, 0xaf, 0xed, 0xd8, 0x4e, 0x2e, 0x4d, 0x69, 0xc5, 0x36, 0x44, 0xd9, 0x43, 0x30, 0x14, 0xf2,
	0x50, 0x32, 0xd6, 0xbc, 0x0d, 0x06, 0x5b, 0xc3, 0x46, 0x9e, 0xd6, 0xa2, 0xc1, 0x60, 0x8f, 0x8e,
	0xa7, 0x25, 0x26, 0xfe, 0x37, 0x4b, 0x69, 0xc8, 0x37, 0xd9, 0xeb, 0x3e, 0xc3, 0xbe, 0xd0, 0x3e,
	0xca, 0x90, 0xac, 0x58, 0xd6, 0xea, 0x94, 0x35, 0x83, 0x81, 0x1f, 0x2c, 0xdd, 0xfd, 0x74, 0xa7,
	0xbb, 0xdf, 0xfd, 0x04, 0x50, 0xcc, 0xbf, 0xf2, 0x71, 0x51, 0xe6, 0x3c, 0x47, 0x2e, 0xdf, 0x16,
	0x94, 0x9d, 0x9f, 0xce, 0x93, 0x3c, 0x5a, 0x45, 0xcb, 0x30, 0xce, 0x2a, 0x43, 0xf0, 0x02, 0x7a,
	0x37, 0x05, 0x2d, 0x43, 0x1e, 0xe7, 0x19, 0x0a, 0xc0, 0xbd, 0x0b, 0x93, 0x35, 0xc5, 0xd6, 0xd0,
	0x1a, 0xf5, 0xaf, 0x8e, 0xc7, 0x12, 0x35, 0xbe, 0x16, 0x20, 0x52, 0x99, 0x82, 0x37, 0x00, 0xd3,
	0x25, 0x8d, 0x56, 0x45, 0x1e, 0x67, 0x1c, 0x9d, 0x43, 0x97, 0xd1, 0x6f, 0x6b, 0x9a, 0x45, 0x15,
	0x68, 0x40, 0xea, 0x35, 0x7a, 0x06, 0xde, 0x97, 0x78, 0x41, 0x19, 0xc7, 0xf6, 0xd0, 0x1a, 0x1d,
	0x13, 0xb5, 0x0a, 0x6e, 0xc0, 0x7d, 0x97, 0xf1, 0x72, 0x7b, 0x08, 0x18, 0x21, 0xe8, 0xdc, 0xc5,
	0x74, 0x83, 0x1d, 0xe9, 0x2f, 0xff, 0x83, 0xf7, 0x00, 0x9f, 0x62, 0xba, 0x99, 0x2e, 0xc3, 0x6c,
	0x41, 0xd1, 0x10, 0xfa, 0x62, 0x37, 0x92, 0xab, 0x52, 0x1d, 0xdc, 0xdc, 0xda, 0x9b, 0xd8, 0x6b,
	0xf0, 0x3f, 0xae, 0xd3, 0x34, 0x3c, 0x2c, 0xb5, 0xe0, 0x12, 0x3c, 0x42, 0xd9, 0x3a, 0xe1, 0x7f,
	0x55, 0xc7, 0x9f, 0x0e, 0xf8, 0x44, 0x1c, 0xc9, 0x38, 0x1a, 0x83, 0x17, 0x25, 0x31, 0xcd, 0xb8,
	0x02, 0x3c, 0x51, 0x00, 0x65, 0x9f, 0x4a, 0xdb, 0xec, 0x88, 0x28, 0x2f, 0xf4, 0x0a, 0xa0, 0x28,
	0xa9, 0xf8, 0xc2, 0x92, 0xca, 0x2c, 0xfa, 0x57, 0xd8, 0xc4, 0xdc, 0x96, 0xf4, 0xb6, 0xb2, 0xcf,
	0x8e, 0x48, 0xc3, 0x1b, 0xbd, 0x04, 0x7f, 0x07, 0x74, 0x24, 0xf0, 0xe9, 0x3d, 0xa0, 0x42, 0xed,
	0xfc, 0x64, 0x7a, 0x79, 0x9a, 0xc6, 0x1c, 0x77, 0x5a, 0xd3, 0x93, 0x36, 0x99, 0x9e, 0xfc, 0x13,
	0xe9, 0x45, 0x35, 0x45, 0xb0, 0xdb, 0x96, 0x9e, 0xa6, 0x90, 0x48, 0x4f, 0x7b, 0x0b, 0xac, 0x6e,
	0x15, 0xf6, 0xda, 0xb0, 0xba, 0xd7, 0x02, 0xab, 0xbd, 0xd1, 0x05, 0x38, 0x61, 0xb4, 0xc2, 0xbe,
	0x04, 0x9d, 0x99, 0xa0, 0xb7, 0xd1, 0x6a, 0x76, 0x44, 0x84, 0x5d, 0x54, 0x20, 0xa3, 0x1b, 0xc9,
	0xa2, 0x6e, 0x5b, 0x05, 0x3e, 0xd0, 0x8d, 0x08, 0x21, 0x2a, 0xa0, 0xfc, 0xae, 0x7d, 0xd5, 0xd0,
	0x60, 0x01, 0x03, 0xa3, 0x29, 0x68, 0x08, 0x76, 0x5e, 0xa8, 0xb6, 0x9d, 0xaa, 0x73, 0xea, 0x81,
	0x22, 0x76, 0x5e, 0xa0, 0xe7, 0xd0, 0xe3, 0x71, 0x4a, 0x19, 0x0f, 0xd3, 0x42, 0xf6, 0xaa, 0x47,
	0xf4, 0x86, 0x20, 0x93, 0x6a, 0xbd, 0x23, 0x4d, 0x6a, 0x15, 0xac, 0xe1, 0xec, 0x5e, 0x27, 0x6b,
	0xf2, 0x5b, 0x9a, 0xfc, 0x06, 0x53, 0xed, 0xbd, 0x4c, 0x75, 0x8c, 0x21, 0xc2, 0xe0, 0x97, 0xb4,
	0x48, 0xe2, 0x28, 0x94, 0x1d, 0x1d, 0x90, 0xdd, 0x32, 0x28, 0xe1, 0xc4, 0xe4, 0xc1, 0x7f, 0x88,
	0xf9, 0x59, 0xd7, 0xb4, 0xe2, 0xcf, 0x63, 0x43, 0x36, 0x8e, 0x76, 0xcc, 0xa3, 0xc3, 0xba, 0x8a,
	0xff, 0xa6, 0x59, 0x0f, 0x84, 0xf8, 0x65, 0xd5, 0x31, 0x1a, 0x22, 0xf4, 0xd8, 0x2b, 0x4c, 0xa0,
	0xaf, 0x87, 0x80, 0x61, 0x67, 0xe8, 0x34, 0x28, 0xac, 0x73, 0x27, 0x4d, 0x2f, 0x21, 0x33, 0x62,
	0x44, 0x19, 0xee, 0x48, 0xf7, 0x9d, 0xcc, 0x48, 0x71, 0x25, 0x95, 0x09, 0x8d, 0xa0, 0xab, 0x86,
	0x9f, 0x61, 0xb7, 0xc5, 0xad, 0xb6, 0x36, 0xaf, 0xe8, 0x99, 0x57, 0xe4, 0x00, 0x7a, 0x8a, 0x5a,
	0xaf, 0xd6, 0xc0, 0xda, 0x06, 0xf6, 0x4f, 0x35, 0x76, 0x1e, 0x52, 0xe3, 0x8e, 0x21, 0xa7, 0x3f,
	0xac, 0x9a, 0x8b, 0x6a, 0x22, 0x5b, 0x43, 0x4f, 0x9a, 0x01, 0x18, 0xb6, 0x8d, 0xca, 0xe9, 0x8e,
	0x34, 0x63, 0x32, 0x74, 0x09, 0x3d, 0x26, 0x95, 0x3e, 0xa6, 0xbb, 0xea, 0x9d, 0x28, 0x88, 0x7a,
	0x01, 0x88, 0x76, 0x68, 0xde, 0xce, 0x35, 0x2b, 0xf3, 0xdd, 0x82, 0x7e, 0x25, 0x04, 0x84, 0x16,
	0xc9, 0xb6, 0x35, 0xc1, 0x83, 0xe6, 0x7f, 0xff, 0xb8, 0xa0, 0x0b, 0xf0, 0x4a, 0xf9, 0xcc, 0x28,
	0x65, 0x1d, 0xd4, 0xea, 0x25, 0x36, 0x89, 0x32, 0xce, 0x3d, 0xf9, 0xbe, 0x4f, 0x7e, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x05, 0x1f, 0x3f, 0xa6, 0x06, 0x08, 0x00, 0x00,
}
