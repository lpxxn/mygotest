package main

import (
	"fmt"
	"github.com/mygotest/gohashdemo/demo3/tiny"
)

func main() {
	completelyRandomSeed := "5SX0TEjkR1mLOw8Gvq2VyJxIFhgCAYidrclDWaM3so9bfzZpuUenKtP74QNH6B"
	tiny_v := tiny.NewTiny(completelyRandomSeed)

	fmt.Println(tiny_v.To(5))
	// E

	fmt.Println(tiny_v.From("E"))
	// 5

}
