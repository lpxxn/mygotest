package main

import (
	"context"
	"go.uber.org/fx"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	app := fx.New(fx.Provide(NewLogger, NewHandler, NewMux),
		fx.Invoke(Register))
	startCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		panic(err)
	}

	//http.Get("http://localhost:8001")
	<-app.Done()
	stopCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}
	// curl http://127.0.0.1:8001/hello
}

func NewLogger() *log.Logger {
	logger := log.New(os.Stdout, "", 0)
	logger.Print("new logger...")
	return logger
}

func NewHandler(l *log.Logger) (http.Handler, error) {
	l.Print("NewHandler...")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
		l.Print("get a request...")
	}), nil
}

func Register(mux *http.ServeMux, h http.Handler, l *log.Logger) {
	l.Print("Register...")
	mux.Handle("/", h)
	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi there"))
	})
}

func NewMux(lc fx.Lifecycle, l *log.Logger) *http.ServeMux {
	l.Print("NewMux...")
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8001",
		Handler: mux,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			l.Print("Starting HTTP server...")
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			l.Print("Stopping HTTP server.")
			return server.Shutdown(ctx)
		},
	})
	return mux
}
