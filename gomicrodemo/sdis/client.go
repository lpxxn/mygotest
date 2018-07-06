package sdis

import (
	"fmt"
	"time"

	client "github.com/coreos/etcd/clientv3"
)

var kHeartBeatInterval = time.Second * 2
var kTTL = time.Second * 5

type Worker struct {
	kapi    client.KV
	client *client.Client
	//leaseId client.LeaseID
	key     string
	extInfo string
	active  bool
	stop    bool
}


func NewWorker(serviceName string, node string, extInfo string, endpoints []string) (*Worker, error) {
	cfg := client.Config{
		Endpoints:               endpoints,
		DialKeepAliveTimeout: time.Second * 2,
	}
	c, err := client.New(cfg)
	if err != nil {
		return nil, err
	}



	worker := &Worker{
		kapi:    client.NewKV(c),
		client: c,
		//leaseId: resp.ID,
		key:     fmt.Sprintf("%s/%s/%s", kRoot, serviceName, node),
		extInfo: extInfo,
		active:  false,
		stop:    false,
	}
	return worker, nil
}

func (w *Worker) Register() {
	w.heartbeat()
	go w.heartbeatPeriod()
}

func (w *Worker) Unregister() {
	w.stop = true
	w.active = false;
	w.kapi.Delete(w.client.Ctx(), w.key)
}

func (w *Worker) IsActive() bool {
	return w.active
}

func (w *Worker) IsStop() bool {
	return w.stop
}

func (w *Worker) heartbeatPeriod() {
	for !w.stop {
		w.heartbeat()
		time.Sleep(kHeartBeatInterval)
	}
	fmt.Println("stop is :", w.stop)
}

func (w *Worker) heartbeat() error {
	resp, err := w.client.Grant(w.client.Ctx(), 30)
	if err != nil {
		w.active = true
		return err
	}
	_, err = w.kapi.Put(w.client.Ctx(), w.key, w.extInfo, client.WithLease(resp.ID))
	w.active = err == nil
	return err
}
