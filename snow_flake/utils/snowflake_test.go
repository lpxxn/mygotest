package utils

import (
	"encoding/binary"
	"fmt"
	"testing"
	"time"
)

func TestConstValue(t *testing.T) {
	println(workerMax)
	println(numberMax)
	println(timeShift)
	println(workerShift)

	nanoInMilli := time.Millisecond.Nanoseconds()
	println(nanoInMilli)
	pt, err := time.Parse("2006-01-02 15:04:05 -0700 UTC", "2010-11-04 01:42:54 +0000 UTC")
	if err != nil {
		panic(err)
	}
	println(pt.UnixNano())
	defaultEpoch := uint64(pt.UnixNano() / nanoInMilli)
	println(defaultEpoch)
}

func TestTwitterEpoch(t *testing.T) {
	//unixTimeStamp := "1432572732"
	//
	//unixIntValue, err := strconv.ParseInt(unixTimeStamp, 10, 64)
	//
	//if err != nil {
	//	t.Fatal(err)
	//}
	var unixIntValue int64 = 1288834974657
	timeStamp := time.Unix(unixIntValue, 0)

	fmt.Println(timeStamp)
	unitTimeInRFC3339 :=timeStamp.Format(time.RFC3339) // converts utc time to RFC3339 format
	fmt.Println("unix time stamp in unitTimeInRFC3339 format :->",unitTimeInRFC3339)

}

func TestSnowFlakeTime(t *testing.T) {
	now := time.Now().UnixNano() / nanoInMilli // 纳秒转毫秒
	var number int64 = 1
	nowVal := now-epoch
	t.Log("nowVal: ", nowVal)
	println((nowVal)<<timeShift)
	println((nowVal)<<timeShift | (0 << workerShift) | 0)
	ID1 := (nowVal)<<timeShift | (1 << workerShift) | (number)
	println(ID1)

	ID2 := (nowVal)<<timeShift | (2 << workerShift) | (number)
	println(ID2)

	// convert int64 to []byte
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, nowVal)
	b := buf[:n]
	t.Logf("b %v , b string: %s", b, string(b))
	// convert []byte to int64
	x, n := binary.Varint(b)
	fmt.Printf("x is: %v, n is: %v\n", x, n)

}
/*
>>> f'{188365481801940992:64b}'
'      1010011101001101010110100001011110110000000000000000000000
 */


func TestSnowFlake1(t *testing.T) {
	worker1, err := NewWorker(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	worker2, err := NewWorker(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	println(worker1.GetId())
	println(worker2.GetId())
}

func TestSnowFlake2(t *testing.T) {
	// 测试脚本

	// 生成节点实例
	worker, err := NewWorker(1)

	if err != nil {
		fmt.Println(err)
		return
	}

	ch := make(chan int64)
	count := 10000
	// 并发 count 个 goroutine 进行 snowflake ID 生成
	for i := 0; i < count; i++ {
		go func() {
			id := worker.GetId()
			ch <- id
		}()
	}

	defer close(ch)

	m := make(map[int64]int)
	for i := 0; i < count; i++  {
		id := <- ch
		// 如果 map 中存在为 id 的 key, 说明生成的 snowflake ID 有重复
		_, ok := m[id]
		if ok {
			t.Error("ID is not unique!\n")
			return
		}
		// 将 id 作为 key 存入 map
		m[id] = i
	}
	// 成功生成 snowflake ID
	fmt.Println("All", count, "snowflake ID Get successed!")
}