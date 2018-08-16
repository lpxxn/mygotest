package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mygotest/packrdemo/gindemo1/utils"
	"github.com/gobuffalo/packr"
)

var a = "aaa"
var b = []string{"a", "b", "c"}

func main() {
	r := gin.Default()
	fmt.Println(b)
	a = "bvb"
	fmt.Println(a)
	// if Allow DirectoryIndex
	//r.Use(static.Serve("/", static.LocalFile("/tmp", true)))
	// set prefix
	//r.Use(static.Serve("/static", static.LocalFile("/tmp", true)))
	box := packr.NewBox("./tmpfs")
	r.Use(utils.Serve("/", box))
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "test")
	})
	// Listen and Server in 0.0.0.0:3000  
	r.Run(":3000")
}
