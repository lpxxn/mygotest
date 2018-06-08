package main

import (
	"log"
	"time"

	sd "github.com/mygotest/servicediscovery/demo1/lib"
)

func main() {
	m, err := sd.NewService("sd-test", []string{
		"http://192.168.3.34:2379",
		"http://192.168.3.18:2379",
		"http://192.168.3.110:2379",
	})
	if err != nil {
		log.Fatal(err)
	}
	for {
		log.Println("all ->", m.GetNodes())
		log.Println("all(strictly) ->", m.GetNodesStrictly())
		time.Sleep(time.Second * 2)
	}
}