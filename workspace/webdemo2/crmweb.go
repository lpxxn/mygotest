package main

import (
	"fmt"

	"github.com/mygotest/workspace/webdemo2/utils"
)

func main() {
	//_, err := utils.Cluster.Set("crmweb", "value111", 0).Result()
	//fmt.Println(err)
	ipstr := utils.HostLocalIp()
	fmt.Println(ipstr)

}
