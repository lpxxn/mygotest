// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gomicrodemo/greeter/srv/proto/common/common.proto

package a_b_common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RespType int32

const (
	RespType_NONE    RespType = 0
	RespType_ASCEND  RespType = 1
	RespType_DESCEND RespType = 2
)

var RespType_name = map[int32]string{
	0: "NONE",
	1: "ASCEND",
	2: "DESCEND",
}
var RespType_value = map[string]int32{
	"NONE":    0,
	"ASCEND":  1,
	"DESCEND": 2,
}

func (x RespType) String() string {
	return proto.EnumName(RespType_name, int32(x))
}
func (RespType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_c3265b1e33464c0a, []int{0}
}

type Request struct {
	Name                 string           `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Values               []int32          `protobuf:"varint,2,rep,packed,name=values" json:"values,omitempty"`
	Header               map[string]*Pair `protobuf:"bytes,3,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Type                 RespType         `protobuf:"varint,4,opt,name=type,enum=a.b.common.RespType" json:"type,omitempty"`
	Content              *any.Any         `protobuf:"bytes,5,opt,name=content" json:"content,omitempty"`
	Msg                  []byte           `protobuf:"bytes,6,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_c3265b1e33464c0a, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Request) GetValues() []int32 {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Request) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Request) GetType() RespType {
	if m != nil {
		return m.Type
	}
	return RespType_NONE
}

func (m *Request) GetContent() *any.Any {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Request) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

type Pair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Values               []string `protobuf:"bytes,2,rep,name=values" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pair) Reset()         { *m = Pair{} }
func (m *Pair) String() string { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()    {}
func (*Pair) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_c3265b1e33464c0a, []int{1}
}
func (m *Pair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pair.Unmarshal(m, b)
}
func (m *Pair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pair.Marshal(b, m, deterministic)
}
func (dst *Pair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pair.Merge(dst, src)
}
func (m *Pair) XXX_Size() int {
	return xxx_messageInfo_Pair.Size(m)
}
func (m *Pair) XXX_DiscardUnknown() {
	xxx_messageInfo_Pair.DiscardUnknown(m)
}

var xxx_messageInfo_Pair proto.InternalMessageInfo

func (m *Pair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Pair) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type Response struct {
	Msg                  string           `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
	Values               []string         `protobuf:"bytes,2,rep,name=values" json:"values,omitempty"`
	Header               map[string]*Pair `protobuf:"bytes,3,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Type                 RespType         `protobuf:"varint,4,opt,name=type,enum=a.b.common.RespType" json:"type,omitempty"`
	Msg2                 string           `protobuf:"bytes,5,opt,name=msg2" json:"msg2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_c3265b1e33464c0a, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Response) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Response) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Response) GetType() RespType {
	if m != nil {
		return m.Type
	}
	return RespType_NONE
}

func (m *Response) GetMsg2() string {
	if m != nil {
		return m.Msg2
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "a.b.common.Request")
	proto.RegisterMapType((map[string]*Pair)(nil), "a.b.common.Request.HeaderEntry")
	proto.RegisterType((*Pair)(nil), "a.b.common.Pair")
	proto.RegisterType((*Response)(nil), "a.b.common.Response")
	proto.RegisterMapType((map[string]*Pair)(nil), "a.b.common.Response.HeaderEntry")
	proto.RegisterEnum("a.b.common.RespType", RespType_name, RespType_value)
}

func init() {
	proto.RegisterFile("gomicrodemo/greeter/srv/proto/common/common.proto", fileDescriptor_common_c3265b1e33464c0a)
}

var fileDescriptor_common_c3265b1e33464c0a = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0x4f, 0xcb, 0xd3, 0x40,
	0x10, 0xc6, 0xdd, 0x24, 0x4d, 0xdb, 0x89, 0x48, 0x58, 0x8a, 0xc4, 0x5e, 0x0c, 0x3d, 0x48, 0x10,
	0xdc, 0x68, 0x3c, 0x58, 0xbc, 0x15, 0x1b, 0x10, 0x84, 0x2a, 0xab, 0x5f, 0x20, 0x6d, 0xc7, 0x58,
	0xec, 0xee, 0xc6, 0xdd, 0xb4, 0x90, 0x4f, 0xe4, 0x67, 0xf4, 0x26, 0xd9, 0x24, 0xd8, 0x16, 0xdf,
	0xd3, 0xfb, 0x9e, 0x32, 0x33, 0x79, 0xe6, 0xcf, 0xf3, 0x5b, 0x78, 0x53, 0x2a, 0x71, 0xd8, 0x69,
	0xb5, 0x47, 0xa1, 0xd2, 0x52, 0x23, 0xd6, 0xa8, 0x53, 0xa3, 0xcf, 0x69, 0xa5, 0x55, 0xad, 0xd2,
	0x9d, 0x12, 0x42, 0xc9, 0xfe, 0xc3, 0x6c, 0x8d, 0x42, 0xc1, 0xb6, 0xac, 0xab, 0xcc, 0x9f, 0x95,
	0x4a, 0x95, 0x47, 0xec, 0xd4, 0xdb, 0xd3, 0xf7, 0xb4, 0x90, 0x4d, 0x27, 0x5b, 0xfc, 0x76, 0x60,
	0xcc, 0xf1, 0xd7, 0x09, 0x4d, 0x4d, 0x29, 0x78, 0xb2, 0x10, 0x18, 0x91, 0x98, 0x24, 0x53, 0x6e,
	0x63, 0xfa, 0x14, 0xfc, 0x73, 0x71, 0x3c, 0xa1, 0x89, 0x9c, 0xd8, 0x4d, 0x46, 0xbc, 0xcf, 0xe8,
	0x3b, 0xf0, 0x7f, 0x60, 0xb1, 0x47, 0x1d, 0xb9, 0xb1, 0x9b, 0x04, 0xd9, 0x73, 0xf6, 0x6f, 0x1f,
	0xeb, 0x07, 0xb2, 0x8f, 0x56, 0x91, 0xcb, 0x5a, 0x37, 0xbc, 0x97, 0xd3, 0x04, 0xbc, 0xba, 0xa9,
	0x30, 0xf2, 0x62, 0x92, 0x3c, 0xc9, 0x66, 0xd7, 0x6d, 0xa6, 0xfa, 0xd6, 0x54, 0xc8, 0xad, 0x82,
	0x32, 0x18, 0xef, 0x94, 0xac, 0x51, 0xd6, 0xd1, 0x28, 0x26, 0x49, 0x90, 0xcd, 0x58, 0xe7, 0x83,
	0x0d, 0x3e, 0xd8, 0x4a, 0x36, 0x7c, 0x10, 0xd1, 0x10, 0x5c, 0x61, 0xca, 0xc8, 0x8f, 0x49, 0xf2,
	0x98, 0xb7, 0xe1, 0xfc, 0x13, 0x04, 0x17, 0x27, 0xb4, 0x82, 0x9f, 0xd8, 0xf4, 0xf6, 0xda, 0x90,
	0xbe, 0x80, 0x91, 0xf5, 0x13, 0x39, 0x76, 0x41, 0x78, 0x79, 0xcd, 0x97, 0xe2, 0xa0, 0x79, 0xf7,
	0xfb, 0xbd, 0xb3, 0x24, 0x8b, 0xd7, 0xe0, 0xb5, 0xa5, 0xff, 0x4c, 0xb9, 0x66, 0x34, 0x1d, 0x18,
	0x2d, 0xfe, 0x10, 0x98, 0xb4, 0x9e, 0x94, 0x34, 0x38, 0x5c, 0xd7, 0xb7, 0x09, 0x53, 0xde, 0xd5,
	0x46, 0x97, 0x37, 0x68, 0xe3, 0x5b, 0x46, 0xed, 0xbc, 0x7b, 0xb2, 0xa5, 0xe0, 0x09, 0x53, 0x66,
	0x16, 0xec, 0x94, 0xdb, 0xf8, 0x41, 0x69, 0xbd, 0x7c, 0xd5, 0x59, 0x6f, 0x57, 0xd2, 0x09, 0x78,
	0x9b, 0xcf, 0x9b, 0x3c, 0x7c, 0x44, 0x01, 0xfc, 0xd5, 0xd7, 0x0f, 0xf9, 0x66, 0x1d, 0x12, 0x1a,
	0xc0, 0x78, 0x9d, 0x77, 0x89, 0xb3, 0xf5, 0xed, 0x93, 0xbe, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff,
	0x8a, 0x8e, 0xc5, 0x40, 0xe9, 0x02, 0x00, 0x00,
}
