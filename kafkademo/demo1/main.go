package main

import (
	"github.com/Shopify/sarama"
	"fmt"
)

func main() {
	//consumer, err := sarama.NewConsumer([]string{"192.168.105.27:9092"}, nil)
	consumer, err := sarama.NewConsumer([]string{"192.168.105.225:9092"}, nil)
	//consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)

	if err != nil {
		panic(err)
	}
	fmt.Println(consumer.Topics())
	p, err := consumer.Partitions("test")
	fmt.Println("err", err, " partitions", p)


	defer func() {
		if err := consumer.Close(); err != nil {
			fmt.Println("close consumer error")
		}
	}()



	// -1默认是放到队尾
	//offset := sarama.OffsetNewest
	var offset int64 = 0
	partition_consumer, err := consumer.ConsumePartition("test", 0, offset)
	if err != nil {
		panic(err)
	}
	defer partition_consumer.Close()

	for {
		select {
		case msg := <- partition_consumer.Messages():
			fmt.Println(string(msg.Value))
		}
	}

}
