package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"sync"
	"time"
)

var ConsumersInfo *CountConsumer = &CountConsumer{Consumers: make([]*nsq.Consumer, 0)}

func main() {
	go readMtMsg("GroupSink", "chgroup1#ephemeral")
	go readMtMsg("MtOrderSubscribe", "order#ephemeral") //
	go readMtMsg("MtDealSubscribe", "dealscribe")

	cleanup := make(chan os.Signal)
	signal.Notify(cleanup, os.Interrupt)
	fmt.Println("Lieten Msg......")

	quit := make(chan bool)

	go func() {
		for _ = range cleanup {
			fmt.Println("Receive an interrupt, stop listen Msg")
			ConsumersInfo.StopAllConsumers()
			quit <- true
		}
	}()

	<-quit
	fmt.Println("stop")
}

func readMtMsg(topicName, channelName string) {
	defer func() {
		fmt.Println("error , topicName :", topicName, " channel name : ", channelName)
		if err := recover(); err != nil {
			fmt.Println("error: ", err)
		}
	}()

	config := nsq.NewConfig()

	config.MaxInFlight = 1000
	config.MaxBackoffDuration = time.Second * 500

	q, err := nsq.NewConsumer(topicName, channelName, config)
	if err != nil {
		panic(err)
	}
	handler := &MyTestHandler{Consumer: q, TopicName: topicName, ChannelName: channelName}
	q.AddHandler(handler)
	ConsumersInfo.Add(q)

	err = q.ConnectToNSQLookupd("192.168.0.105:4161")

	if err != nil {
		fmt.Println("connect nsqd error :", err)
		panic(err)
	}

	<-q.StopChan
	fmt.Println("end .....")
}

type MyTestHandler struct {
	Consumer    *nsq.Consumer
	TopicName   string
	ChannelName string
}

func (handler *MyTestHandler) HandleMessage(message *nsq.Message) error {

	msg := string(message.Body)
	fmt.Printf("TopicName : %s, ChannelName : %s, Msg: %s \n", handler.TopicName, handler.ChannelName, msg)
	return nil
}

type CountConsumer struct {
	sync.Mutex
	Consumers []*nsq.Consumer
}

func (self CountConsumer) Add(consumer *nsq.Consumer) {
	self.Lock()
	self.Consumers = append(self.Consumers, consumer)
	self.Unlock()
}

func (self CountConsumer) Get(index int) *nsq.Consumer {
	self.Lock()
	defer self.Unlock()
	return self.Consumers[index]
}

func (self CountConsumer) StopAllConsumers() bool {
	total := len(self.Consumers)
	for i := 0; i < total; i++ {
		self.Consumers[i].StopChan <- 1
		self.Consumers[i].Stop()
	}

	return true
}
