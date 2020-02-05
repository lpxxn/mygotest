package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	type TTT struct {
		ctx context.Context
	}
	ctx := context.Background()
	ctt, _ := context.WithTimeout(ctx, 4*time.Second)
	fmt.Printf("now == %v\n", time.Now())
	ttt := TTT{
		ctx: ctt,
	}
	payload, _ := json.Marshal(ttt)
	fmt.Println("--------")
	fmt.Println(string(payload))
	time.Sleep(1 * time.Second)

	ppp := new(TTT)
	ppp.ctx, _ = context.WithTimeout(context.Background(), time.Second)
	fmt.Println(ctx)

	err := json.Unmarshal(payload, ppp)
	if ppp.ctx == nil {
		fmt.Println("aaa")
	}
	fmt.Println(ctx)

	fmt.Println(err)
	select {
	case <-ppp.ctx.Done():
		fmt.Printf("now == %v", time.Now())
	default:
		fmt.Printf("wandanle")
	}
}
