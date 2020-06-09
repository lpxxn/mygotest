package mdns_demo1

import (
	"errors"
	"fmt"
	"net"
	"sync"

	"github.com/miekg/dns"
	"golang.org/x/net/ipv4"
)

var (
	mdnsGroupIPv4 = net.ParseIP("224.0.0.251")
	//mdnsGroupIPv6 = net.ParseIP("ff02::fb")

	MdnsWildcardAddrIPv4 = &net.UDPAddr{
		IP:   net.ParseIP("224.0.0.0"),
		Port: 5353,
	}
)

type MClient struct {
	IPv4MulticastConn *net.UDPConn
	IPv4UnicastConn   *net.UDPConn
	ClosedCh          chan struct{}
	closeLock         sync.Mutex
	closed            bool
}

func NewMClient() (*MClient, error) {
	uconn4, err4 := net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4zero, Port: 0})

	if err4 != nil {
		fmt.Printf("[ERR] mdns: Failed to bind to udp port: %v", err4)
		return nil, err4
	}
	if uconn4 == nil {
		uconn4 = &net.UDPConn{}
	}

	mconn4, err4 := net.ListenUDP("udp4", MdnsWildcardAddrIPv4)
	if err4 != nil {
		fmt.Printf("[ERR] mdns: Failed to bind to udp port: %v", err4)
		return nil, err4
	}

	p1 := ipv4.NewPacketConn(mconn4)
	p1.SetMulticastLoopback(true)

	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("net.Interfaces err %#v", err)
		return nil, err4
	}
	var errCount1 int

	for _, iface := range ifaces {
		fmt.Printf("iface name : %#v \n", iface.Name)
		if err := p1.JoinGroup(&iface, &net.UDPAddr{IP: mdnsGroupIPv4}); err != nil {
			errCount1++
		}
	}

	if len(ifaces) == errCount1 {
		return nil, errors.New("failed to join multicast group on all interfaces!")
	}
	return &MClient{
		IPv4MulticastConn: mconn4,
		IPv4UnicastConn:   uconn4,
		ClosedCh:          make(chan struct{}),
	}, nil
}

// Recv is used to receive until we get a shutdown
func (c *MClient) Recv(l *net.UDPConn, msgCh chan *dns.Msg) {
	if l == nil {
		return
	}
	buf := make([]byte, 65536)
	for {
		c.closeLock.Lock()
		if c.closed {
			c.closeLock.Unlock()
			return
		}
		c.closeLock.Unlock()
		n, err := l.Read(buf)
		if err != nil {
			continue
		}
		msg := new(dns.Msg)
		if err := msg.Unpack(buf[:n]); err != nil {
			continue
		}
		select {
		case msgCh <- msg:
		case <-c.ClosedCh:
			return
		}
	}
}

func (c *MClient) SetInterface(iface *net.Interface, loopback bool) error {
	p := ipv4.NewPacketConn(c.IPv4UnicastConn)
	if err := p.JoinGroup(iface, &net.UDPAddr{IP: mdnsGroupIPv4}); err != nil {
		return err
	}
	p = ipv4.NewPacketConn(c.IPv4MulticastConn)
	if err := p.JoinGroup(iface, &net.UDPAddr{IP: mdnsGroupIPv4}); err != nil {
		return err
	}

	if loopback {
		p.SetMulticastLoopback(true)
	}

	return nil
}

func (c *MClient) Close() error {
	c.closeLock.Lock()
	defer c.closeLock.Unlock()

	if c.closed {
		return nil
	}
	c.closed = true

	close(c.ClosedCh)

	if c.IPv4UnicastConn != nil {
		c.IPv4UnicastConn.Close()
	}
	if c.IPv4MulticastConn != nil {
		c.IPv4MulticastConn.Close()
	}

	return nil
}

type MServer struct {
	IPv4MulticastConn *net.UDPConn
	ClosedCh          chan struct{}
	closeLock         sync.Mutex
	closed            bool
}

func NewMServer() (*MServer, error) {
	ipv4List, _ := net.ListenUDP("udp4", MdnsWildcardAddrIPv4)
	if ipv4List == nil {
		return nil, fmt.Errorf("[ERR] mdns: Failed to bind to any udp port!")
	}

	if ipv4List == nil {
		ipv4List = &net.UDPConn{}
	}

	// Join multicast groups to receive announcements
	p1 := ipv4.NewPacketConn(ipv4List)
	p1.SetMulticastLoopback(true)

	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	errCount1 := 0
	for _, iface := range ifaces {
		if err := p1.JoinGroup(&iface, &net.UDPAddr{IP: mdnsGroupIPv4}); err != nil {
			errCount1++
		}
	}
	if len(ifaces) == errCount1 {
		return nil, fmt.Errorf("Failed to join multicast group on all interfaces!")
	}
	s := &MServer{

		IPv4MulticastConn: ipv4List,
		ClosedCh:          make(chan struct{}),
		closeLock:         sync.Mutex{},
		closed:            false,
	}
	//go s.RevN(s.IPv4MulticastConn)
	return s, nil
}

var ipv4Addr = &net.UDPAddr{
	IP:   mdnsGroupIPv4,
	Port: 5353,
}

func (s *MServer) SendMulticast(msg *dns.Msg) error {
	buf, err := msg.Pack()
	if err != nil {
		return err
	}
	if s.IPv4MulticastConn != nil {
		if _, err := s.IPv4MulticastConn.WriteToUDP(buf, ipv4Addr); err != nil {
			fmt.Println("write to udp err: ", err.Error())
		}
	}
	return nil
}

func (s *MServer) Close() error {
	s.closeLock.Lock()
	defer s.closeLock.Unlock()

	if s.closed {
		return nil
	}
	s.closed = true

	close(s.ClosedCh)

	if s.IPv4MulticastConn != nil {
		s.IPv4MulticastConn.Close()
	}

	return nil
}

func (s *MServer) RevN(l *net.UDPConn) {
	if l == nil {
		return
	}
	buf := make([]byte, 65536)
	for {
		s.closeLock.Lock()
		if s.closed {
			s.closeLock.Unlock()
			return
		}
		s.closeLock.Unlock()
		fmt.Println("begin read")
		_, err := l.Read(buf)
		if err != nil {
			fmt.Println("err: ", err.Error())
			continue
		}
		fmt.Println(string(buf))
	}
}
