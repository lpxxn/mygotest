package demo1

import (
	"encoding/json"
	"log/syslog"
	"testing"
	"time"
)

type myLog struct {
	Status  int       `json:"status"`
	Code    string    `json:"code"`
	Desc    string    `json:"desc"`
	Created time.Time `json:"created"`
}

func TestSysDial1(t *testing.T) {
	sys, err := syslog.Dial("udp", "192.168.10.144:514", syslog.LOG_LOCAL1|syslog.LOG_INFO, "myApp")
	if err != nil {
		t.Fatal(err)
	}
	l := &myLog{
		Status:  200,
		Code:    "OK",
		Desc:    "Success",
		Created: time.Now(),
	}
	bytes, _ := json.Marshal(l)
	//if err := sys.Info("@cee:" + string(bytes)); err != nil {
	if err := sys.Info(string(bytes)); err != nil {
		t.Fatal(err)
	}
	/*
	 {"@timestamp":"2021-03-01T10:46:04+08:00","host":"lideMacBook-Pro.local","severity":"info","facility":"local1","tag":"myApp[11360]:","programname":"myApp","fromhost":"192.168.10.138","fromhostip":"192.168.10.138", "status": 200, "code": "OK", "desc": "Success", "created": "2021-03-01T10:46:04.224676+08:00" }

	*/
}

func TestName(t *testing.T) {

}
func BenchmarkName(b *testing.B) {
	sys, err := syslog.Dial("udp", "192.168.10.144:514", syslog.LOG_LOCAL1|syslog.LOG_INFO, "myApp")
	if err != nil {
		b.Fatal(err)
	}
	l := &myLog{
		Status:  200,
		Code:    "OK",
		Desc:    "Success",
		Created: time.Now(),
	}
	bytes, _ := json.Marshal(l)
	idx := 0
	for i := 0; i < b.N; i++ {
		//if err := sys.Info("@cee:" + string(bytes)); err != nil {
		if err := sys.Info(string(bytes)); err != nil {
			b.Fatal(err)
			return
		}
		idx++
	}
	b.Log(idx)
}
