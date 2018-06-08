package etcd_service_discovery

import (
	"fmt"
	client "github.com/coreos/etcd/clientv3"
	"log"
	"strings"
	"sync"
	"time"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

var kRoot = "imservice"

type Service struct {
	sync.RWMutex
	kapi   client.KV
	client *client.Client
	key    string
	nodes  map[string]string
	active bool
}

func NewService(serviceName string, endpoints []string) (*Service, error) {
	cfg := client.Config{
		Endpoints:               endpoints,
		DialKeepAliveTimeout: time.Second * 2,
	}
	c, err := client.New(cfg)
	if err != nil {
		return nil, err
	}
	master := &Service{
		kapi:   client.NewKV(c),
		client: c,
		key:    fmt.Sprintf("%s/%s/", kRoot, serviceName),
		nodes:  make(map[string]string),
		active: true,
	}
	master.fetch()

	/// `fetch` Timer may work well too?
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

func (m *Service) addNode(node, extInfo string) {
	m.Lock()
	defer m.Unlock()
	node = strings.TrimLeft(node, m.key)
	m.nodes[node] = extInfo
}

func (m *Service) addNodebyte(node, extInfo []byte) {
	m.Lock()
	defer m.Unlock()
	nodestr := strings.TrimLeft(string(node), m.key)
	m.nodes[nodestr] = string(extInfo)
}

func (m *Service) delNode(node string) {
	m.Lock()
	defer m.Unlock()
	node = strings.TrimLeft(node, m.key)
	delete(m.nodes, node)
}


func (m *Service) watch() {
	rch := m.client.Watch(m.client.Ctx(), m.key, client.WithPrefix())
	for resp := range rch {

		m.active = true
		//log.Println("loop active ->", m.active)
		for _, ev := range resp.Events {
			switch ev.Type {
			case mvccpb.PUT:
				m.addNodebyte(ev.Kv.Key, ev.Kv.Value)
				break
			case mvccpb.DELETE:
				m.delNode(string(ev.Kv.Key))
				break
			default:
				log.Println("watchme!!!", "resp ->", resp)
			}
		}
	}
}

func (m *Service) fetch() error {
	resp, err := m.kapi.Get(m.client.Ctx(), m.key)
	if err != nil {
		return err
	}
	if resp.Count > 0 {
		for _, v := range resp.Kvs {
			m.addNodebyte(v.Key, v.Value)
		}
	}
	return err
}

