package goredis

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis"
)
// 连接 cluster也是没有问题的，failover后一般30秒到1分钟左右就能恢复
func TestRedisClient(t *testing.T) {
	// cafeteria-c-test1.lbqctm.clustercfg.cnw1.cache.amazonaws.com.cn:6379
	client := redis.NewClient(&redis.Options{
		Addr:               "cafeteria-c-test1.lbqctm.clustercfg.cnw1.cache.amazonaws.com.cn:6379",
	})
	key1 := "a"
	for {
		time.Sleep(time.Second * 2)
		cmd := client.Set(context.Background(), key1, "a", -1)
		if cmd.Err() != nil {
			t.Log(cmd.Err())
			continue
		}
		t.Log(cmd.Val())
		strCmd := client.Get(context.Background(), key1)
		if strCmd.Err() != nil {
			t.Log(strCmd.Err())
			continue
		}
		t.Log(strCmd.Val())

	}
}

func TestRedisCluster(t *testing.T) {
	// cafeteria-c-test1.lbqctm.clustercfg.cnw1.cache.amazonaws.com.cn:6379
	client := redis.NewClusterClient(&redis.ClusterOptions{Addrs: []string{"cafeteria-c-test1.lbqctm.clustercfg.cnw1.cache.amazonaws.com.cn:6379"}})
	key1 := "a"
	for {
		time.Sleep(time.Second * 2)
		cmd := client.Set(context.Background(), key1, "a", -1)
		if cmd.Err() != nil {
			t.Log(cmd.Err())
			continue
		}
		t.Log(cmd.Val())
		strCmd := client.Get(context.Background(), key1)
		if strCmd.Err() != nil {
			t.Log(strCmd.Err())
			continue
		}
		t.Log(strCmd.Val())

	}
}

// aws master slave方式
func TestRedisGo(t *testing.T) {
	//read cafeteria-test1-ro.lbqctm.ng.0001.cnw1.cache.amazonaws.com.cn:6379
	//Primary  cafeteria-test1.lbqctm.ng.0001.cnw1.cache.amazonaws.com.cn:6379
	opt := redis.ClusterOptions{
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,

		MaxRedirects: 8,

		PoolSize:           10,
		PoolTimeout:        30 * time.Second,
		IdleTimeout:        time.Minute,
		IdleCheckFrequency: 100 * time.Millisecond,
	}
	opt.ClusterSlots = func(ctx context.Context) ([]redis.ClusterSlot, error) {
		slots := []redis.ClusterSlot{{
			Start: 0,
			End:   16383,
			Nodes: []redis.ClusterNode{{
				Addr: "cafeteria-test1.lbqctm.ng.0001.cnw1.cache.amazonaws.com.cn:6379",
			}, {
				Addr: "cafeteria-test1-ro.lbqctm.ng.0001.cnw1.cache.amazonaws.com.cn:6379",
			}},
		}}
		return slots, nil
	}
	client := redis.NewClusterClient(&opt)
	key1 := "a"
	for {
		time.Sleep(time.Second * 2)
		cmd := client.Set(context.Background(), key1, "a", -1)
		if cmd.Err() != nil {
			t.Log(cmd.Err())
			continue
		}
		t.Log(cmd.Val())
		strCmd := client.Get(context.Background(), key1)
		if strCmd.Err() != nil {
			t.Log(strCmd.Err())
			continue
		}
		t.Log(strCmd.Val())

	}
}
