package main

import (
	"flag"
	"fmt"
	sd "github.com/mygotest/servicediscovery/demo1/lib"
	"log"
	"time"
)

func main() {
	name := flag.String("name", fmt.Sprintf("%d", time.Now().Unix()), "des")
	extInfo := "lhq-demo..."

	flag.Parse()
	w, err := sd.NewWorker("sd-test", *name, extInfo, []string{
		"http://192.168.3.34:2379",
		"http://192.168.3.18:2379",
		"http://192.168.3.110:2379",
	})
	if err != nil {
		log.Fatal(err)
	}
	w.Register()
	log.Println("name ->", *name, "extInfo ->", extInfo)

	go func() {
		time.Sleep(time.Second * 60)
		w.Unregister()
	}()

	for {
		log.Println("isActive ->", w.IsActive())
		log.Println("isStop ->", w.IsStop())
		time.Sleep(time.Second * 2)
		//服务退出
		if w.IsStop() {
			return
		}
	}
}

