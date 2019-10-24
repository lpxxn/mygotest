package main_test

import (
	"testing"
	"regexp"
)

func Benchmark_Reg(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		if ok, _ := regexp.MatchString(`^ZONE\d*$`, "ZONE12122323"); ok {
			//fmt.Println("ZONE12122323 ok")
		}

		//if ok, _ := regexp.MatchString(`^ZONE\d*$`, "zone12122323"); ok {
		//	//fmt.Println("zone12122323 ok")
		//}
	}
}

var re = regexp.MustCompile(`^ZONE\d*$`)
// good
func Benchmark_Reg2(b *testing.B) {
	//var re = regexp.MustCompile(`^ZONE\d*$`)

	for i := 0; i < b.N; i ++ {
		if ok := re.MatchString("ZONE12122323"); ok {
			//fmt.Println("ZONE12122323 ok")
		}

		//if ok, _ := regexp.MatchString(`^ZONE\d*$`, "zone12122323"); ok {
		//	//fmt.Println("zone12122323 ok")
		//}
	}
}

func TestCafeteriaName(t *testing.T) {
	var re = regexp.MustCompile(`^<.*>(.*)~.*\[(.*)_[0-9]*\]$`)
	var str = `<食堂菜单>2019-03-18~infinity[按天重复_27]`

	for i, match := range re.FindAllString(str, -1) {
		t.Logf(match, "found at index\n", i)
	}
}