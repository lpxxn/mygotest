package main

import (
"github.com/Shopify/sarama"
"fmt"
	"github.com/mygotest/kafkademo/utils"
	"time"
	"os"
	"os/signal"
)
var addrs = []string{"192.168.105.240:9092"}
const testTopic = "lptest2"


func main() {

	//kafclient, _ := sarama.NewClient(addrs, nil)
	//b, _ := kafclient.Leader("lptest4", 6)
	//
	//defer b.Close()
	//ts, _ := kafclient.Topics()
	//fmt.Println(ts)

	go KafkaConsumer("0000", 0)
	go KafkaConsumer("11111", 1)
	go KafkaConsumer("2222", 2)
	go KafkaConsumer("33333", 3)


	go NewPublish()


	stopSignal := make(chan os.Signal)
	signal.Notify(stopSignal, os.Interrupt)
	quit := make(chan bool)
	go func() {
		for _ = range stopSignal {
			quit <- true
		}
	}()
	fmt.Print("Running service...")




	<-quit
	fmt.Print("StopServer")
}

var TProducer sarama.AsyncProducer

func NewPublish() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	var err error
	TProducer, err = sarama.NewAsyncProducer(addrs, config)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			//time.Sleep(time.Second * 2)
			NewPublishMsg()
		}
	}()


	for {
		select {
		case err := <-TProducer.Errors():
			if err != nil {
				fmt.Print("producer error")
				panic(err)
			}
		}
	}
}


func NewPublishMsg() {
	message := &sarama.ProducerMessage{Topic: testTopic}
	message.Value =  sarama.StringEncoder(utils.RandomStr(time.Now().UnixNano(), int64(utils.RandomInt(3, 10))))
	TProducer.Input()<- message
}


func KafkaConsumer(c string, partitionIndex int32) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	//consumer, err := sarama.NewConsumer([]string{"192.168.105.27:9092"}, nil)
	//consumer, err := sarama.NewConsumer(addrs, nil)
	//consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)

	client, err := sarama.NewClient(addrs, config)

	client.RefreshCoordinator("aaaa")
	consumer, err := sarama.NewConsumerFromClient(client)



	if err != nil {
		panic(err)
	}
	fmt.Println(consumer.Topics())
	p, err := consumer.Partitions(testTopic)
	fmt.Println("err", err, " partitions", p)


	defer func() {
		if err := consumer.Close(); err != nil {
			fmt.Println("close consumer error")
		}
	}()



	// -1默认是放到队尾
	var offset int64 = 0//sarama.OffsetNewest
	//var offset int64 = 0
	partition_consumer, err := consumer.ConsumePartition(testTopic, partitionIndex, offset)
	if err != nil {
		panic(err)
	}
	defer partition_consumer.Close()

	for {
		select {
		case msg := <- partition_consumer.Messages():
			fmt.Println("c->", c, "  " , string(msg.Value), " offset: ", msg.Offset)
		}
	}
}