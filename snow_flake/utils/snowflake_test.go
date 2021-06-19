package utils

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"testing"
	"time"
	"unicode"
)

func TestConstValue(t *testing.T) {
	t.Log(workerMax)
	t.Log(numberMax)
	t.Log(timeShift)
	t.Log(workerShift)

	const workMax2 int64 = 1<<workerBits - 1
	t.Log(workMax2)
	t.Log(time.Duration(-1))
	time.Sleep(time.Duration(time.Second * -10))
	nanoInMilli := time.Millisecond.Nanoseconds()
	t.Log(nanoInMilli)
	pt, err := time.Parse("2006-01-02 15:04:05 -0700 UTC", "2010-11-04 01:42:54 +0000 UTC")
	if err != nil {
		panic(err)
	}
	t.Log(pt.UnixNano())
	defaultEpoch := uint64(pt.UnixNano() / nanoInMilli)
	t.Log(defaultEpoch)
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
	timeStamp := time.Unix(0, unixIntValue*nanoInMilli)

	t.Log(timeStamp)
	unitTimeInRFC3339 := timeStamp.Format(time.RFC3339) // converts utc time to RFC3339 format
	t.Log("unix time stamp in unitTimeInRFC3339 format :->", unitTimeInRFC3339)
	var t41 int64 = 2199023255552
	//var t42 int64 = 4398046511104
	unixIntValue = t41
	timeStamp = time.Unix(0, unixIntValue*nanoInMilli)

	t.Log(timeStamp)
	unitTimeInRFC3339 = timeStamp.Format(time.RFC3339) // converts utc time to RFC3339 format
	t.Log("unix time stamp in unitTimeInRFC3339 format :->", unitTimeInRFC3339)

	unixIntValue -= 1288834974657
	timeStamp = time.Unix(0, unixIntValue*nanoInMilli)

	t.Log(timeStamp)
	unitTimeInRFC3339 = timeStamp.Format(time.RFC3339) // converts utc time to RFC3339 format
	t.Log("unix time stamp in unitTimeInRFC3339 format :->", unitTimeInRFC3339)

}

func TestSnowFlakeTime(t *testing.T) {
	t.Log(time.Millisecond.Milliseconds(), "  ", time.Millisecond.Microseconds(), " ", time.Millisecond.Nanoseconds())
	curtTime := time.Now()
	currentUnixTime := curtTime.UnixNano()
	now := currentUnixTime / nanoInMilli // 纳秒转毫秒
	t.Log(now, "  ", currentUnixTime>>20, "  ", "  \n")
	t.Log(time.Unix(0, int64(time.Millisecond)*now), "   ", time.Unix(0, int64(time.Millisecond)*(currentUnixTime>>20)))

	var number int64 = 1
	nowVal := now - epoch
	t.Log("now:", now, "  nowVal: ", nowVal)
	t.Log((nowVal) << timeShift)
	t.Log((nowVal)<<timeShift | (0 << workerShift) | 0)
	ID1 := (nowVal)<<timeShift | (1 << workerShift) | (number)
	t.Log(ID1)

	ID2 := (nowVal)<<timeShift | (2 << workerShift) | (number)
	t.Log(ID2)

	// convert int64 to []byte
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, nowVal)
	b := buf[:n]
	t.Logf("b %v , b string: %s\n", b, string(b))
	// convert []byte to int64
	x, n := binary.Varint(b)
	t.Logf("x is: %v, n is: %v\n", x, n)

}

/*
>>> v1 = 44920820475
>>> f'{v1:b}'
'101001110101011111010101001011111011'
>>> v2=v1<<22
>>> f'{v2:b}'
'1010011101010111110101010010111110110000000000000000000000'
>>> len(f'{v2:b}')
58
>>> v1 = 1570626560124
>>> v2=v1<<22
>>> v2
6587685263634333696
>>> f'{v2:b}'
'101101101101100001010001101100000011111000000000000000000000000'
>>> len(f'{v2:b}')
63
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
	for i := 0; i < count; i++ {
		id := <-ch
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

/*
https://juejin.im/post/5c75132f51882562276c5065

41位，用来记录时间戳（毫秒）。
3) 41位可以表示2^41−1个数字，如果只用来表示正整数（计算机中正数包含0），
可以表示的数值范围是：0 至 2^41−1，减1是因为可表示的数值范围是从0开始算的，而不是1。
也就是说41位可以表示2^41−1个毫秒的值，转化成单位年则是(2^41−1)/(1000∗60∗60∗24∗365)=69年

>>> (2**41 - 1) / (1000*60*60*24*365)
69.73057000098301
41bit:用来记录时间戳，这里可以记录69年，如果设置好起始时间比如今年是2018年，那么可以用到2089年，到时候怎么办？要是这个系统能用69年，我相信这个系统早都重构了好多次了。

sonyflake 是 39 bits for time in units of 10 msec
>>> (2**39-1)/(100*60*60*24*365)
174.32642500221968

标识在 snowflake_test
*/

func TestParse(t *testing.T) {
	v, err := strconv.ParseInt("999999999999999999", 10, 64)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)
	unicode.IsNumber(1)
}

func TestV1(t *testing.T) {
	x := ^uint32(0) // x is 0xffffffff
	i := int(x)     // i is -1 on 32-bit systems, 0xffffffff on 64-bit
	t.Log(i)

	f1 := 0xffffffff
	t.Log(f1)

	f2 := 0xffffffff
	t.Log(int32(f2))

	y, m, d := time.Now().Date()
	t.Log(y, m, d)
	t.Logf("%02d", int(time.Now().Month()))
}

func TestV2(t *testing.T) {
	t.Log(makeMilliseconds())
}
func makeMilliseconds() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

//  64 bit
//  1 bit unused | 39 id  549755813888   | 15 year 2**15=32768  | 4 month 2**4=16 | 5 day 2**5=32 |
//                                       |                           24                           |

func MakeLong(left int32, right int32) int64 {
	//implicit conversion of left to a long
	var res int64 = int64(left)

	//shift the bits creating an empty space on the right
	// ex: 0x0000CFFF becomes 0xCFFF0000
	res = res << 32

	//combine the bits on the right with the previous value
	// ex: 0xCFFF0000 | 0x0000ABCD becomes 0xCFFFABCD
	res = res | int64(right) //uint first to prevent loss of signed bit

	//return the combined result
	return res
}
