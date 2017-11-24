package main

import (
	"fmt"
	"github.com/mygotest/zapdemo/utils/zaplogger"
	"os"
	"path"
	"path/filepath"
)

func main() {
	zaplogger.InitLogger()
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)

	logPath := path.Join(exPath, "crm.log")
	fmt.Println(logPath)

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)
	zaplogger.InitLogger().Error("abcde")
	zaplogger.InitLogger().Info("info:adfa")
	zaplogger.InitLogger().Panic("panic")
}
