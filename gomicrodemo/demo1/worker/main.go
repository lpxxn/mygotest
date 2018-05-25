package main

import (
	"log"
	"time"
	"github.com/mygotest/gomicrodemo/demo1/lib"
	"fmt"
	"github.com/coreos/etcd/pkg/transport"
)

func main(){

	serviceName := "s-test2"
	serviceInfo := lib.ServiceInfo{IP:"192.168.0.25"}


	tlsInfo := transport.TLSInfo{
		CertFile:      "/Users/li/certs/s1.pem",
		KeyFile:       "/Users/li/certs/s1-key.pem",
		TrustedCAFile: "/Users/li/certs/etcd-root-ca.pem",
	}

	tls, err := tlsInfo.ClientConfig()

	if err != nil {
		log.Fatal(err)
	}


	s, err := lib.NewService(serviceName, serviceInfo,[]string {
		"https://127.0.0.1:2379",
		"https://127.0.0.1:22379",
		"https://127.0.0.1:32379",
	}, tls)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name:%s, ip:%s\n", s.Name, s.Info.IP)


	go func() {
		time.Sleep(time.Second*50)
		s.Stop()
	}()

	s.Start()
}
