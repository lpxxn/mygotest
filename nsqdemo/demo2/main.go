package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

func main() {

	topicName := "publishtest"
	//topicName := "GroupSink"
	msgCount := 2
	for i := 0; i < msgCount; i++ {
		//time.Sleep(time.Millisecond * 20)
		go readMessage(topicName, i)
	}

	//cleanup := make(chan os.Signal, 1)
	cleanup := make(chan os.Signal)
	signal.Notify(cleanup, os.Interrupt)
	fmt.Println("server is running....")

	quit := make(chan bool)
	go func() {
		//for _ = range cleanup {
		//	fmt.Println("Received an interrupt , stoping service ...")
		//	quit <- true
		//}
		select {
		case <-cleanup:
			fmt.Println("Received an interrupt , stoping service ...")
			for _, ele := range consumers {
				ele.StopChan <- 1
				ele.Stop()
			}
			quit <- true
		}
	}()
	<-quit
	fmt.Println("Shutdown server....")
}

type ConsumerHandle struct {
	q       *nsq.Consumer
	msgGood int
}

var consumers []*nsq.Consumer = make([]*nsq.Consumer, 0)
var mux *sync.Mutex = &sync.Mutex{}

func (h *ConsumerHandle) HandleMessage(message *nsq.Message) error {
	msg := string(message.Body) + "  " + strconv.Itoa(h.msgGood)
	fmt.Println(msg)

	return nil
}

func readMessage(topicName string, msgCount int) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error: ", err)
		}
	}()

	config := nsq.NewConfig()
	config.MaxInFlight = 1000
	config.MaxBackoffDuration = 500 * time.Second
	config.DialTimeout = 10 * time.Second

	//q, _ := nsq.NewConsumer(topicName, "ch" + strconv.Itoa(msgCount), config)
	//q, _ := nsq.NewConsumer(topicName, "ch" + strconv.Itoa(msgCount) + "#ephemeral", config)
	q, _ := nsq.NewConsumer(topicName, "ch"+strconv.Itoa(msgCount), config)

	h := &ConsumerHandle{q: q, msgGood: msgCount}
	q.AddHandler(h)

	err := q.ConnectToNSQLookupd("13.125.77.114:9002")
	//err := q.ConnectToNSQLookupd("192.168.0.105:4161")
	//err := q.ConnectToNSQDs([]string{"192.168.0.105:4161"})
	//err := q.ConnectToNSQD("192.168.0.49:4150")
	//err := q.ConnectToNSQD("192.168.0.105:4415")
	if err != nil {
		fmt.Println("conect nsqd error")
		log.Println(err)
	}
	mux.Lock()
	consumers = append(consumers, q)
	mux.Unlock()
	<-q.StopChan
	fmt.Println("end....")
}
