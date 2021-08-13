package main

import (
	"context"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestCancel1(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8090/hello", nil)
	if err != nil {
		t.Fatal(err)
	}
	//client := http.Client{Timeout: 2 * time.Second }
	client := http.Client{Timeout: 200 * time.Second}
	ctx, cancel := context.WithCancel(req.Context())
	req = req.WithContext(ctx)
	time.AfterFunc(time.Second * 2, func() {
		cancel()
	})
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", b)
}
// https://www.sohamkamani.com/golang/context-cancellation-and-values/

/*
net/http/server.go

func (cr *connReader) startBackgroundRead() {
	cr.lock()
	defer cr.unlock()
	if cr.inRead {
		panic("invalid concurrent Body.Read call")
	}
	if cr.hasByte {
		return
	}
	cr.inRead = true
	cr.conn.rwc.SetReadDeadline(time.Time{})
	go cr.backgroundRead()
}

// It may be called from multiple goroutines.
func (cr *connReader) handleReadError(_ error) {
	cr.conn.cancelCtx()
	cr.closeNotify()
}

 */