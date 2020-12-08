package tutorial

import (
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang/protobuf/proto"
)

func TestRWpb(t *testing.T) {
	studentList := &StudentList{
		Students: []*Student{&Student{Name: "li", Age: 10},
			{Name: "peng", Age: 12},
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

	t.Log("----------------")

	typ = reflect.TypeOf(StudentList{})
	t.Logf("Struct is %d bytes long\n", typ.Size())
	n = typ.NumField()
	for i := 0; i < n; i++ {
		field := typ.Field(i)
		t.Logf("%s at offset %v, size=%d, align=%d\n",
			field.Name, field.Offset, field.Type.Size(),
			field.Type.Align())
	}
	// 为什么string是16个字节呢？因为string的结构包含两个域，
	//一个是指向Data的指针，占8个字节，一个是表示string长度的len，占8个字节
}
