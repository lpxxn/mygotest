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

	r.Use(static.Serve("/", static.LocalFile("./tmpfs", true)))
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "test")
	})
	// Listen and Server in 0.0.0.0:3000  
	r.Run(":3000")
}
