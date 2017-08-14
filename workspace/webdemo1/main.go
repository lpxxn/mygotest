package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"

)

// Binding from JSON
type Login struct {
	User     string `form:"user1" json:"user1" binding:"required"`
	Password string `form:"password1" json:"password1" binding:"required"`
}



func main()  {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("paing inner")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Example for binding JSON ({"user": "manu", "password": "123"})
	r.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if c.BindJSON(&json) == nil {
			if json.User == "manu" && json.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		}
	})

	// Example for binding a HTML form (user=manu&password=123)
	r.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// This will infer what binder to use depending on the content-type header.
		if c.Bind(&form) == nil {
			if form.User == "manu" && form.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		}
	})
	r.RunTLS(":9065", "./ca.crt", "./ca.key")
	//r.Run() // listen and serve on 0.0.0.0:8080
}
