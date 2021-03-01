package demo1

import (
	"log/syslog"
	"testing"
	"time"

	"github.com/sirupsen/logrus"

	lSyslog "github.com/sirupsen/logrus/hooks/syslog"
)

func TestLogrus1(t *testing.T) {
	// Use the Airbrake hook to report errors that have Error severity or above to
	// an exception tracker. You can create custom hooks, see the Hooks section.
	log := logrus.New()
	hook, err := lSyslog.NewSyslogHook("udp", "192.168.10.144:514", syslog.LOG_LOCAL1|syslog.LOG_INFO, "haha")
	if err != nil {
		log.Error("Unable to connect to local syslog daemon")
	} else {
		log.AddHook(hook)
	}
	l := &myLog{
		Status:  200,
		Code:    "OK",
		Desc:    "Success",
		Created: time.Now(),
	}
	_ = l
	//bytes, _ := json.Marshal(l)
	log.WithField("status", "my_OK").Info("absafsef")
	//log.Info(l)
}
