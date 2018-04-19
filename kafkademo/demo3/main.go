package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	cluster "github.com/bsm/sarama-cluster"
)

var addrs = []string{"192.168.105.240:9092"}

func main() {

	go ConsumerTest("11111")
	go ConsumerTest("22222")

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

var testtop = "lptest2"
func ConsumerTest(c string) {
	// init (custom) config, enable errors and notifications
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true


	// init consumer
	brokers := addrs
	//topics := []string{"test", "lptest", "lptest2"}
	topics := []string{testtop}

	/*
	client, _ := cluster.NewClient(brokers, config)
	offset_manager, _ := sarama.NewOffsetManagerFromClient("my_group", client)
	defer offset_manager.Close()
	of_o, err := offset_manager.ManagePartition(testtop, 0)
	of_o.ResetOffset(0, "")

	of_o1, err := offset_manager.ManagePartition(testtop, 1)
	of_o1.ResetOffset(0, "")

	//
	of_o2, err := offset_manager.ManagePartition(testtop, 2)
	of_o2.ResetOffset(1, "")

	of_o3, err := offset_manager.ManagePartition(testtop, 3)
	of_o3.ResetOffset(1, "")

	consumer, err := cluster.NewConsumerFromClient(client, "my_group", topics)
	*/
	consumer, err := cluster.NewConsumer(brokers, "my-consumer-group", topics, config)

	consumer.SetTopicPartitionOffset = func() map[string]map[int32]int64 {
		return map[string]map[int32]int64{
			topics[0]: map[int32]int64{0: -1, 1: 500, 2: 30},
		}
	}


	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	// trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// consume errors
	go func() {
		for err := range consumer.Errors() {
			log.Printf("Error: %s\n", err.Error())
		}
	}()

	// consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			log.Printf("Rebalanced: %+v\n", ntf)
		}
	}()



	//consumer.ResetOffset(&sarama.ConsumerMessage{Topic: testtop, Offset: -1, Partition: 0}, "")
	//consumer.ResetOffset(&sarama.ConsumerMessage{Topic: testtop, Offset: -1, Partition: 1}, "")
	//consumer.ResetPartitionOffset(testtop, 0, -1, "")
	//consumer.ResetPartitionOffset(testtop, 2, -1, "")
	//consumer.ResetPartitionOffset(testtop, 1, -1, "0")
 	//consumer.CommitOffsets()
	// consume messages, watch signals
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				//fmt.Println("consomer: ", c, " offset :", msg.Offset)
				fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\t%s Consumer: %s Offset: %d\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value, c, msg.Offset)
				consumer.MarkOffset(msg, "")	// mark message as processed
			}
		case <-signals:
			return
		}
	}
}