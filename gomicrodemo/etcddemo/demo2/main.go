package main

import (
	"fmt"
	"time"
	"context"
	"github.com/coreos/etcd/clientv3/concurrency"
	"github.com/coreos/etcd/clientv3"
	"log"
)

func main() {

	endpoints := []string{"http://192.168.3.34:2379", "http://192.168.3.18:2379", "http://192.168.3.110:2379"}
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// get a grant
	ctxc, cancelc := context.WithTimeout(context.Background(), 3*time.Second)
	lresp, err := cli.Grant(ctxc, 2)
	if err != nil {
		log.Fatal(err)
	}
	defer cancelc()
	fmt.Println("lease", lresp.ID)

	
	// create two separate sessions for lock competition
	s1, err := concurrency.NewSession(cli, concurrency.WithTTL(10))
	if err != nil {
		log.Fatal(err)
	}
	m1 := concurrency.NewMutex(s1, "/xxxsx/")

	s2, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	m2 := concurrency.NewMutex(s2, "/xxxsx/")

	// acquire lock for s1
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	if err := m1.Lock(ctx); err != nil {
		fmt.Println("m1 lock error")
		log.Fatal("timeout", err)
	}
	defer cancel()
	fmt.Println("acquired lock for s1")

	m2Locked := make(chan struct{})
	go func() {
		defer close(m2Locked)
		// wait until s1 is unlocked /xxxsx/
		if err := m2.Lock(context.TODO()); err != nil {
			fmt.Println("m2 lock error")
			log.Fatal(err)
		}
		fmt.Println("release m2 session")
		s2.Close()
	}()

	m3Locked := make(chan struct{})
	go func() {
		defer close(m3Locked)
		s3, err := concurrency.NewSession(cli)
		if err != nil {
			log.Fatal(err)
		}
		m3 := concurrency.NewMutex(s3, "/xxxsx/")

		if err := m3.Lock(context.TODO()); err != nil {
			fmt.Println("m3 lock error")
			log.Fatal(err)
		}
		s3.Close()
		fmt.Println("m3 acquired lock")
	}()

	s1.Close()



	<-m2Locked
	fmt.Println("acquired lock for s2")

	<-m3Locked

}
