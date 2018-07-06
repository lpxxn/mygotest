package main

import (
	"fmt"
	"time"
	"context"
	"github.com/coreos/etcd/clientv3/concurrency"
	"github.com/coreos/etcd/clientv3"
	"log"
)

// 这个是git 上有问题的。
// demo2 做了修改，要释放session才行
func main() {

	endpoints := []string{"http://192.168.3.34:2379", "http://192.168.3.18:2379", "http://192.168.3.110:2379"}
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// get a grant
	ctxc, cancelc := context.WithTimeout(context.Background(), 1*time.Second)
	lresp, err := cli.Grant(ctxc, 2)
	if err != nil {
		log.Fatal(err)
	}
	defer cancelc()
	fmt.Println("lease", lresp.ID)

	// create two separate sessions for lock competition
	s1, err := concurrency.NewSession(cli, concurrency.WithLease(lresp.ID))
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()
	m1 := concurrency.NewMutex(s1, "/xxxsx/")

	s2, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s2.Close()
	m2 := concurrency.NewMutex(s2, "/xxxsx/")

	// acquire lock for s1
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	if err := m1.Lock(ctx); err != nil {
		log.Fatal("timeout", err)
	}
	defer cancel()
	fmt.Println("acquired lock for s1")

	m2Locked := make(chan struct{})
	go func() {
		defer close(m2Locked)
		// wait until s1 is unlocked /xxxsx/
		if err := m2.Lock(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	<-m2Locked
	fmt.Println("acquired lock for s2")
}
