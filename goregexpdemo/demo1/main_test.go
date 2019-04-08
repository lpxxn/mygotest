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