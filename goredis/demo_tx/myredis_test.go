package txdemo

import (
	"testing"
	"github.com/go-redis/redis"
	"github.com/gin-gonic/gin/json"
	"github.com/json-iterator/go"
	"strconv"
)


const uidSayCount = "saycount:uid"
const uClient = "saycount:client"
const uid = "uid"


func Benchmark_M(b *testing.B) {
	var client = redis.NewClient(&redis.Options{
		Addr:     "192.168.3.212:6379",
		Password: "", // no password set
		DB:       3,  // use default DB
	})


	//pip := client.TxPipeline()
	//pip.HMSet("uid", map[string]interface{}{"aaa": 123})
	//
	//
	//_, err := pip.Exec()
	//if err != nil {
	//	panic(err)
	//}

	b.ResetTimer()
	//for i:= 0; i < b.N; i++ {
	//
	//	client.HMSet("uid", map[string]interface{}{"aaa": 123})
	//}

	for i:= 0; i < b.N; i++ {

		z, err := client.ZRevRangeWithScores(uidSayCount, 0, 0).Result()
		client.ZCard(uClient).Result()
		client.ZScore(uidSayCount, strconv.Itoa(i)).Result()
		pip := client.TxPipeline()
		pip.HMSet(uid, map[string]interface{}{"aaa": 123})
		z, err = pip.ZRevRangeWithScores(uidSayCount, 0, 0).Result()
		var s float64 = 0
		if len(z) > 0 {
			s = z[0].Score + 1
		}
		PipFunc(pip, uidSayCount, redis.Z{Score: s, Member: i})
		PipFunc(pip, uClient, redis.Z{Score: s, Member: i})
		pip.ZRange(uClient, 0, 500).Result()


		_, err = pip.Exec()
		if err != nil {
			panic(err)
		}

		json.Marshal(struct {
			Name string
			Age int
			Desc string
		}{Name: "李鹏一二三", Age: i, Desc: "Desc 测试"})
		json.Marshal(struct {
			Name string
			Age int
			Desc string
		}{Name: "李鹏一二三", Age: i, Desc: "Desc 测试"})
		json.Marshal(struct {
			Name string
			Age int
			Desc string
		}{Name: "李鹏一二三", Age: i, Desc: "Desc 测试"})
	}
}


func Benchmark_MW(b *testing.B) {
	var client = redis.NewClient(&redis.Options{
		Addr:     "192.168.3.212:6379",
		Password: "", // no password set
		DB:       3,  // use default DB
	})



	b.ResetTimer()
	//for i:= 0; i < b.N; i++ {
	//
	//	client.HMSet("uid", map[string]interface{}{"aaa": 123})
	//}

	for i:= 0; i < b.N; i++ {

		client.HMSet(uid, map[string]interface{}{"aaa": 123})
		z, err := client.ZRevRangeWithScores(uidSayCount, 0, 0).Result()
		//fmt.Println(z)
		var s float64 = 0
		if len(z) > 0 {
			s = z[0].Score + 1
		}
		client.ZAdd(uidSayCount, redis.Z{Score: s, Member: i})
		client.ZAdd(uClient, redis.Z{Score: s, Member: i})
		client.ZRange(uClient, 0, 500).Result()
		client.ZCard(uClient).Result()

		json.Marshal(struct {
			Name string
			Age int
			Desc string
		}{Name: "李鹏一二三", Age: i, Desc: "Desc 测试"})
		json.Marshal(struct {
			Name string
			Age int
			Desc string
		}{Name: "李鹏一二三", Age: i, Desc: "Desc 测试"})
		json.Marshal(struct {
			Name string
			Age int
			Desc string
		}{Name: "李鹏一二三", Age: i, Desc: "Desc 测试"})
		if err != nil {
			panic(err)
		}
	}
}


func Test_M(t *testing.T) {
	var client = redis.NewClient(&redis.Options{
		Addr:     "192.168.3.212:6379",
		Password: "", // no password set
		DB:       3,  // use default DB
	})


	client.HMSet("uid", map[string]interface{}{"aaa": 123})

}

func Benchmark_Mar1(b *testing.B) {
	for i := 0; i< b.N; i++ {
		v, _:=json.Marshal(struct {
			Name string
			Age int
			Desc string
		}{Name: "李鹏一二三", Age: i, Desc: "Desc 测试"})
		_ = v
	}
}

func Benchmark_Mar2(b *testing.B) {
	for i := 0; i< b.N; i++ {
		v, _:=jsoniter.Marshal(struct {
			Name string
			Age int
			Desc string
		}{Name: "李鹏一二三", Age: i, Desc: "Desc 测试"})
		_ = v
	}
}


//go test -test.bench=".*"
