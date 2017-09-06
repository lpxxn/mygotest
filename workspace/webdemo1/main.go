package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	"context"
	"os"
	"os/signal"
	"time"
	"github.com/mygotest/workspace/webdemo1/src/urls"
	"github.com/gin-contrib/static"

)

// Binding from JSON
type Login struct {
	User     string `form:"user1" json:"user1" binding:"required"`
	Password string `form:"password1" json:"password1" binding:"required"`
}

func init() {
	fmt.Println("init func")
}

func main() {
	r := gin.Default()
	// github.com/gin-contrib/static
	r.Use(static.Serve("/", static.LocalFile("./src/www", true)))
	//r.Static("/", "./src/www/index.html")

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "hello world my name is lp"})
	})

	//gin.SetMode(gin.ReleaseMode)
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

	for url, value := range urls.UrlsGetmap {
		r.GET(url, value)
	}
	for url, value := range urls.UrlsPostMap {
		r.POST(url, value)
	}

	/*
	capi := r.Group("api")

	{
		capi.GET("/ping", func(c *gin.Context) {
			fmt.Println("paing inner")
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		// Example for binding JSON ({"user": "manu", "password": "123"})
		capi.POST("/loginJSON", func(c *gin.Context) {
			var json Login
			if c.BindJSON(&json) == nil {
				if json.User == "manu" && json.Password == "123" {
					c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
				} else {
					c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
				}
			}
		})

	}

	*/

	//
	srv := &http.Server{
		Addr:    ":9064",
		Handler: r,
	}

	// start the https server
	go func() {
		//if err := http.ListenAndServeTLS(":9064", "server.crt", "server.key", r); err != nil {
		//	fmt.Printf("listen: %s \n", err)
		//}

		if err := srv.ListenAndServeTLS("server.crt", "server.key"); err != nil {
			fmt.Printf("listen: %s \n", err)
		}
	}()

	// start http server
	go http.ListenAndServe(":9065", r)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	fmt.Println("Server is Running....")
	<-quit
	fmt.Println("Shutdown Server .....")

	// test timeout
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Server Shutdown: %v \n", err)
	}

	// test timeout
	//select {
	//case <-time.After(1 * time.Second):
	//	fmt.Println("overslept")
	//case <-ctx.Done():
	//	fmt.Println("ctx.Done()")
	//	fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	//}

	fmt.Println("server exists")

	//r.RunTLS(":9065", "./server.crt", "./server.key")
	//r.Run() // listen and serve on 0.0.0.0:8080
}
