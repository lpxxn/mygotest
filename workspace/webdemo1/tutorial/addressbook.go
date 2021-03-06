// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addressbook.proto

/*
Package tutorial is a generated protocol buffer package.

It is generated from these files:
	addressbook.proto

It has these top-level messages:
	Person
	AddressBook
*/
package tutorial

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 2
	Person_WORK   Person_PhoneType = 1
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	2: "HOME",
	1: "WORK",
}
var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   2,
	"WORK":   1,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}
func (Person_PhoneType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Person struct {
	Name   string                `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Id     int32                 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Email  string                `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Phones []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones" json:"phones,omitempty"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

type Person_PhoneNumber struct {
	Number string           `protobuf:"bytes,1,opt,name=number" json:"number,omitempty"`
	Type   Person_PhoneType `protobuf:"varint,2,opt,name=type,enum=tutorial.Person_PhoneType" json:"type,omitempty"`
}

func (m *Person_PhoneNumber) Reset()                    { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string            { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()               {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

type AddressBook struct {
	People []*Person `protobuf:"bytes,1,rep,name=people" json:"people,omitempty"`
}

func (m *AddressBook) Reset()                    { *m = AddressBook{} }
func (m *AddressBook) String() string            { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()               {}
func (*AddressBook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterType((*Person)(nil), "tutorial.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "tutorial.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "tutorial.AddressBook")
	proto.RegisterEnum("tutorial.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
}

func init() { proto.RegisterFile("addressbook.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xdd, 0x34, 0x5d, 0xda, 0x09, 0x94, 0x38, 0x88, 0x84, 0xe2, 0x21, 0xe4, 0x14, 0x10,
	0xf6, 0x50, 0x05, 0xcf, 0x16, 0x0a, 0x8a, 0xd6, 0x94, 0x45, 0xf1, 0x9c, 0x90, 0x01, 0x43, 0x93,
	0xcc, 0xb2, 0x49, 0x0f, 0xfd, 0xef, 0x1e, 0xa4, 0x9b, 0x28, 0x22, 0xde, 0xde, 0xcc, 0x7b, 0x7c,
	0xbb, 0xf3, 0xe0, 0x3c, 0x2f, 0x4b, 0x4b, 0x5d, 0x57, 0x30, 0xef, 0x95, 0xb1, 0xdc, 0x33, 0xce,
	0xfa, 0x43, 0xcf, 0xb6, 0xca, 0xeb, 0xe4, 0x53, 0x80, 0xdc, 0x91, 0xed, 0xb8, 0x45, 0x04, 0xbf,
	0xcd, 0x1b, 0x8a, 0x26, 0xb1, 0x48, 0xe7, 0xda, 0x69, 0x5c, 0x80, 0x57, 0x95, 0x91, 0x17, 0x8b,
	0x74, 0xaa, 0xbd, 0xaa, 0xc4, 0x0b, 0x98, 0x52, 0x93, 0x57, 0x75, 0x24, 0x5c, 0x68, 0x18, 0xf0,
	0x16, 0xa4, 0xf9, 0xe0, 0x96, 0xba, 0xc8, 0x8f, 0x27, 0x69, 0xb0, 0xba, 0x52, 0xdf, 0x7c, 0x35,
	0xb0, 0xd5, 0xee, 0x64, 0xbf, 0x1c, 0x9a, 0x82, 0xac, 0x1e, 0xb3, 0xcb, 0x37, 0x08, 0x7e, 0xad,
	0xf1, 0x12, 0x64, 0xeb, 0xd4, 0xc8, 0x1e, 0x27, 0x54, 0xe0, 0xf7, 0x47, 0x43, 0xee, 0x13, 0x8b,
	0xd5, 0xf2, 0x7f, 0xf4, 0xeb, 0xd1, 0x90, 0x76, 0xb9, 0xe4, 0x1a, 0xe6, 0x3f, 0x2b, 0x04, 0x90,
	0xdb, 0x6c, 0xfd, 0xf8, 0xbc, 0x09, 0xcf, 0x70, 0x06, 0xfe, 0x43, 0xb6, 0xdd, 0x84, 0xde, 0x49,
	0xbd, 0x67, 0xfa, 0x29, 0x14, 0xc9, 0x1d, 0x04, 0xf7, 0x43, 0x3b, 0x6b, 0xe6, 0x3d, 0xa6, 0x20,
	0x0d, 0xb1, 0xa9, 0x29, 0x12, 0xee, 0x90, 0xf0, 0xef, 0x6b, 0x7a, 0xf4, 0x0b, 0xe9, 0x8a, 0xbc,
	0xf9, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xac, 0x28, 0x14, 0xd0, 0x5d, 0x01, 0x00, 0x00,
}
