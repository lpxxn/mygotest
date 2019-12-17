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
}

func TestSonyFlaked2(t *testing.T) {
	var i int
	start := time.Now()
	for ; time.Since(start) < time.Second; {
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
