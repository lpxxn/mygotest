package tutorial

import (
	"fmt"
	"math/bits"
	"reflect"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang/protobuf/proto"
)

func TestRWpb(t *testing.T) {
	s := "21342345"
	t.Log(strings.Trim(s, "0123456789"))
	studentList := &StudentList{
		Class:   "一",
		Teacher: "李老师",
		Students: []*Student{&Student{Id: 123121212211212, Name: "床前明月光，疑是地上霜", Age: 10},
			{Id: 789, Name: "一去二三里，山村四五家", Age: 12},
		},
	}
	body, err := proto.Marshal(studentList)
	if err != nil {
		t.Error(err)
	}
	t.Log(body)

	rev := &StudentList{}
	err = proto.Unmarshal(body, rev)
	if err != nil {
		t.Error(err)
	}
	spew.Dump(rev)
}

func TestOffset(t *testing.T) {
	sl1 := new(Student)
	t.Logf("%p\n", &sl1.Name)
	t.Logf("%p\n", &sl1.Age)
	typ := reflect.TypeOf(Student{})
	t.Logf("Struct is %d bytes long\n", typ.Size())
	n := typ.NumField()
	for i := 0; i < n; i++ {
		field := typ.Field(i)
		t.Logf("%s at offset %v, size=%d, align=%d\n",
			field.Name, field.Offset, field.Type.Size(),
			field.Type.Align())
	}

	typFunc := func(name string, typ reflect.Type) {
		t.Log("----------------")
		t.Logf("%s is %d bytes long\n", name, typ.Size())
		n = typ.NumField()
		for i := 0; i < n; i++ {
			field := typ.Field(i)
			t.Logf("%s at offset %v, size=%d, align=%d\n",
				field.Name, field.Offset, field.Type.Size(),
				field.Type.Align())
		}
	}
	typFunc("StructList", reflect.TypeOf(StudentList{}))
	typFunc("School", reflect.TypeOf(School{}))

	// 为什么string是16个字节呢？因为string的结构包含两个域，
	//一个是指向Data的指针，占8个字节，一个是表示string长度的len，占8个字节
}

func TestSizeVarint(t *testing.T) {
	t.Log(SizeVarint(123456))
}
// SizeVarint returns the encoded size of a varint.
// The size is guaranteed to be within 1 and 10, inclusive.
func SizeVarint(v uint64) int {
	// This computes 1 + (bits.Len64(v)-1)/7.
	// 9/64 is a good enough approximation of 1/7
	bitLen := bits.Len64(v)
	fmt.Println(9/64)
	fmt.Println(bitLen)
	return int(9*uint32(bitLen)+64) / 64
}
