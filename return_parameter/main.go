package main

import "fmt"

type terp string

func (terp) Rp() (a, b string) {
	a = "aaa"
	b = "bbb"
	return
}

func (t terp) RpW() (a, b string) {
	return t.Rp()
}

func main() {
	var t terp
	ra, rb := t.RpW()
	fmt.Println(ra, "  ", rb)
}