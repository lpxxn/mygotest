package main

import (
	"fmt"
	"runtime"
)

func main() {
	for skip := 0; ; skip++ {
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}

		p := runtime.FuncForPC(pc)
		fnfile, fnline := p.FileLine(0)

		fmt.Printf("skip = %d, pc = 0x%08X\n", skip, pc)
		fmt.Printf("  func: file = %s, line = L%03d, name = %s, entry = 0x%08X\n", fnfile, fnline, p.Name(), p.Entry())
		fmt.Printf("  call: file = %s, line = L%03d\n", file, line)
	}
}

/*
https://chai2010.cn/advanced-go-programming-book/ch3-asm/ch3-06-func-again.html
其中runtime.Caller先获取当时的PC寄存器值，以及文件和行号。然后根据PC寄存器表示的指令位置，
通过runtime.FuncForPC函数获取函数的基本信息。Go语言是如何实现这种特性的呢？

Go语言作为一门静态编译型语言，在执行时每个函数的地址都是固定的，函数的每条指令也是固定的。如果针对每个函数和函数的每个指令生成一个地址表格（也叫PC表格），那么在运行时我们就可以根据PC寄存器的值轻松查询到指令当时对应的函数和位置信息。而Go语言也是采用类似的策略，只不过地址表格经过裁剪，舍弃了不必要的信息。因为要在运行时获取任意一个地址的位置，必然是要有一个函数调用，因此我们只需要为函数的开始和结束位置，以及每个函数调用位置生成地址表格就可以了。同时地址是有大小顺序的，在排序后可以通过只记录增量来减少数据的大小；在查询时可以通过二分法加快查找的速度。
 */
