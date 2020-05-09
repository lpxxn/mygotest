package demo_pipline

import (
	"testing"

	"github.com/go-redis/redis"
	"github.com/mygotest/goredis"
)

func TestPipeline1(t *testing.T) {
	client, err := goredis.RClient()
	if err != nil {
		panic(err)
	}
	rev, _ := client.Ping().Result()
	t.Log(rev)

	p1 := client.Pipeline()
	pong := p1.Ping()
	p1.Set("a:bc", "adfwe", 0)
	p1Get := p1.Get("a:bc")
	p1GetNotExist := p1.Get("not:exists:12234asdf")

	t.Log("pong err: ", pong.Err(), "pong: value", pong.Val())
	t.Log("p1Get err: ", p1Get.Err(), "p1Get: value", p1Get.Val())
	t.Log("p1GetNotExist err: ", p1GetNotExist.Err(), "p1GetNotExist: ", p1GetNotExist.Val())
	revCmd, err := p1.Exec()
	if err != redis.Nil {
		t.Error("error", err)
	}
	t.Log("after exec ----------------: ")
	t.Log("pong err: ", pong.Err(), "pong: value", pong.Val())
	t.Log("p1Get err: ", p1Get.Err(), "p1Get: value", p1Get.Val())
	t.Log("p1GetNotExist err: ", p1GetNotExist.Err(), "p1GetNotExist: ", p1GetNotExist.Val())
	t.Log("-----------")
	for _, cmd := range revCmd {
		t.Log(cmd.Name(), "  value: ", cmd.String(), " err: ", cmd.Err())
	}
}

func TestPipeline2(t *testing.T) {
	client, err := goredis.RClient()
	if err != nil {
		panic(err)
	}
	rev, _ := client.Ping().Result()
	t.Log(rev)

	revCmd, err := client.Pipelined(func(p1 redis.Pipeliner) error {
		p1.Ping()
		p1.Set("a:bc", "adfwe", 0)
		p1.Get("a:bc")
		p1.Get("not:exists:12234asdf")
		return nil
	})

	if err != redis.Nil {
		t.Error(err)
	}
	for _, cmd := range revCmd {
		t.Log(cmd.Name(), "  value: ", cmd.String(), " err: ", cmd.Err())
	}
}
