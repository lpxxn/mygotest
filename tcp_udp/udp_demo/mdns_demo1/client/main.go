package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/miekg/dns"
	"github.com/mygotest/tcp_udp/udp_demo/mdns_demo1"
)

func main() {
	client, err := mdns_demo1.NewMServer()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	hostName, err := os.Hostname()
	if err != nil {
		fmt.Errorf("could not determine host: %v", err)
		return
	}
	hostName = fmt.Sprintf("%s.", hostName)
	//	name := fmt.Sprintf("%s.%s.%s.", sd.Instance, trimDot(sd.Service), trimDot(sd.Domain))
	// go.micro.srv.greeter-bed881eb-4191-4d51-8eb9-3ab2ad4076a1.go.micro.srv.greeter.micro.
	name := "lilili.abc.srv.test-bed881eb-4191-4d51-8eb9-3ab2ad4076a1.go.micro.srv.greeter.micro."
	q := new(dns.Msg)
	q.SetQuestion(name, dns.TypePTR)
	q.RecursionDesired = false
	var defaultTTL uint32 = 120
	srv := &dns.SRV{
		Hdr: dns.RR_Header{
			Name:   name,
			Rrtype: dns.TypeSRV,
			Class:  dns.ClassINET,
			Ttl:    defaultTTL,
		},
		Priority: 0,
		Weight:   0,
		Port:     uint16(8888),
		Target:   hostName,
	}
	txt := &dns.TXT{
		Hdr: dns.RR_Header{
			Name:   name,
			Rrtype: dns.TypeTXT,
			Class:  dns.ClassINET,
			Ttl:    defaultTTL,
		},
		Txt: []string{"abc", "hello", "world"},
	}
	q.Ns = []dns.RR{srv, txt}
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	_ = randomizer
	fmt.Println("begin send")
	//for i := 0; i < 3; i++ {
	if err := client.SendMulticast(q); err != nil {
		panic(err)
	}
	//time.Sleep(time.Duration(randomizer.Intn(250)) * time.Millisecond)
	//}

	fmt.Println("send end")
}

/*
;; opcode: QUERY, status: NOERROR, id: 31774
;; flags:; QUERY: 1, ANSWER: 0, AUTHORITY: 2, ADDITIONAL: 0

;; QUESTION SECTION:
;lilili.abc.srv.test-bed881eb-4191-4d51-8eb9-3ab2ad4076a1.go.micro.srv.greeter.micro.	IN	 PTR

;; AUTHORITY SECTION:
lilili.abc.srv.test-bed881eb-4191-4d51-8eb9-3ab2ad4076a1.go.micro.srv.greeter.micro.	120	IN	SRV	0 0 8888 li-peng-mc-macbook.local.
lilili.abc.srv.test-bed881eb-4191-4d51-8eb9-3ab2ad4076a1.go.micro.srv.greeter.micro.	120	IN	TXT	"abc" "hello" "world"
go.micro.srv.greeter-2f39c4c8-db31-4559-9c59-805278ea88fa.go.micro.srv.greeter.micro.
li-peng-mc-macbook.local.
	m, err := NewMDNSService(
		"hostname", // instance name
		"_foobar._tcp", //service
		"local.",// domain
		"testhost.", // hostname
		80, // port
		[]net.IP{net.IP([]byte{192, 168, 0, 42}), net.ParseIP("2620:0:1000:1900:b0c2:d0b2:c411:18bc")},
		[]string{"Local web server"}) // TXT

*/
