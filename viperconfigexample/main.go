package main

import (
	"github.com/mygotest/viperconfigexample/config"
	"github.com/spf13/viper"
	"log"
)

func main() {
	v := viper.New()
	v.SetConfigName("comconfig")
	v.AddConfigPath("config/")
	//v.AddConfigPath(".")

	var commonConfig config.CommonConfig

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := v.Unmarshal(&commonConfig)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	viper.SetConfigName("config")
	viper.AddConfigPath("./config" + commonConfig.Server.Addr)

	var configuration config.Configuration
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err = viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	log.Printf("database uri is %s", configuration.Database.ConnectionUri)
	log.Printf("port for this application is %d", configuration.Server.Port)
}
