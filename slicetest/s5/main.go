package main

import (
	"fmt"
	"time"
)


func main() {

	go test()
	time.Sleep(time.Minute)
}

func test() {
	go transfer()

}

func transfer() {
	defer deferRecove()
	uid, cid := getCommunitiesByWallet()

	fmt.Println("uid: ", uid, "  cid: ", cid)
	if uid == int64(0) {
		panic("transfer: transfer uid empty")
	}
	fmt.Println("run after")
}

func getCommunitiesByWallet()(int64, int64) {
	defer deferRecove()

	panic("getCommunitiesByWallet error")

	return 1, 2
}


func deferRecove() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}