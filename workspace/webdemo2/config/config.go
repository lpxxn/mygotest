package config

import (
	"github.com/mygotest/workspace/webdemo2/utils/zaplogger"
	"github.com/spf13/viper"
)

var config *viper.Viper

func InitConfig(env string) {
	var err error
	v := viper.New()

	v.SetConfigType("yaml")
	v.SetConfigName(env)
	v.AddConfigPath("config/")
	err = v.ReadInConfig()
	if err != nil {
		zaplogger.InitLogger().Panic("read config error")
	}
	config = v
}

// https://github.com/vsouza/go-gin-boilerplate/blob/master/config/config.go
