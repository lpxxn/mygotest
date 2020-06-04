package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	nats "github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

// nats-streaming-server -DV --cluster_id=lp-test-cluster

func main() {
	// Connect to NATS
	opts := []nats.Option{nats.Name("NATS Streaming Example Subscriber")}
	localHost := "nats://192.168.10.208:4222"

	nc, err := nats.Connect(localHost, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	clusterID := "lp-test-cluster"
	clientID := "lp-stan-sub"
	//sc, err := stan.Connect(clusterID, clientID, stan.NatsConn(nc),
	//	stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
	//		log.Fatalf("Connection lost, reason: %v", reason)
	//	}))

	sc, err := stan.Connect(clusterID, clientID, stan.NatsConn(nc))
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, localHost)
	}
	log.Printf("Connected to %s clusterID: [%s] clientID: [%s]\n", localHost, clusterID, clientID)

	startOpt := stan.StartWithLastReceived()

	/*
		var startSeq int64 = 1

		startOpt := stan.StartAt(pb.StartPosition_NewOnly)

		if startSeq != 0 {
			startOpt = stan.StartAtSequence(startSeq)
		}
			else if deliverLast {
				startOpt = stan.StartWithLastReceived()
			} else if deliverAll && !newOnly {
				startOpt = stan.DeliverAllAvailable()
			} else if startDelta != "" {
				ago, err := time.ParseDuration(startDelta)
				if err != nil {
					sc.Close()
					log.Fatal(err)
				}
				startOpt = stan.StartAtTimeDelta(ago)

			}
	*/
	i := 0
	mcb := func(msg *stan.Msg) {
		i++
		printMsg(msg, i)
	}

	//sub, err := sc.QueueSubscribe(subj, qgroup, mcb, startOpt, stan.DurableName(durable))
	subj := "test1"
	durable := "lp-durable"
	sub, err := sc.Subscribe(subj, mcb, startOpt, stan.DurableName(durable))
	if err != nil {
		sc.Close()
		log.Fatal(err)
	}
	_ = sub
	log.Printf("Listening on [%s], clientID=[%s], durable=[%s]\n", subj, clientID, durable)

	// Wait for a SIGINT (perhaps triggered by user with CTRL-C)
	// Run cleanup when signal is received
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			fmt.Printf("\nReceived an interrupt, unsubscribing and closing connection...\n\n")
			// Do not unsubscribe a durable on exit, except if asked to.
			//if durable == "" || unsubscribe {
			//sub.Unsubscribe()
			//}
			sc.Close()
			cleanupDone <- true
		}
	}()
	<-cleanupDone

}

func printMsg(m *stan.Msg, i int) {
	log.Printf("[#%d] Received: %s\n", i, m)
}

// 这样做后需要处理断开重连
// github.com/micro/go-plugins/broker/stan
