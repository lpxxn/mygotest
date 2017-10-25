package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"github.com/pkg/errors"
	"log"
	"time"
)

func main() {
	config := nsq.NewConfig()
	// 随便给哪个ip发都可以
	w1, _ := nsq.NewProducer("192.168.0.105:4150", config)
	//w1, _ := nsq.NewProducer("192.168.0.49:4150", config)

	err1 := w1.Ping()
	if err1 != nil {
		log.Fatal("should not be able to ping after Stop()")
		return
	}
	defer w1.Stop()
	topicName := "publishtest"
	msgCount := 2
	for i := 1; i < msgCount; i++ {
		err1 := w1.Publish(topicName, []byte("测试测试publis test case"))
		if err1 != nil {
			log.Fatal("error")
		}
	}
}

type ConsumerHandle struct {
	q       *nsq.Consumer
	msgGood int
}

func (h *ConsumerHandle) HandleMessage(message *nsq.Message) error {
	msg := string(message.Body)
	fmt.Println(msg)
	if msg == "bad_test_case" {
		return errors.New("fail this message")
	}

	if msg != "multipublish_test_case" && msg != "public_test_case" {
		return errors.New("message was not corrent" + msg)
	}

	return nil
}

func readMessage(topicName string, msgCount int) {

	config := nsq.NewConfig()
	config.DefaultRequeueDelay = 0
	config.MaxBackoffDuration = 50 * time.Millisecond

	q, _ := nsq.NewConsumer(topicName, "ch", config)

	h := &ConsumerHandle{q: q}
	q.AddHandler(h)

	err := q.ConnectToNSQD("192.168.0.105:4150")
	if err != nil {
		log.Fatal(err)
	}

	<-q.StopChan
}
