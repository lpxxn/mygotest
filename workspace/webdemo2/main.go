package main

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"

	"github.com/mygotest/workspace/webdemo2/server"
	"github.com/mygotest/workspace/webdemo2/utils"
	"github.com/mygotest/workspace/webdemo2/utils/zaplogger"
)

func main() {
	//_, err := utils.Cluster.Set("crmweb", "value111", 0).Result()
	//fmt.Println(err)
	ipstr := utils.HostLocalIp()
	fmt.Println(ipstr)

	r := server.NewRouter()
	httpsServer := &http.Server{
		Addr:    ":9101",
		Handler: r,
	}
	//fmt.Println(httpsServer)
	zaplogger.Info("msg", zap.Any("htttps: ", httpsServer.Addr))
	go func() {
		if err := httpsServer.ListenAndServeTLS("server.crt", "server.key"); err != nil {
			//if err := http.ListenAndServeTLS(":9101", "server.crt", "server.key", r); err != nil {
			zaplogger.Panic("TLS Server Error", zap.Error(err))
		}
		fmt.Println("server run tls port: ", httpsServer.Addr)
	}()

	go func() {
		if err := http.ListenAndServe(":9100", r); err != nil {
			zaplogger.Panic("Http Server Error", zap.Error(err))
		}
		fmt.Println("server run port :", "9100")
	}()

	stopSignal := make(chan os.Signal)
	signal.Notify(stopSignal, os.Interrupt)
	quit := make(chan bool)

	go func() {
		for _ = range stopSignal {
			fmt.Println("Receive an interrup, Begin Stop.....")

			quit <- true
		}
	}()
	fmt.Println("Running Service ....")
	<-quit
	fmt.Println("Stop Server")

}
