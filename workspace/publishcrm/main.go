package main

import (
	"fmt"
	sconfig "./serverconfig"

)

var data sconfig.AutoGenerated

func main() {
	fmt.Println("publish crm..")
	sconfig.TestFun()
}
