package sdis

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/coreos/etcd/clientv3"
)

var kRoot = "gateway"

type SNode struct {
	OpType mvccpb.Event_EventType
	Node string
	Value string
}

type Service struct {
	sync.RWMutex
	kapi   clientv3.KV
	client *clientv3.Client
	key    string
	nodes  map[string]string
	active bool
	nodeChan chan *SNode
}

func NewService(serviceName string, endpoints []string, ch chan *SNode) (*Service, error) {
	cfg := clientv3.Config{
		Endpoints:               endpoints,
		DialKeepAliveTimeout: time.Second * 2,
	}
	c, err := clientv3.New(cfg)
	if err != nil {
		return nil, err
	}
	master := &Service{
		kapi:   clientv3.NewKV(c),
		client: c,
		key:    fmt.Sprintf("%s/%s/", kRoot, serviceName),
		nodes:  make(map[string]string),
		active: true,
		nodeChan: ch,
	}
	master.fetch()

	go master.watch()

	return master, err
}



func (m *Service) GetNodesStrictly() map[string]string {
	//log.Println("strictly active ->", m.active)
	if !m.active {
		return nil
	}
	return m.GetNodes()
}

func (m *Service) GetNodes() map[string]string {
	m.RLock()
	defer m.RUnlock()
	return m.nodes
}

func (m *Service) addNode(eventType mvccpb.Event_EventType, node, extInfo string) {
	m.Lock()
	defer m.Unlock()
	node = strings.TrimLeft(node, m.key)

	if m.nodeChan != nil {
		m.nodeChan <- &SNode{OpType:eventType, Node:node, Value: extInfo}
	}

	m.nodes[node] = extInfo

}

func (m *Service) delNode(eventType mvccpb.Event_EventType, node string) {
	m.Lock()
	defer m.Unlock()
	node = strings.TrimLeft(node, m.key)
	delete(m.nodes, node)
	if m.nodeChan != nil {
		m.nodeChan <- &SNode{OpType:eventType, Node:node}
	}
}


func (m *Service) watch() {
	rch := m.client.Watch(m.client.Ctx(), m.key, clientv3.WithPrefix())
	for resp := range rch {

		m.active = true
		//log.Println("loop active ->", m.active)
		for _, ev := range resp.Events {
			switch ev.Type {
			case mvccpb.PUT:
				m.addNode(mvccpb.PUT, string(ev.Kv.Key), string(ev.Kv.Value))
				break
			case mvccpb.DELETE:
				m.delNode(mvccpb.DELETE, string(ev.Kv.Key))
				break
			default:
				log.Println("watchme!!!", "resp ->", resp)
			}
		}
	}
}

func (m *Service) fetch() error {
	resp, err := m.kapi.Get(m.client.Ctx(), m.key, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	if resp.Count > 0 {
		for _, v := range resp.Kvs {
			m.addNode(mvccpb.PUT, string(v.Key), string(v.Value))
		}
	}
	return err
}

