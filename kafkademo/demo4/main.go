package main

import (
	"github.com/Shopify/sarama"
	"time"
	"github.com/mygotest/kafkademo/utils"
)

var addrs = []string{"192.168.105.240:9092"}
const testTopic = "lptest2"
func main() {

	var err error
	TProducer, err := sarama.NewAsyncProducer(addrs, nil)
	if err != nil {
		panic(err)
	}

	for {
		time.Sleep(time.Second * 2)
		message := &sarama.ProducerMessage{Topic: testTopic}
		message.Value =  sarama.StringEncoder(utils.RandomStr(time.Now().UnixNano(), int64(utils.RandomInt(3, 10))))
		TProducer.Input()<- message
	}
}