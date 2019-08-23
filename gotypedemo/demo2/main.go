package main

import "github.com/mygotest/gotypedemo/demo2/libs"

func main() {
	var lib libs.TLib = libs.NewLib("abde")
	lib.String()

}
