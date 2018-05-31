package main

import (
	"testing"
)

func Test_ReidsClient(t *testing.T) {

	_, err := RClient()
	if err != nil {
		t.Error(err)
	}
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

