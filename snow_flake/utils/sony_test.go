package utils

import (
	"testing"
	"time"
)

var sf *Sonyflake

func init() {
	var st Settings
	st.StartTime = time.Now()

	sf = NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}

}
func TestSonyTime(t *testing.T) {
	t.Log(toSonyflakeTime(time.Now()))
	sf.NextID()
	sf.NextID()
	sf.NextID()
	sf.NextID()
}

func TestSonyFlaked2(t *testing.T) {
	var i int
	start := time.Now()
	for time.Since(start) < time.Second {
		sf.NextID()
		i++
	}
	t.Log(i)
}

func TestLower16BitPrivateIP(t *testing.T) {
	t.Log(privateIPv4())
	t.Log(lower16BitPrivateIP())
	t.Log(uint16(255)<<8 + uint16(255)) //65535
	/*
		>>> 2 ** 16
		65536
	*/
}

func TestAdd1(t *testing.T) {
	// 二进制都要为1
	//const max = 255
	const max = 3

	t.Log(max)
	logValue := func(v uint16) uint16 {
		//const max = uint16(4)
		r := (v + 1) & max
		t.Log(v, r)
		return r
	}
	//var v uint16 = 2
	var i uint16 = 0
	for ; i < 10; i++ {
		logValue(i)
	}
}

func TestNumber1(t *testing.T) {
	h := 12
	min := 15
	fHour := func() {
		t.Log("before: ", h, "  ", min)
		minBit := 6
		a := h << minBit
		b := a | min
		hour := b >> minBit
		minRev := b & (1<<minBit - 1)
		t.Log(" after: ", hour, "  ", minRev)
		if h != hour || minRev != min {
			t.Fatal("error")
		}
	}
	fHour()
	h = 11
	min = 53
	fHour()

	h = 3
	min = 6
	fHour()
}
