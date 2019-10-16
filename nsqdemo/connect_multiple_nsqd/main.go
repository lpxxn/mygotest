package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/nsqio/go-nsq"
)

func main() {
	//adds := []string{"127.0.0.1:7000", "127.0.0.1:8000"}
	adds := []string{"127.0.0.1:4150"}
	config := nsq.NewConfig()
	config.MaxInFlight = 1000
	config.MaxBackoffDuration = 5 * time.Second
	config.DialTimeout = 10 * time.Second

	topicName := "testTopic1"
	c, _ := nsq.NewConsumer(topicName, "ch1", config)
	testHandler := &MyTestHandler{consumer: c}

	c.AddHandler(testHandler)
	if err := c.ConnectToNSQDs(adds); err != nil {
		panic(err)
	}
	stats := c.Stats()
	if stats.Connections == 0 {
		panic("stats report 0 connections (should be > 0)")
	}
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	fmt.Println("server is running....")
	<-stop
}

type MyTestHandler struct {
	consumer *nsq.Consumer
}

type A struct {
	ID int `json:"id"`
}

func (m MyTestHandler) HandleMessage(message *nsq.Message) error {
	fmt.Println(string(message.Body))
	rev := &A{}
	if err := json.Unmarshal(message.Body, rev); err != nil {
		return err
	}
	fmt.Println("returned message A id", rev.ID)
	if rev.ID == 1 || rev.ID == 3 {
		fmt.Println("err")
		return errors.New("err")
	}
	return nil
}
