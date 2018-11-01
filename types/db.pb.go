// Code generated by protoc-gen-go. DO NOT EDIT.
// source: db.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// merkle avl tree
type LeafNode struct {
	Key    []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value  []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Height int32  `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	Size   int32  `protobuf:"varint,4,opt,name=size" json:"size,omitempty"`
}

func (m *LeafNode) Reset()                    { *m = LeafNode{} }
func (m *LeafNode) String() string            { return proto.CompactTextString(m) }
func (*LeafNode) ProtoMessage()               {}
func (*LeafNode) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *LeafNode) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *LeafNode) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *LeafNode) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *LeafNode) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type InnerNode struct {
	LeftHash  []byte `protobuf:"bytes,1,opt,name=leftHash,proto3" json:"leftHash,omitempty"`
	RightHash []byte `protobuf:"bytes,2,opt,name=rightHash,proto3" json:"rightHash,omitempty"`
	Height    int32  `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	Size      int32  `protobuf:"varint,4,opt,name=size" json:"size,omitempty"`
}

func (m *InnerNode) Reset()                    { *m = InnerNode{} }
func (m *InnerNode) String() string            { return proto.CompactTextString(m) }
func (*InnerNode) ProtoMessage()               {}
func (*InnerNode) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *InnerNode) GetLeftHash() []byte {
	if m != nil {
		return m.LeftHash
	}
	return nil
}

func (m *InnerNode) GetRightHash() []byte {
	if m != nil {
		return m.RightHash
	}
	return nil
}

func (m *InnerNode) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *InnerNode) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type MAVLProof struct {
	LeafHash   []byte       `protobuf:"bytes,1,opt,name=leafHash,proto3" json:"leafHash,omitempty"`
	InnerNodes []*InnerNode `protobuf:"bytes,2,rep,name=innerNodes" json:"innerNodes,omitempty"`
	RootHash   []byte       `protobuf:"bytes,3,opt,name=rootHash,proto3" json:"rootHash,omitempty"`
}

func (m *MAVLProof) Reset()                    { *m = MAVLProof{} }
func (m *MAVLProof) String() string            { return proto.CompactTextString(m) }
func (*MAVLProof) ProtoMessage()               {}
func (*MAVLProof) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *MAVLProof) GetLeafHash() []byte {
	if m != nil {
		return m.LeafHash
	}
	return nil
}

func (m *MAVLProof) GetInnerNodes() []*InnerNode {
	if m != nil {
		return m.InnerNodes
	}
	return nil
}

func (m *MAVLProof) GetRootHash() []byte {
	if m != nil {
		return m.RootHash
	}
	return nil
}

type StoreNode struct {
	Key        []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value      []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	LeftHash   []byte `protobuf:"bytes,3,opt,name=leftHash,proto3" json:"leftHash,omitempty"`
	RightHash  []byte `protobuf:"bytes,4,opt,name=rightHash,proto3" json:"rightHash,omitempty"`
	Height     int32  `protobuf:"varint,5,opt,name=height" json:"height,omitempty"`
	Size       int32  `protobuf:"varint,6,opt,name=size" json:"size,omitempty"`
	ParentHash []byte `protobuf:"bytes,7,opt,name=parentHash,proto3" json:"parentHash,omitempty"`
}

func (m *StoreNode) Reset()                    { *m = StoreNode{} }
func (m *StoreNode) String() string            { return proto.CompactTextString(m) }
func (*StoreNode) ProtoMessage()               {}
func (*StoreNode) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *StoreNode) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *StoreNode) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *StoreNode) GetLeftHash() []byte {
	if m != nil {
		return m.LeftHash
	}
	return nil
}

func (m *StoreNode) GetRightHash() []byte {
	if m != nil {
		return m.RightHash
	}
	return nil
}

func (m *StoreNode) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *StoreNode) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *StoreNode) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

type LocalDBSet struct {
	KV []*KeyValue `protobuf:"bytes,2,rep,name=KV" json:"KV,omitempty"`
}

func (m *LocalDBSet) Reset()                    { *m = LocalDBSet{} }
func (m *LocalDBSet) String() string            { return proto.CompactTextString(m) }
func (*LocalDBSet) ProtoMessage()               {}
func (*LocalDBSet) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *LocalDBSet) GetKV() []*KeyValue {
	if m != nil {
		return m.KV
	}
	return nil
}

type LocalDBList struct {
	Prefix    []byte `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Key       []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Direction int32  `protobuf:"varint,3,opt,name=direction" json:"direction,omitempty"`
	Count     int32  `protobuf:"varint,4,opt,name=count" json:"count,omitempty"`
}

func (m *LocalDBList) Reset()                    { *m = LocalDBList{} }
func (m *LocalDBList) String() string            { return proto.CompactTextString(m) }
func (*LocalDBList) ProtoMessage()               {}
func (*LocalDBList) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *LocalDBList) GetPrefix() []byte {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *LocalDBList) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *LocalDBList) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

func (m *LocalDBList) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type LocalDBGet struct {
	Keys [][]byte `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (m *LocalDBGet) Reset()                    { *m = LocalDBGet{} }
func (m *LocalDBGet) String() string            { return proto.CompactTextString(m) }
func (*LocalDBGet) ProtoMessage()               {}
func (*LocalDBGet) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *LocalDBGet) GetKeys() [][]byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

type LocalReplyValue struct {
	Values [][]byte `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (m *LocalReplyValue) Reset()                    { *m = LocalReplyValue{} }
func (m *LocalReplyValue) String() string            { return proto.CompactTextString(m) }
func (*LocalReplyValue) ProtoMessage()               {}
func (*LocalReplyValue) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

func (m *LocalReplyValue) GetValues() [][]byte {
	if m != nil {
		return m.Values
	}
	return nil
}

type StoreSet struct {
	StateHash []byte      `protobuf:"bytes,1,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	KV        []*KeyValue `protobuf:"bytes,2,rep,name=KV" json:"KV,omitempty"`
	Height    int64       `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
}

func (m *StoreSet) Reset()                    { *m = StoreSet{} }
func (m *StoreSet) String() string            { return proto.CompactTextString(m) }
func (*StoreSet) ProtoMessage()               {}
func (*StoreSet) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

func (m *StoreSet) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *StoreSet) GetKV() []*KeyValue {
	if m != nil {
		return m.KV
	}
	return nil
}

func (m *StoreSet) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type StoreDel struct {
	StateHash []byte `protobuf:"bytes,1,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	Height    int64  `protobuf:"varint,2,opt,name=height" json:"height,omitempty"`
}

func (m *StoreDel) Reset()                    { *m = StoreDel{} }
func (m *StoreDel) String() string            { return proto.CompactTextString(m) }
func (*StoreDel) ProtoMessage()               {}
func (*StoreDel) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

func (m *StoreDel) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *StoreDel) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type StoreSetWithSync struct {
	Storeset *StoreSet `protobuf:"bytes,1,opt,name=storeset" json:"storeset,omitempty"`
	Sync     bool      `protobuf:"varint,2,opt,name=sync" json:"sync,omitempty"`
}

func (m *StoreSetWithSync) Reset()                    { *m = StoreSetWithSync{} }
func (m *StoreSetWithSync) String() string            { return proto.CompactTextString(m) }
func (*StoreSetWithSync) ProtoMessage()               {}
func (*StoreSetWithSync) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{10} }

func (m *StoreSetWithSync) GetStoreset() *StoreSet {
	if m != nil {
		return m.Storeset
	}
	return nil
}

func (m *StoreSetWithSync) GetSync() bool {
	if m != nil {
		return m.Sync
	}
	return false
}

type StoreGet struct {
	StateHash []byte   `protobuf:"bytes,1,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	Keys      [][]byte `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (m *StoreGet) Reset()                    { *m = StoreGet{} }
func (m *StoreGet) String() string            { return proto.CompactTextString(m) }
func (*StoreGet) ProtoMessage()               {}
func (*StoreGet) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{11} }

func (m *StoreGet) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *StoreGet) GetKeys() [][]byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

type StoreReplyValue struct {
	Values [][]byte `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (m *StoreReplyValue) Reset()                    { *m = StoreReplyValue{} }
func (m *StoreReplyValue) String() string            { return proto.CompactTextString(m) }
func (*StoreReplyValue) ProtoMessage()               {}
func (*StoreReplyValue) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{12} }

func (m *StoreReplyValue) GetValues() [][]byte {
	if m != nil {
		return m.Values
	}
	return nil
}

type PruneData struct {
	// 对应keyHash下的区块高度
	Height int64 `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
	// hash+prefix的长度
	Lenth int32 `protobuf:"varint,2,opt,name=lenth" json:"lenth,omitempty"`
}

func (m *PruneData) Reset()                    { *m = PruneData{} }
func (m *PruneData) String() string            { return proto.CompactTextString(m) }
func (*PruneData) ProtoMessage()               {}
func (*PruneData) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{13} }

func (m *PruneData) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *PruneData) GetLenth() int32 {
	if m != nil {
		return m.Lenth
	}
	return 0
}

type DeleteNodeMap struct {
	MpNode map[string]bool `protobuf:"bytes,1,rep,name=mpNode" json:"mpNode,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *DeleteNodeMap) Reset()                    { *m = DeleteNodeMap{} }
func (m *DeleteNodeMap) String() string            { return proto.CompactTextString(m) }
func (*DeleteNodeMap) ProtoMessage()               {}
func (*DeleteNodeMap) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{14} }

func (m *DeleteNodeMap) GetMpNode() map[string]bool {
	if m != nil {
		return m.MpNode
	}
	return nil
}

// 用于存储db Pool数据的Value
type StoreValuePool struct {
	Values [][]byte `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (m *StoreValuePool) Reset()                    { *m = StoreValuePool{} }
func (m *StoreValuePool) String() string            { return proto.CompactTextString(m) }
func (*StoreValuePool) ProtoMessage()               {}
func (*StoreValuePool) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{15} }

func (m *StoreValuePool) GetValues() [][]byte {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*LeafNode)(nil), "types.LeafNode")
	proto.RegisterType((*InnerNode)(nil), "types.InnerNode")
	proto.RegisterType((*MAVLProof)(nil), "types.MAVLProof")
	proto.RegisterType((*StoreNode)(nil), "types.StoreNode")
	proto.RegisterType((*LocalDBSet)(nil), "types.LocalDBSet")
	proto.RegisterType((*LocalDBList)(nil), "types.LocalDBList")
	proto.RegisterType((*LocalDBGet)(nil), "types.LocalDBGet")
	proto.RegisterType((*LocalReplyValue)(nil), "types.LocalReplyValue")
	proto.RegisterType((*StoreSet)(nil), "types.StoreSet")
	proto.RegisterType((*StoreDel)(nil), "types.StoreDel")
	proto.RegisterType((*StoreSetWithSync)(nil), "types.StoreSetWithSync")
	proto.RegisterType((*StoreGet)(nil), "types.StoreGet")
	proto.RegisterType((*StoreReplyValue)(nil), "types.StoreReplyValue")
	proto.RegisterType((*PruneData)(nil), "types.PruneData")
	proto.RegisterType((*DeleteNodeMap)(nil), "types.DeleteNodeMap")
	proto.RegisterType((*StoreValuePool)(nil), "types.StoreValuePool")
}

func init() { proto.RegisterFile("db.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 595 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0x95, 0xed, 0x26, 0x9f, 0x3d, 0xed, 0x47, 0xab, 0x15, 0x42, 0x56, 0x55, 0x41, 0xb4, 0xe2,
	0x10, 0x84, 0x30, 0x88, 0x5c, 0x5a, 0xc4, 0x01, 0xaa, 0xa0, 0x82, 0x9a, 0xa2, 0xca, 0x91, 0x82,
	0xc4, 0x01, 0x69, 0xeb, 0x4e, 0x1a, 0xab, 0xee, 0xae, 0xb1, 0x37, 0x08, 0x73, 0xe6, 0x4f, 0xf1,
	0xef, 0x90, 0x67, 0xd7, 0xb1, 0x8b, 0x5a, 0x42, 0x4f, 0xd9, 0x37, 0x99, 0xbc, 0x79, 0xef, 0xcd,
	0x6e, 0xc0, 0x3f, 0x3f, 0x8b, 0xf2, 0x42, 0x69, 0xc5, 0x7a, 0xba, 0xca, 0xb1, 0xdc, 0xdd, 0x4a,
	0xd4, 0xd5, 0x95, 0x92, 0xa6, 0xc8, 0xbf, 0x80, 0x3f, 0x41, 0x31, 0xff, 0xa8, 0xce, 0x91, 0xed,
	0x80, 0x77, 0x89, 0x55, 0xe8, 0x0c, 0x9c, 0xe1, 0x56, 0x5c, 0x1f, 0xd9, 0x7d, 0xe8, 0x7d, 0x13,
	0xd9, 0x12, 0x43, 0x97, 0x6a, 0x06, 0xb0, 0x07, 0xd0, 0x5f, 0x60, 0x7a, 0xb1, 0xd0, 0xa1, 0x37,
	0x70, 0x86, 0xbd, 0xd8, 0x22, 0xc6, 0x60, 0xa3, 0x4c, 0x7f, 0x60, 0xb8, 0x41, 0x55, 0x3a, 0xf3,
	0xaf, 0x10, 0x7c, 0x90, 0x12, 0x0b, 0x1a, 0xb0, 0x0b, 0x7e, 0x86, 0x73, 0xfd, 0x5e, 0x94, 0x0b,
	0x3b, 0x65, 0x85, 0xd9, 0x1e, 0x04, 0x45, 0xcd, 0x42, 0x5f, 0x9a, 0x71, 0x6d, 0xe1, 0x4e, 0x23,
	0x97, 0x10, 0x9c, 0xbc, 0x9d, 0x4d, 0x4e, 0x0b, 0xa5, 0xe6, 0x66, 0xa4, 0x98, 0x5f, 0x1f, 0x69,
	0x30, 0x7b, 0x01, 0x90, 0x36, 0xda, 0xca, 0xd0, 0x1d, 0x78, 0xc3, 0xcd, 0x97, 0x3b, 0x11, 0xa5,
	0x14, 0xad, 0x44, 0xc7, 0x9d, 0x9e, 0x9a, 0xad, 0x50, 0xca, 0x68, 0xf4, 0x0c, 0x5b, 0x83, 0xf9,
	0x2f, 0x07, 0x82, 0xa9, 0x56, 0x05, 0xde, 0x29, 0xcb, 0x6e, 0x24, 0xde, 0xdf, 0x22, 0xd9, 0xb8,
	0x3d, 0x92, 0xde, 0x8d, 0x91, 0xf4, 0xdb, 0x48, 0xd8, 0x43, 0x80, 0x5c, 0x14, 0x28, 0x0d, 0xd5,
	0x7f, 0x44, 0xd5, 0xa9, 0xf0, 0x67, 0x00, 0x13, 0x95, 0x88, 0x6c, 0x7c, 0x38, 0x45, 0xcd, 0x1e,
	0x81, 0x7b, 0x3c, 0xb3, 0x79, 0x6c, 0xdb, 0x3c, 0x8e, 0xb1, 0x9a, 0xd5, 0x82, 0x63, 0xf7, 0x78,
	0xc6, 0x2f, 0x61, 0xd3, 0xb6, 0x4f, 0xd2, 0x52, 0xd7, 0x4a, 0xf2, 0x02, 0xe7, 0xe9, 0x77, 0x6b,
	0xd7, 0xa2, 0x26, 0x03, 0xb7, 0xcd, 0x60, 0x0f, 0x82, 0xf3, 0xb4, 0xc0, 0x44, 0xa7, 0x4a, 0xda,
	0x4d, 0xb6, 0x85, 0x3a, 0xa1, 0x44, 0x2d, 0xa5, 0xb6, 0xdb, 0x34, 0x80, 0x0f, 0x56, 0xda, 0x8e,
	0x90, 0xdc, 0x5d, 0x62, 0x65, 0xb6, 0xb5, 0x15, 0xd3, 0x99, 0x3f, 0x81, 0x6d, 0xea, 0x88, 0x31,
	0xcf, 0x8c, 0xca, 0x5a, 0x12, 0xe5, 0xdb, 0x34, 0x5a, 0xc4, 0x05, 0xf8, 0xb4, 0xa3, 0xda, 0xe6,
	0x1e, 0x04, 0xa5, 0x16, 0x1a, 0x3b, 0x77, 0xa3, 0x2d, 0xac, 0x0d, 0xe1, 0x8f, 0x2b, 0xe9, 0x35,
	0xf9, 0xf3, 0x37, 0x76, 0xc4, 0x18, 0xb3, 0x35, 0x23, 0x5a, 0x06, 0xf7, 0x1a, 0xc3, 0x14, 0x76,
	0x1a, 0x91, 0x9f, 0x52, 0xbd, 0x98, 0x56, 0x32, 0x61, 0x4f, 0xc1, 0x2f, 0xeb, 0x5a, 0x89, 0x9a,
	0x88, 0x5a, 0x51, 0x4d, 0x6b, 0xbc, 0x6a, 0xa0, 0x2b, 0x50, 0xc9, 0x84, 0x68, 0xfd, 0x98, 0xce,
	0xfc, 0xb5, 0x95, 0x75, 0xb4, 0xd6, 0xf9, 0x2d, 0x11, 0xd3, 0xaf, 0xff, 0x21, 0xe2, 0x03, 0x08,
	0x4e, 0x8b, 0xa5, 0xc4, 0xb1, 0xd0, 0xa2, 0x63, 0xd1, 0xe9, 0x5a, 0xac, 0x57, 0x9d, 0xa1, 0xd4,
	0xe6, 0xa5, 0xf7, 0x62, 0x03, 0xf8, 0x4f, 0x07, 0xfe, 0x1f, 0x63, 0x86, 0x9a, 0xde, 0xd0, 0x89,
	0xc8, 0xd9, 0x3e, 0xf4, 0xaf, 0xf2, 0x1a, 0x84, 0x0e, 0x6d, 0x62, 0x60, 0x4d, 0x5f, 0xeb, 0x8a,
	0x4e, 0xa8, 0xe5, 0x9d, 0xd4, 0x45, 0x15, 0xdb, 0xfe, 0xdd, 0x03, 0xd8, 0xec, 0x94, 0xbb, 0xef,
	0x31, 0xb8, 0xe1, 0x3d, 0xfa, 0xf6, 0x3d, 0xbe, 0x72, 0xf7, 0x1d, 0x3e, 0x84, 0x7b, 0x64, 0x96,
	0x7c, 0x9e, 0x2a, 0x95, 0x75, 0xbc, 0x3a, 0x5d, 0xaf, 0x87, 0x8f, 0x3f, 0xf3, 0x8b, 0x54, 0x67,
	0xe2, 0x2c, 0x1a, 0x8d, 0xa2, 0x44, 0x3e, 0x4f, 0x16, 0x22, 0x95, 0xa3, 0xd1, 0xea, 0x93, 0xc4,
	0x9e, 0xf5, 0xe9, 0xaf, 0x76, 0xf4, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xe2, 0x23, 0x00, 0x8b,
	0x05, 0x00, 0x00,
}
