package main

import (
	"io"
	"os"
	"os/signal"
	"sync"

	"github.com/mygotest/tcp_udp/udp_demo/demo1"
)

type ReadBuf struct {
	bytes []byte
	n     int
}

func (r *ReadBuf) GetBytes() []byte {
	return r.bytes[:r.n]
}

func (r *ReadBuf) Read(p []byte) (int, error) {
	if r.n == 0 {
		return 0, io.EOF
	}
	n := r.n
	copied := copy(p, r.bytes[:n])
	r.n -= copied
	return n, nil
}

func main() {
	hostPort := ":9876"
	server, err := demo1.NewTUDPTransport(hostPort)
	if err != nil {
		panic(err)
	}
	readBufPool := &sync.Pool{
		New: func() interface{} {
			return &ReadBuf{
				bytes: make([]byte, demo1.MaxLength),
				n:     0,
			}
		},
	}
	go func() {
		for {
			readBuf := readBufPool.Get().(*ReadBuf)
			n, err := server.Read(readBuf.bytes)
			if err == nil {
				readBuf.n = n

			}
		}
	}()
	osCh := make(chan os.Signal)
	signal.Notify(osCh, os.Interrupt)
	<-osCh
}
