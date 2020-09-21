package main

import (
	"context"
	"testing"
)

func Test_ReidsClient(t *testing.T) {

	r, err := RClient()
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	r.Set(ctx, "m_a", "aaaa", 0).Result()
	r.Set(ctx, "m_b", "bbb", 0).Result()
	r.Set(ctx, "m_c", "ccc", 0).Result()

	rev, err := r.MGet(ctx, "m_a", "m_c", "m_zzz", "m_b", "m_d").Result()
	t.Log(err)
	t.Log(rev)

}

func Benchmark_RedisRemove(b *testing.B) {
	client, _ := RClient()
	b.StopTimer()

	b.StartTimer()
	for i := 0; i < b.N; i++ { //use b.N for looping
		go testSetPop(client)
		go testSetRemove(client, []string{"nufHwmIZz", "pXsQyzYB", "pXsQyzYB"})
		go testSetRemove(client, []string{"pXsQyzYB"})
		go testSetPop(client)
		go testSetRemove(client, []string{"CCVvJMNf", "YyHjCvE", "VfEbOm"})
		go testSetRemove(client, []string{"Mng", "BKGbkEnWqbmFE", "KSAVEtHr"})
	}

}

/*

go test -test.bench=".*"
*/
