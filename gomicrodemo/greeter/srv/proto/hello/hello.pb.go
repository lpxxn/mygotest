// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello/hello.proto

package go_micro_srv_greeter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/mygotest/gomicrodemo/greeter/srv/proto/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

func init() { proto.RegisterFile("hello/hello.proto", fileDescriptor_hello_b711b2f51dfaa0cc) }

var fileDescriptor_hello_b711b2f51dfaa0cc = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0xcd, 0xb1, 0x0a, 0xc2, 0x40,
	0x0c, 0xc6, 0x71, 0x41, 0x74, 0xe8, 0x66, 0xed, 0xd4, 0xd1, 0x07, 0x48, 0x40, 0xdd, 0x9c, 0x05,
	0x67, 0x7d, 0x82, 0xf6, 0x0c, 0xd7, 0x42, 0xaf, 0x5f, 0xbd, 0xa4, 0x85, 0xbe, 0xbd, 0xf4, 0xec,
	0xe2, 0x92, 0x40, 0xf8, 0xf3, 0x4b, 0x76, 0x68, 0xa4, 0xeb, 0xc0, 0x69, 0xd2, 0x10, 0x61, 0xc8,
	0x0b, 0x0f, 0x0a, 0xad, 0x8b, 0x20, 0x8d, 0x13, 0xf9, 0x28, 0x62, 0x12, 0xcb, 0xbb, 0x6f, 0xad,
	0x19, 0x6b, 0x72, 0x08, 0x1c, 0x66, 0x0f, 0x13, 0x35, 0xf6, 0x48, 0xe1, 0x5b, 0x02, 0x78, 0x0d,
	0x59, 0xe3, 0xc4, 0x09, 0x62, 0x87, 0x10, 0xd0, 0xaf, 0xeb, 0x87, 0x9f, 0x6f, 0xd9, 0xf6, 0x55,
	0xcd, 0xf9, 0x35, 0xdb, 0x3d, 0x96, 0x97, 0xf9, 0x91, 0x2a, 0x4a, 0xe8, 0x92, 0x3c, 0xe5, 0x33,
	0x8a, 0x5a, 0x59, 0xfc, 0x1f, 0x75, 0x40, 0xaf, 0x72, 0xda, 0xd4, 0xfb, 0x64, 0x5c, 0xbe, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x01, 0x3e, 0xa8, 0x55, 0xb5, 0x00, 0x00, 0x00,
}