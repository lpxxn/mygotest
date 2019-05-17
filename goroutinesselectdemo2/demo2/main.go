package main

import (
	"fmt"

	gerr "github.com/go-errors/errors"
	"github.com/mygotest/goreflect/callerinfo"
)

func main() {
	ch1 := make(chan *int, 1)
	defer func() {
		fmt.Print("recover \n")
		if rec := recover(); rec != nil {
			errStr := gerr.Wrap(rec, 2).ErrorStack()
			fmt.Println(errStr)
		}
		v2 := 12
		ch1 <- &v2
	}()

	v1 := 1
	ch1 <- &v1

	var vint *int
	select {
	case vint = <-ch1:
		fmt.Println("run <-ch1")
	default:
		fmt.Println("run default")
		v2 := 2323
		vint = &v2
	}
	fmt.Println(vint, *vint)
	select {
	case vint = <-ch1:
		fmt.Println("run <-ch1")
	default:
		fmt.Println("run default")
		v2 := 2323
		vint = &v2
	}
	fmt.Println(vint, *vint)
	rev := callerinfo.RetrieveCallInfo()
	fmt.Printf("%#v \n", rev)

	panic("hehe")
	<-ch1
}

// go build -ldflags "-s -w"
