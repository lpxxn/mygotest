package main

import (
	"github.com/nsqio/go-nsq"
	"time"
	"log"
	"fmt"
	"strconv"
	"os"
	"os/signal"
)

func main () {

	topicName := "publishtest"
	msgCount := 2
	for i := 0; i < msgCount; i++ {
		//time.Sleep(time.Millisecond * 20)
		go readMessage(topicName, i)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	fmt.Println("server is running....")
	<-quit
	fmt.Println("Shutdown server....")

}

type ConsumerHandle struct {
	q *nsq.Consumer
	msgGood int
}

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

	//q, _ := nsq.NewConsumer(topicName, "ch" + strconv.Itoa(msgCount), config)
	//q, _ := nsq.NewConsumer(topicName, "ch" + strconv.Itoa(msgCount) + "#ephemeral", config)
	q, _ := nsq.NewConsumer(topicName, "ch" + strconv.Itoa(msgCount), config)

	h := &ConsumerHandle{q: q, msgGood:msgCount}
	q.AddHandler(h)

	err := q.ConnectToNSQLookupd("192.168.0.105:4161")
	//err := q.ConnectToNSQDs([]string{"192.168.0.105:4161"})
	//err := q.ConnectToNSQD("192.168.0.49:4150")
	//err := q.ConnectToNSQD("192.168.0.105:4415")
	if err != nil {
		fmt.Println("conect nsqd error")
		log.Println(err)
	}
	<-q.StopChan
	fmt.Println("end....")
}