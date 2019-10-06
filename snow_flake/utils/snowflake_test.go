package utils

import (
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