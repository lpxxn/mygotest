package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	"context"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-contrib/static"
	"github.com/go-redis/redis"
	"github.com/mygotest/workspace/webdemo1/src/urls"
	"github.com/mygotest/workspace/webdemo1/tutorial"
	"net"
	"os"
	"os/signal"
	"strings"
	"time"

	"flag"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/mygotest/workspace/webdemo1/src/utils"
	"path"
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
	t := time.Now()
	fmt.Println(t)

	persondatajsonpath := flag.String("persiondatapath", "./mock_data/mock_person_data.json", "the path of person data")
	currentPath, err := os.Getwd()
	if err != nil {
		os.Exit(1)
		return
	}
	*persondatajsonpath = path.Join(currentPath, *persondatajsonpath)
	utils.PersonDataPath = *persondatajsonpath
	personData := utils.GetPersionInfo()
	fmt.Println(len(*personData))

	r := gin.Default()

	// sassion
	store, _ := sessions.NewRedisStore(10, "tcp", "192.168.0.105:6379", "", []byte("mysessionsecrit"))
	store.Options(sessions.Options{
		MaxAge:   86400,
		//Domain:".lp.com",
	})

	r.Use(sessions.Sessions("workino_session", store))
	// allow all origins
	r.Use(cors.Default())

	// github.com/gin-contrib/static
	r.Use(static.Serve("/", static.LocalFile("./src/www", true)))
	//r.Static("/", "./src/www/index.html")

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "hello world my name is lp"})
	})

	//gin.SetMode(gin.ReleaseMode)
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("paing inner")
		hostName, _ := os.Hostname()

		ifaces, err := net.Interfaces()
		var ips []string
		var ip string
		if err == nil {
			for _, i := range ifaces {
				addrs, err := i.Addrs()
				if err == nil {
					for _, addr := range addrs {
						var netIp net.IP
						switch v := addr.(type) {
						case *net.IPNet:
							netIp = v.IP
						case *net.IPAddr:
							netIp = v.IP
						}
						ips = append(ips, netIp.String())
					}
				}
			}
			ip = strings.Join(ips, " ,")
		}

		// redis
		client := redis.NewClient(&redis.Options{
			Addr:     "redistest:6379",
			Password: "",
			DB:       0,
		})
		var redisInfo string
		redisInfo = client.Ping().String()

		c.JSON(200, gin.H{
			"message":   "pong",
			"hostName":  hostName,
			"hostIp":    ip,
			"redisInfo": redisInfo,
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

	r.POST("/testproto1", func(c *gin.Context) {
		var p1 tutorial.Person
		if c.Bind(&p1) == nil {
			c.JSON(http.StatusOK, gin.H{"p1": p1})
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
		if err := http.ListenAndServeTLS(":9064", "server.crt", "server.key", r); err != nil {
			fmt.Printf("listen: %s \n", err)
		}

		//if err := srv.ListenAndServeTLS("server.crt", "server.key"); err != nil {
		//	fmt.Printf("listen: %s \n", err)
		//}
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
