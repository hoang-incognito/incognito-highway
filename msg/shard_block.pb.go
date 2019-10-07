// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shard_block.proto

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

type GetBlockShardByHeightRequest struct {
	Shard                int32    `protobuf:"varint,1,opt,name=Shard,proto3" json:"Shard,omitempty"`
	Specific             bool     `protobuf:"varint,2,opt,name=Specific,proto3" json:"Specific,omitempty"`
	FromHeight           uint64   `protobuf:"varint,3,opt,name=FromHeight,proto3" json:"FromHeight,omitempty"`
	ToHeight             uint64   `protobuf:"varint,4,opt,name=ToHeight,proto3" json:"ToHeight,omitempty"`
	Heights              []uint64 `protobuf:"varint,5,rep,packed,name=Heights,proto3" json:"Heights,omitempty"`
	FromPool             bool     `protobuf:"varint,6,opt,name=FromPool,proto3" json:"FromPool,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockShardByHeightRequest) Reset()         { *m = GetBlockShardByHeightRequest{} }
func (m *GetBlockShardByHeightRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlockShardByHeightRequest) ProtoMessage()    {}
func (*GetBlockShardByHeightRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc8063b71d9d45c7, []int{0}
}

func (m *GetBlockShardByHeightRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockShardByHeightRequest.Unmarshal(m, b)
}
func (m *GetBlockShardByHeightRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockShardByHeightRequest.Marshal(b, m, deterministic)
}
func (m *GetBlockShardByHeightRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockShardByHeightRequest.Merge(m, src)
}
func (m *GetBlockShardByHeightRequest) XXX_Size() int {
	return xxx_messageInfo_GetBlockShardByHeightRequest.Size(m)
}
func (m *GetBlockShardByHeightRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockShardByHeightRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockShardByHeightRequest proto.InternalMessageInfo

func (m *GetBlockShardByHeightRequest) GetShard() int32 {
	if m != nil {
		return m.Shard
	}
	return 0
}

func (m *GetBlockShardByHeightRequest) GetSpecific() bool {
	if m != nil {
		return m.Specific
	}
	return false
}

func (m *GetBlockShardByHeightRequest) GetFromHeight() uint64 {
	if m != nil {
		return m.FromHeight
	}
	return 0
}

func (m *GetBlockShardByHeightRequest) GetToHeight() uint64 {
	if m != nil {
		return m.ToHeight
	}
	return 0
}

func (m *GetBlockShardByHeightRequest) GetHeights() []uint64 {
	if m != nil {
		return m.Heights
	}
	return nil
}

func (m *GetBlockShardByHeightRequest) GetFromPool() bool {
	if m != nil {
		return m.FromPool
	}
	return false
}

type GetBlockShardByHeightResponse struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockShardByHeightResponse) Reset()         { *m = GetBlockShardByHeightResponse{} }
func (m *GetBlockShardByHeightResponse) String() string { return proto.CompactTextString(m) }
func (*GetBlockShardByHeightResponse) ProtoMessage()    {}
func (*GetBlockShardByHeightResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc8063b71d9d45c7, []int{1}
}

func (m *GetBlockShardByHeightResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockShardByHeightResponse.Unmarshal(m, b)
}
func (m *GetBlockShardByHeightResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockShardByHeightResponse.Marshal(b, m, deterministic)
}
func (m *GetBlockShardByHeightResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockShardByHeightResponse.Merge(m, src)
}
func (m *GetBlockShardByHeightResponse) XXX_Size() int {
	return xxx_messageInfo_GetBlockShardByHeightResponse.Size(m)
}
func (m *GetBlockShardByHeightResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockShardByHeightResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockShardByHeightResponse proto.InternalMessageInfo

func (m *GetBlockShardByHeightResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetBlockShardByHashRequest struct {
	Shard                int32    `protobuf:"varint,1,opt,name=Shard,proto3" json:"Shard,omitempty"`
	Hashes               [][]byte `protobuf:"bytes,2,rep,name=Hashes,proto3" json:"Hashes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockShardByHashRequest) Reset()         { *m = GetBlockShardByHashRequest{} }
func (m *GetBlockShardByHashRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlockShardByHashRequest) ProtoMessage()    {}
func (*GetBlockShardByHashRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc8063b71d9d45c7, []int{2}
}

func (m *GetBlockShardByHashRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockShardByHashRequest.Unmarshal(m, b)
}
func (m *GetBlockShardByHashRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockShardByHashRequest.Marshal(b, m, deterministic)
}
func (m *GetBlockShardByHashRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockShardByHashRequest.Merge(m, src)
}
func (m *GetBlockShardByHashRequest) XXX_Size() int {
	return xxx_messageInfo_GetBlockShardByHashRequest.Size(m)
}
func (m *GetBlockShardByHashRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockShardByHashRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockShardByHashRequest proto.InternalMessageInfo

func (m *GetBlockShardByHashRequest) GetShard() int32 {
	if m != nil {
		return m.Shard
	}
	return 0
}

func (m *GetBlockShardByHashRequest) GetHashes() [][]byte {
	if m != nil {
		return m.Hashes
	}
	return nil
}

type GetBlockShardByHashResponse struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockShardByHashResponse) Reset()         { *m = GetBlockShardByHashResponse{} }
func (m *GetBlockShardByHashResponse) String() string { return proto.CompactTextString(m) }
func (*GetBlockShardByHashResponse) ProtoMessage()    {}
func (*GetBlockShardByHashResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc8063b71d9d45c7, []int{3}
}

func (m *GetBlockShardByHashResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockShardByHashResponse.Unmarshal(m, b)
}
func (m *GetBlockShardByHashResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockShardByHashResponse.Marshal(b, m, deterministic)
}
func (m *GetBlockShardByHashResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockShardByHashResponse.Merge(m, src)
}
func (m *GetBlockShardByHashResponse) XXX_Size() int {
	return xxx_messageInfo_GetBlockShardByHashResponse.Size(m)
}
func (m *GetBlockShardByHashResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockShardByHashResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockShardByHashResponse proto.InternalMessageInfo

func (m *GetBlockShardByHashResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetBlockShardByHeightRequest)(nil), "msg.GetBlockShardByHeightRequest")
	proto.RegisterType((*GetBlockShardByHeightResponse)(nil), "msg.GetBlockShardByHeightResponse")
	proto.RegisterType((*GetBlockShardByHashRequest)(nil), "msg.GetBlockShardByHashRequest")
	proto.RegisterType((*GetBlockShardByHashResponse)(nil), "msg.GetBlockShardByHashResponse")
}

func init() { proto.RegisterFile("shard_block.proto", fileDescriptor_bc8063b71d9d45c7) }

var fileDescriptor_bc8063b71d9d45c7 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0xce, 0x48, 0x2c,
	0x4a, 0x89, 0x4f, 0xca, 0xc9, 0x4f, 0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce,
	0x2d, 0x4e, 0x57, 0x3a, 0xc4, 0xc8, 0x25, 0xe3, 0x9e, 0x5a, 0xe2, 0x04, 0x12, 0x0f, 0x06, 0x29,
	0x71, 0xaa, 0xf4, 0x48, 0xcd, 0x4c, 0xcf, 0x28, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11,
	0x12, 0xe1, 0x62, 0x05, 0x8b, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x41, 0x38, 0x42, 0x52,
	0x5c, 0x1c, 0xc1, 0x05, 0xa9, 0xc9, 0x99, 0x69, 0x99, 0xc9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x1c,
	0x41, 0x70, 0xbe, 0x90, 0x1c, 0x17, 0x97, 0x5b, 0x51, 0x7e, 0x2e, 0xc4, 0x18, 0x09, 0x66, 0x05,
	0x46, 0x0d, 0x96, 0x20, 0x24, 0x11, 0x90, 0xde, 0x90, 0x7c, 0xa8, 0x2c, 0x0b, 0x58, 0x16, 0xce,
	0x17, 0x92, 0xe0, 0x62, 0x87, 0xb0, 0x8a, 0x25, 0x58, 0x15, 0x98, 0x35, 0x58, 0x82, 0x60, 0x5c,
	0x90, 0x2e, 0x90, 0x19, 0x01, 0xf9, 0xf9, 0x39, 0x12, 0x6c, 0x10, 0x1b, 0x61, 0x7c, 0x25, 0x63,
	0x2e, 0x59, 0x1c, 0x7e, 0x28, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x12, 0xe2, 0x62, 0x71, 0x49,
	0x2c, 0x49, 0x04, 0xfb, 0x81, 0x27, 0x08, 0xcc, 0x56, 0xf2, 0xe2, 0x92, 0x42, 0xd7, 0x94, 0x58,
	0x9c, 0x81, 0xdf, 0xdb, 0x62, 0x5c, 0x6c, 0x20, 0x45, 0xa9, 0xc5, 0x12, 0x4c, 0x0a, 0xcc, 0x1a,
	0x3c, 0x41, 0x50, 0x9e, 0x92, 0x21, 0x97, 0x34, 0x56, 0xb3, 0x70, 0x5b, 0xef, 0xc4, 0x1a, 0x05,
	0x0a, 0xff, 0x24, 0x36, 0x70, 0x5c, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x97, 0xce, 0xcf,
	0x27, 0xa0, 0x01, 0x00, 0x00,
}
