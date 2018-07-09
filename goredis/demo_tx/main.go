package txdemo

import (
	"github.com/go-redis/redis"
)
var client = redis.NewClient(&redis.Options{
Addr:     "192.168.3.212:6379",
Password: "", // no password set
DB:       3,  // use default DB
})


func PipFunc(pip redis.Pipeliner, key string, z redis.Z) {
	pip.ZAdd(key, z)
}