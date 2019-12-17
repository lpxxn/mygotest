package main

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"github.com/sony/sonyflake"
	"time"
)

var sf *sonyflake.Sonyflake
var sf2 *sonyflake.Sonyflake

func init() {
	var st sonyflake.Settings
	st.StartTime = time.Now()

	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}

	var st2 sonyflake.Settings
	st2.StartTime = time.Now()
	st2.MachineID = func() (u uint16, err error) {
		return 0, nil
	}
	st2.CheckMachineID = func(u uint16) bool {
		return true
	}
	sf2 = sonyflake.NewSonyflake(st2)
	if sf2 == nil {
		panic("sonyflake not created")
	}
}

func main() {
	consumer := make(chan uint64, 10)
	go func() {
		for {
			id, err := sf2.NextID()
			if err != nil {
				fmt.Printf("error: %#v", err)
				return
			}
			consumer <- id
		}
	}()
	go func() {
		for {
			id, err := sf.NextID()
			if err != nil {
				fmt.Printf("error: %#v", err)
				return
			}
			consumer <- id
		}
	}()
	set := mapset.NewSet()
	count := 0
	start := time.Now()
	for ; time.Since(start) < time.Second; {
		idx := <-consumer
		if set.Contains(idx) {
			fmt.Printf("duplicate idx %d", idx)
			panic(" duplicate idx")
		} else {
			set.Add(idx)
		}
		count++
		fmt.Println(idx, "  count: ", count)
	}
}
