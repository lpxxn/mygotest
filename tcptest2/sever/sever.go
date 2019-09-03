package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("Launching server ....")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", "127.0.0.1:9091")

	ctx := defaultServerContext()
	tcpSrv := &tcpServer{ctx: ctx}
	TcpServer(ln, tcpSrv)

	out := make(chan os.Signal)
	signal.Notify(out, os.Interrupt)
	<-out
	fmt.Print("end connection")
	//// accept connection on port
	//conn, _ := ln.Accept()
	//// run loop forever (or util ctrl-c)
	//for {
	//	// will listen for message to process ending in newline(\n)
	//	message, _ := bufio.NewReader(conn).ReadString('\n')
	//	// output message received
	//	fmt.Println("Message Received: ", string(message))
	//
	//	// sample process for string received
	//	newmessage := strings.ToUpper(message)
	//	// send new string back to client
	//	conn.Write([]byte(newmessage + "\n"))
	//}
}

type TCPHandler interface {
	Handle(net.Conn)
}

func TcpServer(listener net.Listener, handler TCPHandler) error {
	log.Printf("TCP: listening on %s", listener.Addr())
	var revErr error = nil
	for {
		clientConn, err := listener.Accept()
		if err != nil {
			if nErr, ok := err.(net.Error); ok && nErr.Temporary() {
				log.Printf("temporary Accept() failure : %s", err)
				runtime.Gosched()
				continue
			}
			log.Printf("listener.Accept() error %#v", err)
			revErr = err
			break
		}

		go handler.Handle(clientConn)
	}

	log.Printf("TCP: closing %s", listener.Addr())
	return revErr
}

// ---------
type Message struct {
	ID        int64
	Body      []byte
	Timestamp int64
	Attempts  uint16

	ClientID int64
}

// ---------
const defaultBufferSize = 16 * 1024

type Client struct {
	ctx      *ServerContext
	exitChan chan int

	net.Conn
	LastUpdate        int64
	ID                int64
	Address           string
	HostName          string
	TcpPort           int
	HeartbeatInterval time.Duration

	writeLock sync.RWMutex
	Reader    *bufio.Reader
	Writer    *bufio.Writer

	OutputBufferSize    int
	OutputBufferTimeout time.Duration

	memoryMsgChan chan *Message
}

func newClient(id int64, conn net.Conn, ctx *ServerContext) *Client {
	hostName, port, _ := net.SplitHostPort(conn.RemoteAddr().String())
	iPort, _ := strconv.Atoi(port)

	return &Client{
		ctx:        ctx,
		Conn:       conn,
		ID:         id,
		LastUpdate: time.Now().Unix(),
		Address:    conn.RemoteAddr().String(),
		HostName:   hostName,
		TcpPort:    iPort,
		exitChan:   make(chan int),

		Reader:              bufio.NewReaderSize(conn, defaultBufferSize),
		Writer:              bufio.NewWriterSize(conn, defaultBufferSize),
		OutputBufferSize:    defaultBufferSize,
		OutputBufferTimeout: 250 * time.Millisecond,

		memoryMsgChan: make(chan *Message, 500),
	}
}

func (c *Client) String() string {
	return c.RemoteAddr().String()
}

func (c *Client) Flush() error {
	var zeroTime time.Time
	if c.HeartbeatInterval > 0 {
		c.SetWriteDeadline(time.Now().Add(c.HeartbeatInterval))
	} else {
		c.SetWriteDeadline(zeroTime)
	}
	return c.Writer.Flush()
}

// ------------
type ServerInfo struct {
	sync.RWMutex

	clientLock       sync.RWMutex
	clientIDSequence int64
	clients          map[int64]*Client
	exitChan         chan int
}

func (s *ServerInfo) AddClient(clientID int64, client *Client) {
	s.clientLock.Lock()
	s.clients[clientID] = client
	s.clientLock.Unlock()
}

func (s *ServerInfo) RemoveClient(clientID int64) {
	s.clientLock.Lock()
	if _, ok := s.clients[clientID]; ok {
		delete(s.clients, clientID)
	}
	s.clientLock.Unlock()
}

type ServerContext struct {
	serverInfo *ServerInfo
}

func defaultServerContext() *ServerContext {
	return &ServerContext{serverInfo: &ServerInfo{clientIDSequence: 0, clients: map[int64]*Client{}, exitChan: make(chan int)}}
}

// ------------ tcpServer
type tcpServer struct {
	ctx *ServerContext
}

// 处理具体的一个connection
func (p *tcpServer) Handle(clientConn net.Conn) {
	log.Printf("TCP: new client(%s)", clientConn.RemoteAddr())
	prot := &protocolV1{ctx: p.ctx}
	if err := prot.ConnectionLoop(clientConn); err != nil {
		log.Printf("client err: %#v", err)
	}
}

type protocolV1 struct {
	ctx *ServerContext
}

func (p *protocolV1) ConnectionLoop(conn net.Conn) error {
	clientID := atomic.AddInt64(&p.ctx.serverInfo.clientIDSequence, 1)
	client := newClient(clientID, conn, p.ctx)
	p.ctx.serverInfo.AddClient(client.ID, client)
	readyPumpMsg := make(chan bool)
	go p.messagePump(client, readyPumpMsg)
	<-readyPumpMsg

	var zeroTime time.Time
	var err error
	for {
		if client.HeartbeatInterval > 0 {
			client.SetReadDeadline(time.Now().Add(client.HeartbeatInterval * 2))
		} else {
			client.SetReadDeadline(zeroTime)
		}
		var allData []byte
		allData, err = client.Reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
		client.memoryMsgChan <- &Message{ID: 1, Body: allData}
	}
	log.Printf("protocolV1 ConnectionLoop exit loop client %s", client.String())
	conn.Close()
	close(client.exitChan)
	p.ctx.serverInfo.RemoveClient(client.ID)
	return err
}

func (p *protocolV1) messagePump(client *Client, ready chan bool) {
	outputBufferTicker := time.NewTicker(client.OutputBufferTimeout)
	var flusherChan <-chan time.Time

	msgChan := client.memoryMsgChan
	flushed := true

	var err error

	// ready
	close(ready)

	for {
		if flushed {
			flusherChan = nil
		} else {
			flusherChan = outputBufferTicker.C
		}
		select {
		case <-flusherChan:
			client.writeLock.Lock()
			err = client.Flush()
			client.writeLock.Unlock()
			if err != nil {
				goto exit
			}
			flushed = true

		case msg := <-msgChan:
			msg.Attempts++
			_, err = p.SendMsg(client, msg)
			if err != nil {
				goto exit
			}
			flushed = false

		case <-client.exitChan:
			goto exit
		}
	}
exit:
	log.Printf("client %s exiting connection loop ", client.String())
	outputBufferTicker.Stop()
	outputBufferTicker = nil
	if err != nil {
		fmt.Printf("protocolV1 messagePump error : %#v", err)
	}
}

func (p *protocolV1) SendMsg(client *Client, msg *Message) (int, error) {
	log.Printf("protocolV1 writig msg(%s) to client(%s)", string(msg.Body), client.String())
	var buf = &bytes.Buffer{}
	buf.WriteString("server side received msg: ")
	buf.Write(msg.Body)

	client.writeLock.Lock()
	err := binary.Write(client.Writer, binary.BigEndian, int32(len(buf.Bytes())))
	if err != nil {
		return 0, err
	}
	n, err := client.Writer.Write(buf.Bytes())
	if err != nil {
		return 0, err
	}
	client.writeLock.Unlock()

	return n + 4, nil
}
