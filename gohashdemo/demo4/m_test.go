package main_test

import (
	"strings"
	"testing"
	"unicode"
)

// https://www.practical-go-lessons.com/chap-27-enum-iota-and-bitmask
func TestName(t *testing.T) {
	var x, x1 uint8
	x = 1
	t.Logf("%08b\n", x)
	x1 = 2
	t.Logf("%08b\n", x1)
	y := x | x1
	t.Logf("%08b\n", y)
	z := y & x
	t.Logf("%08b\n", z)

	var roles byte = isCaptain | isMedic | canFlyJupiter
	//Prints a binary representation.
	t.Logf("%b\n", roles)
	t.Logf("%b\n", isCaptain)
	t.Logf("%b\n", isTrooper)
	t.Logf("%b\n", isMedic)

	t.Logf("Is Captain? %v\n", isCaptain&roles == isCaptain)
	t.Logf("Is Trooper? %v", isTrooper&roles == isTrooper)

	t.Log(isDig("12341234123401234234"))
	t.Log(isDig("12341234123401234a234"))
	t.Log(isDig("12341234#123041234234"))

	 s1, s2, s3 := StuList()
	 z1 := append(s1, s2...)
	 t.Log(z1)
	 t.Log(append(z1, s3...))

}

func StuList() (s1, s2, s3 []*Stu) {
	return
}

type Stu struct {
	Age int
	Name string
}

func isDig(s string) bool {
	return strings.IndexFunc(s, func(c rune) bool { return c < '0' || c > '9' }) == -1
}

func isDig2(s string) bool {
	return strings.IndexFunc(s, unicode.IsDigit) == -1
}

const (
	isCaptain = 1 << iota
	isTrooper
	isMedic

	canFlyMars
	canFlyJupiter
	canFlyMoon
)