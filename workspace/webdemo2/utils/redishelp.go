package utils

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/mygotest/workspace/webdemo2/utils/zaplogger"
	"time"
)

//var Cluster *redis.ClusterClient = nil
var Cluster *redis.Client = nil

func init() {
	ConnectCluster()
}
func ConnectCluster() {
	//Cluster = redis.NewClusterClient(&redis.ClusterOptions{
	//	Addrs:        []string{"192.168.0.105:6379"},
	//	PoolSize:     1000,
	//	PoolTimeout:  2 * time.Minute,
	//	IdleTimeout:  10 * time.Minute,
	//	ReadTimeout:  2 * time.Minute,
	//	WriteTimeout: 1 * time.Minute,
	//	// Password: password,
	//})
	Cluster = redis.NewClient(&redis.Options{
		Addr:         "192.168.0.105:6379",
		Password:     "",
		PoolSize:     1000,
		PoolTimeout:  2 * time.Minute,
		IdleTimeout:  10 * time.Minute,
		ReadTimeout:  2 * time.Minute,
		WriteTimeout: 1 * time.Minute,
		//DB:           0,
	})
	_, err := Cluster.Ping().Result()
	if err != nil {
		fmt.Println(err)
		zaplogger.InitLogger().Panic("Redis error")
	}
}
