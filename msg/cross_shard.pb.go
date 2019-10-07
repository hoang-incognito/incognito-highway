// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cross_shard.proto

package msg

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetBlockCrossShardByHeightRequest struct {
	FromShard            int32    `protobuf:"varint,1,opt,name=FromShard,proto3" json:"FromShard,omitempty"`
	ToShard              int32    `protobuf:"varint,2,opt,name=ToShard,proto3" json:"ToShard,omitempty"`
	Specific             bool     `protobuf:"varint,3,opt,name=Specific,proto3" json:"Specific,omitempty"`
	FromHeight           uint64   `protobuf:"varint,4,opt,name=FromHeight,proto3" json:"FromHeight,omitempty"`
	ToHeight             uint64   `protobuf:"varint,5,opt,name=ToHeight,proto3" json:"ToHeight,omitempty"`
	Heights              []uint64 `protobuf:"varint,6,rep,packed,name=Heights,proto3" json:"Heights,omitempty"`
	FromPool             bool     `protobuf:"varint,7,opt,name=FromPool,proto3" json:"FromPool,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockCrossShardByHeightRequest) Reset()         { *m = GetBlockCrossShardByHeightRequest{} }
func (m *GetBlockCrossShardByHeightRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlockCrossShardByHeightRequest) ProtoMessage()    {}
func (*GetBlockCrossShardByHeightRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_780bd9b6271edabf, []int{0}
}

func (m *GetBlockCrossShardByHeightRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockCrossShardByHeightRequest.Unmarshal(m, b)
}
func (m *GetBlockCrossShardByHeightRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockCrossShardByHeightRequest.Marshal(b, m, deterministic)
}
func (m *GetBlockCrossShardByHeightRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockCrossShardByHeightRequest.Merge(m, src)
}
func (m *GetBlockCrossShardByHeightRequest) XXX_Size() int {
	return xxx_messageInfo_GetBlockCrossShardByHeightRequest.Size(m)
}
func (m *GetBlockCrossShardByHeightRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockCrossShardByHeightRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockCrossShardByHeightRequest proto.InternalMessageInfo

func (m *GetBlockCrossShardByHeightRequest) GetFromShard() int32 {
	if m != nil {
		return m.FromShard
	}
	return 0
}

func (m *GetBlockCrossShardByHeightRequest) GetToShard() int32 {
	if m != nil {
		return m.ToShard
	}
	return 0
}

func (m *GetBlockCrossShardByHeightRequest) GetSpecific() bool {
	if m != nil {
		return m.Specific
	}
	return false
}

func (m *GetBlockCrossShardByHeightRequest) GetFromHeight() uint64 {
	if m != nil {
		return m.FromHeight
	}
	return 0
}

func (m *GetBlockCrossShardByHeightRequest) GetToHeight() uint64 {
	if m != nil {
		return m.ToHeight
	}
	return 0
}

func (m *GetBlockCrossShardByHeightRequest) GetHeights() []uint64 {
	if m != nil {
		return m.Heights
	}
	return nil
}

func (m *GetBlockCrossShardByHeightRequest) GetFromPool() bool {
	if m != nil {
		return m.FromPool
	}
	return false
}

type GetBlockCrossShardByHeightResponse struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockCrossShardByHeightResponse) Reset()         { *m = GetBlockCrossShardByHeightResponse{} }
func (m *GetBlockCrossShardByHeightResponse) String() string { return proto.CompactTextString(m) }
func (*GetBlockCrossShardByHeightResponse) ProtoMessage()    {}
func (*GetBlockCrossShardByHeightResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_780bd9b6271edabf, []int{1}
}

func (m *GetBlockCrossShardByHeightResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockCrossShardByHeightResponse.Unmarshal(m, b)
}
func (m *GetBlockCrossShardByHeightResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockCrossShardByHeightResponse.Marshal(b, m, deterministic)
}
func (m *GetBlockCrossShardByHeightResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockCrossShardByHeightResponse.Merge(m, src)
}
func (m *GetBlockCrossShardByHeightResponse) XXX_Size() int {
	return xxx_messageInfo_GetBlockCrossShardByHeightResponse.Size(m)
}
func (m *GetBlockCrossShardByHeightResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockCrossShardByHeightResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockCrossShardByHeightResponse proto.InternalMessageInfo

func (m *GetBlockCrossShardByHeightResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetBlockCrossShardByHashRequest struct {
	FromShard            int32    `protobuf:"varint,1,opt,name=FromShard,proto3" json:"FromShard,omitempty"`
	ToShard              int32    `protobuf:"varint,2,opt,name=ToShard,proto3" json:"ToShard,omitempty"`
	Hashes               [][]byte `protobuf:"bytes,3,rep,name=Hashes,proto3" json:"Hashes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockCrossShardByHashRequest) Reset()         { *m = GetBlockCrossShardByHashRequest{} }
func (m *GetBlockCrossShardByHashRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlockCrossShardByHashRequest) ProtoMessage()    {}
func (*GetBlockCrossShardByHashRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_780bd9b6271edabf, []int{2}
}

func (m *GetBlockCrossShardByHashRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockCrossShardByHashRequest.Unmarshal(m, b)
}
func (m *GetBlockCrossShardByHashRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockCrossShardByHashRequest.Marshal(b, m, deterministic)
}
func (m *GetBlockCrossShardByHashRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockCrossShardByHashRequest.Merge(m, src)
}
func (m *GetBlockCrossShardByHashRequest) XXX_Size() int {
	return xxx_messageInfo_GetBlockCrossShardByHashRequest.Size(m)
}
func (m *GetBlockCrossShardByHashRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockCrossShardByHashRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockCrossShardByHashRequest proto.InternalMessageInfo

func (m *GetBlockCrossShardByHashRequest) GetFromShard() int32 {
	if m != nil {
		return m.FromShard
	}
	return 0
}

func (m *GetBlockCrossShardByHashRequest) GetToShard() int32 {
	if m != nil {
		return m.ToShard
	}
	return 0
}

func (m *GetBlockCrossShardByHashRequest) GetHashes() [][]byte {
	if m != nil {
		return m.Hashes
	}
	return nil
}

type GetBlockCrossShardByHashResponse struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockCrossShardByHashResponse) Reset()         { *m = GetBlockCrossShardByHashResponse{} }
func (m *GetBlockCrossShardByHashResponse) String() string { return proto.CompactTextString(m) }
func (*GetBlockCrossShardByHashResponse) ProtoMessage()    {}
func (*GetBlockCrossShardByHashResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_780bd9b6271edabf, []int{3}
}

func (m *GetBlockCrossShardByHashResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockCrossShardByHashResponse.Unmarshal(m, b)
}
func (m *GetBlockCrossShardByHashResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockCrossShardByHashResponse.Marshal(b, m, deterministic)
}
func (m *GetBlockCrossShardByHashResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockCrossShardByHashResponse.Merge(m, src)
}
func (m *GetBlockCrossShardByHashResponse) XXX_Size() int {
	return xxx_messageInfo_GetBlockCrossShardByHashResponse.Size(m)
}
func (m *GetBlockCrossShardByHashResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockCrossShardByHashResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockCrossShardByHashResponse proto.InternalMessageInfo

func (m *GetBlockCrossShardByHashResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetBlockCrossShardByHeightRequest)(nil), "msg.GetBlockCrossShardByHeightRequest")
	proto.RegisterType((*GetBlockCrossShardByHeightResponse)(nil), "msg.GetBlockCrossShardByHeightResponse")
	proto.RegisterType((*GetBlockCrossShardByHashRequest)(nil), "msg.GetBlockCrossShardByHashRequest")
	proto.RegisterType((*GetBlockCrossShardByHashResponse)(nil), "msg.GetBlockCrossShardByHashResponse")
}

func init() { proto.RegisterFile("cross_shard.proto", fileDescriptor_780bd9b6271edabf) }

var fileDescriptor_780bd9b6271edabf = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x15, 0x9c, 0xa4, 0xe5, 0xd4, 0x05, 0x0f, 0xc8, 0x42, 0x08, 0x8c, 0x27, 0x4f, 0x2c,
	0x48, 0x88, 0x39, 0x20, 0x60, 0x44, 0x6e, 0x27, 0x16, 0x14, 0x82, 0x69, 0x22, 0x5a, 0x2e, 0xcd,
	0x99, 0x81, 0x47, 0xe6, 0x2d, 0xd0, 0x39, 0x4d, 0x61, 0xa0, 0x2c, 0xdd, 0xee, 0xbb, 0xff, 0xee,
	0x7e, 0x9f, 0x0f, 0x0e, 0xaa, 0x0e, 0x89, 0x9e, 0xa8, 0x2e, 0xbb, 0x97, 0xf3, 0xb6, 0xc3, 0x80,
	0x52, 0x2c, 0x69, 0x6e, 0xbe, 0x12, 0x38, 0xbb, 0xf3, 0xa1, 0x58, 0x60, 0xf5, 0x76, 0xcd, 0x25,
	0x53, 0xae, 0x28, 0x3e, 0xef, 0x7d, 0x33, 0xaf, 0x83, 0xf3, 0xab, 0x0f, 0x4f, 0x41, 0x1e, 0xc3,
	0xfe, 0x6d, 0x87, 0xcb, 0xa8, 0xa9, 0x44, 0x27, 0x36, 0x73, 0x3f, 0x09, 0xa9, 0x60, 0x34, 0xc3,
	0x5e, 0xdb, 0x8b, 0xda, 0x80, 0xf2, 0x08, 0xc6, 0xd3, 0xd6, 0x57, 0xcd, 0x6b, 0x53, 0x29, 0xa1,
	0x13, 0x3b, 0x76, 0x1b, 0x96, 0x27, 0x00, 0x3c, 0xa2, 0x37, 0x52, 0xa9, 0x4e, 0x6c, 0xea, 0x7e,
	0x65, 0xb8, 0x77, 0x86, 0x6b, 0x35, 0x8b, 0xea, 0x86, 0xd9, 0xb1, 0x8f, 0x48, 0xe5, 0x5a, 0xd8,
	0xd4, 0x0d, 0xc8, 0x5d, 0x3c, 0xe3, 0x01, 0x71, 0xa1, 0x46, 0xbd, 0xe3, 0xc0, 0xe6, 0x0a, 0xcc,
	0x7f, 0xab, 0x52, 0x8b, 0xef, 0xe4, 0xa5, 0x84, 0xf4, 0xa6, 0x0c, 0x65, 0x5c, 0x73, 0xe2, 0x62,
	0x6c, 0x56, 0x70, 0xfa, 0x67, 0x67, 0x49, 0xf5, 0xae, 0x5f, 0x74, 0x08, 0x39, 0x8f, 0xf1, 0xa4,
	0x84, 0x16, 0x76, 0xe2, 0xd6, 0x64, 0x2e, 0x41, 0x6f, 0xb7, 0xdc, 0xfe, 0xd4, 0x22, 0x7b, 0xe4,
	0xbb, 0x3e, 0xe7, 0xf1, 0xc6, 0x17, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x51, 0x53, 0xb6,
	0xf8, 0x01, 0x00, 0x00,
}
