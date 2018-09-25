package main

import (
	"log"
	"time"
	"github.com/mygotest/gomicrodemo/demo1/lib"
	"fmt"
	"github.com/coreos/etcd/pkg/transport"
)

func main() {

	tlsInfo := transport.TLSInfo{
		CertFile:      "/Users/li/certs/s1.pem",
		KeyFile:       "/Users/li/certs/s1-key.pem",
		TrustedCAFile: "/Users/li/certs/etcd-root-ca.pem",
	}

	tls, err := tlsInfo.ClientConfig()

	if err != nil {
		log.Fatal(err)
	}



	m, err := lib.NewMaster([]string{
		"https://127.0.0.1:2379",
		"https://127.0.0.1:22379",
		"https://127.0.0.1:32379",
	}, "services/", tls)

	if err != nil {
		log.Fatal(err)
	}

	for {
		for k, v := range  m.Nodes {
			fmt.Printf("node:%s, ip=%s\n", k, v.Info.IP)
		}
		fmt.Printf("nodes num = %d\n",len(m.Nodes))
		time.Sleep(time.Second * 5)
	}

}


/**
https://studygolang.com/articles/12165



 */