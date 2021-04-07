package dig1

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestErrorType(t *testing.T) {
	_errType := reflect.TypeOf((*error)(nil))
	// *error   *error
	t.Log(_errType, " ", _errType.String())
	_errT := _errType.Elem()
	// error   error
	t.Log(_errT, " ", _errT.String())
}

func TestUintptr1(t *testing.T) {
	data := []byte("abcd")
	for i := 0; i < len(data); i++ {
		ptr := unsafe.Pointer(uintptr(unsafe.Pointer(&data[0])) + uintptr(i)*unsafe.Sizeof(data[0]))
		fmt.Printf("%c,", *(*byte)(unsafe.Pointer(ptr)))
	}
	fmt.Printf("\n")
}

/*
要理解上述代码，首选需要了解两个原则，分别是：

其他类型的指针只能转化为unsafe.Pointer，也只有unsafe.Pointer才能转化成任意类型的指针
只有uintptr才支持加减操作，而uintptr是一个非负整数，表示地址值，没有类型信息，以字节为单位
for循环的ptr赋值是该例子中的重点代码，它表示：

把data的第0个元素的地址，转化为unsafe.Pointer，再把它转换成uintptr，用于加减运算，即（uintptr(unsafe.Pointer(&data[0])) ）
加上第i个元素的偏移量，得到一个新的uintptr值，计算方法为i每个元素所占的字节数，即（+ uintptr(i)unsafe.Sizeof(data[0])）
把新的uintptr再转化为unsafe.Pointer，用于在后续的打印操作中，转化为实际类型的指针

 */
/*
https://zhuanlan.zhihu.com/p/240856451

能说说uintptr和unsafe.Pointer的区别吗？

怎么答
unsafe.Pointer只是单纯的通用指针类型，用于转换不同类型指针，它不可以参与指针运算；
而uintptr是用于指针运算的，GC 不把 uintptr 当指针，也就是说 uintptr 无法持有对象， uintptr 类型的目标会被回收；
unsafe.Pointer 可以和 普通指针 进行相互转换；
unsafe.Pointer 可以和 uintptr 进行相互转换。
*/
func TestUintptr(t *testing.T) {
	type W struct {
		b int32
		c int64
	}

	var w *W = new(W)
	//这时w的变量打印出来都是默认值0，0
	fmt.Println(w.b, w.c)

	//现在我们通过指针运算给b变量赋值为10
	b := unsafe.Pointer(uintptr(unsafe.Pointer(w)) + unsafe.Offsetof(w.b))
	*((*int)(b)) = 10
	//此时结果就变成了10，0
	fmt.Println(w.b, w.c)
}
/*
uintptr(unsafe.Pointer(w)) 获取了 w 的指针起始值
unsafe.Offsetof(w.b) 获取 b 变量的偏移量
两个相加就得到了 b 的地址值，将通用指针 Pointer 转换成具体指针 ((*int)(b))，通过 * 符号取值，然后赋值。*((*int)(b)) 相当于把 (*int)(b) 转换成 int 了，最后对变量重新赋值成 10，这样指针运算就完成了
 */
