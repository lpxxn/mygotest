package demo1

import (
	"bytes"
	"errors"
	"net"
	"sync/atomic"
)

//MaxLength of UDP packet
const MaxLength = 65000

var ClosedUDPErr = errors.New("UDP connection Closed")

func NewTUDPTransport(hostPort string) (*TUDPTransport, error) {
	addr, err := net.ResolveUDPAddr("udp", hostPort)
	if err != nil {
		return nil, err
	}
	conn, err := net.ListenUDP(addr.Network(), addr)
	if err != nil {
		return nil, err
	}
	return &TUDPTransport{conn: conn, addr: conn.LocalAddr()}, nil
}

type TUDPTransport struct {
	conn     *net.UDPConn
	addr     net.Addr
	writeBuf bytes.Buffer
	closed   uint32 // atomic flag
}

func (p *TUDPTransport) IsOpen() bool {
	return atomic.LoadUint32(&p.closed) == 0
}

func (p *TUDPTransport) Addr() net.Addr {
	return p.addr
}

// Read reads one UDP packet and puts it in the specified buf
func (p *TUDPTransport) Read(buf []byte) (int, error) {
	if !p.IsOpen() {
		return 0, ClosedUDPErr
	}
	return p.conn.Read(buf)
}

func (p *TUDPTransport) Write(buf []byte) (int, error) {
	if !p.IsOpen() {
		return 0, ClosedUDPErr
	}
	if len(p.writeBuf.Bytes())+len(buf) > MaxLength {
		return 0, errors.New("exceed max length")
	}
	return p.writeBuf.Write(buf)
}

func (p *TUDPTransport) Flush() error {
	if !p.IsOpen() {
		return ClosedUDPErr
	}
	_, err := p.conn.Write(p.writeBuf.Bytes())
	p.writeBuf.Reset()
	return err
}

func NewTUPPClientTransport(desHostPort string) (*TUDPTransport, error) {
	destAddr, err := net.ResolveUDPAddr("udp", desHostPort)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialUDP(destAddr.Network(), nil, destAddr)
	if err != nil {
		return nil, err
	}
	return &TUDPTransport{
		conn:     conn,
		addr:     destAddr,
		writeBuf: bytes.Buffer{},
		closed:   0,
	}, nil
}
