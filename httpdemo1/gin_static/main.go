package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// if Allow DirectoryIndex
	//r.Use(static.Serve("/", static.LocalFile("/tmp", true)))
	// set prefix
	//r.Use(static.Serve("/static", static.LocalFile("/tmp", true)))

	//r.Use(static.Serve("/", static.LocalFile("/tmp", false)))
	r.Use(static.Serve("/", static.LocalFile("/Users/li/Downloads", true)))
	// 下面这个要index.html
	//r.Use(static.Serve("/", static.LocalFile("/Users/li/Downloads", false)))
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "test")
	})
	r.Run(":8088")
}
