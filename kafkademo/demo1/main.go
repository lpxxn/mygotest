package main

import (
	"github.com/Shopify/sarama"
	"fmt"
)


var addrs = []string{"192.168.105.240:9092"}
const testTopic = "lptest2"

func main() {
	//consumer, err := sarama.NewConsumer([]string{"192.168.105.27:9092"}, nil)
	consumer, err := sarama.NewConsumer(addrs, nil)
	//consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)

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
	offset := sarama.OffsetNewest
	//var offset int64 = 0
	partition_consumer, err := consumer.ConsumePartition(testTopic, 1, offset)

	if err != nil {
		panic(err)
	}
	defer partition_consumer.Close()

	for {
		select {
		case msg := <- partition_consumer.Messages():
			fmt.Println("partions 1", string(msg.Value))
		}
	}

}
