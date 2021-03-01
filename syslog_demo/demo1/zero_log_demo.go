package demo1

import (
	"io"
	"testing"
)

func TestZeroLog1(t *testing.T) {

	//sw := &syslogTestWriter{}
	//log := zerolog.New(zerolog.SyslogLevelWriter(sw))
	//log.Trace().Msg("trace")
	//log.Debug().Msg("debug")
	//log.Info().Msg("info")
	//log.Warn().Msg("warn")
	//log.Error().Msg("error")
	//log.Log().Msg("nolevel")
	//want := []syslogEvent{
	//	{"Debug", `{"level":"debug","message":"debug"}` + "\n"},
	//	{"Info", `{"level":"info","message":"info"}` + "\n"},
	//	{"Warning", `{"level":"warn","message":"warn"}` + "\n"},
	//	{"Err", `{"level":"error","message":"error"}` + "\n"},
	//	{"Info", `{"message":"nolevel"}` + "\n"},
	//}
	//if got := sw.events; !reflect.DeepEqual(got, want) {
	//	t.Errorf("Invalid syslog message routing: want %v, got %v", want, got)
	//}
}

type syslogEvent struct {
	level string
	msg   string
}
type syslogTestWriter struct {
	events []syslogEvent
	writer io.Writer
}

func (w *syslogTestWriter) Write(p []byte) (int, error) {
	return w.Write(p)
}
func (w *syslogTestWriter) Trace(m string) error {
	w.events = append(w.events, syslogEvent{"Trace", m})
	return nil
}
func (w *syslogTestWriter) Debug(m string) error {
	w.events = append(w.events, syslogEvent{"Debug", m})
	return nil
}
func (w *syslogTestWriter) Info(m string) error {
	w.events = append(w.events, syslogEvent{"Info", m})
	return nil
}
func (w *syslogTestWriter) Warning(m string) error {
	w.events = append(w.events, syslogEvent{"Warning", m})
	return nil
}
func (w *syslogTestWriter) Err(m string) error {
	w.events = append(w.events, syslogEvent{"Err", m})
	return nil
}
func (w *syslogTestWriter) Emerg(m string) error {
	w.events = append(w.events, syslogEvent{"Emerg", m})
	return nil
}
func (w *syslogTestWriter) Crit(m string) error {
	w.events = append(w.events, syslogEvent{"Crit", m})
	return nil
}
