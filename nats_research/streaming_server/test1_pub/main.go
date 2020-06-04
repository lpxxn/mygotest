package main

import (
	"fmt"
	"log"
	"time"

	nats "github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

func main() {
	// Connect to NATS
	opts := []nats.Option{nats.Name("NATS Streaming Example Subscriber")}
	//localHost := "nats://192.168.10.208:4222"
	localHost := stan.DefaultNatsURL

	nc, err := nats.Connect(localHost, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	clusterID := "lp-test-cluster"
	clientID := "lp-stan-pub"
	sc, err := stan.Connect(clusterID, clientID, stan.NatsConn(nc),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, localHost)
	}
	log.Printf("Connected to %s clusterID: [%s] clientID: [%s]\n", localHost, clusterID, clientID)

	subj := "test1"
	sync := true
	msg := []byte(fmt.Sprintf("Hello World %d", time.Now().Unix()))
	if sync {
		err = sc.Publish(subj, msg)
		if err != nil {
			log.Fatalf("Error during publish: %v\n", err)
		}
		log.Printf("Published [%s] : '%s'\n", subj, msg)
	} else {
		//glock.Lock()
		//guid, err = sc.PublishAsync(subj, msg, acb)
		//if err != nil {
		//	log.Fatalf("Error during async publish: %v\n", err)
		//}
		//glock.Unlock()
		//if guid == "" {
		//	log.Fatal("Expected non-empty guid to be returned.")
		//}
		//log.Printf("Published [%s] : '%s' [guid: %s]\n", subj, msg, guid)
		//
		//select {
		//case <-ch:
		//	break
		//case <-time.After(5 * time.Second):
		//	log.Fatal("timeout")
		//}
	}
	fmt.Println("end push msg !")
}
