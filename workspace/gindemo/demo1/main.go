package main

import (
	"github.com/gin-gonic/gin"
	"strings"
	"net/http"
	"fmt"
	gerr "github.com/go-errors/errors"
	"go.uber.org/zap"
	"github.com/mygotest/workspace/gindemo/demo1/interviewlogger"
	"github.com/mygotest/workspace/gindemo/demo1/zaplogger"
)

func main() {
	r := gin.Default()
	r.Use(globalRecover)

	r.GET("hi", func(context *gin.Context) {
		panic("error")
	})

	g := r.Group("t2")
	g.Use(InterViewLog)
	g.POST("u1", func(c *gin.Context) {
		var obj interface{}
		c.Bind(&obj)
		// body is already read
		// 如果提前读取，那么c.Bind()方法无效
		//fmt.Println(obj)
		//x, _ := ioutil.ReadAll(c.Request.Body)
		//fmt.Printf("aaaaaaaaa %s", string(x))
		c.JSON(http.StatusOK, gin.H{"ok": "ok"})
	})
	g.GET("u2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": "ok"})
	})
	r.GET("u1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": "ok"})
	})
	if err := http.ListenAndServe(":9100", r); err != nil {
		panic(err)
	}
	fmt.Println("server run port :", "9100")

	select {}
}


func InterViewLog(c *gin.Context) {

	defer interViewLogDetail(c)
	c.Next()
}

func interViewLogDetail(c *gin.Context) {
	defer func() {
		if err:= recover(); err != nil {
			zaplogger.Error("InterViewLogger Error:", zap.String("url", c.Request.RequestURI), zap.Error(err.(error)))
		}
	}()


	param := c.Request.URL.String()

	if c.Request.Method == "POST" {
		var obj interface{}
		c.Bind(&obj)
		fmt.Println(obj)
		interviewlogger.LogInterView("interview", zap.String("url", param))
	}
	interviewlogger.LogInterView("interview", zap.String("url", param))

}

func globalRecover(c *gin.Context) {
	defer func(c *gin.Context) {
		if rec := recover(); rec != nil {
			// that recovery also handle XHR's
			// you need handle it
			err := gerr.Wrap(rec, 2).ErrorStack()
			if XHR(c) {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err,
				})
			} else {
				//c.HTML(http.StatusOK, "500", gin.H{})
				c.JSON(http.StatusInsufficientStorage, gin.H{"error": err})
			}
		}
	}(c)
	c.Next()
}

func XHR(c *gin.Context) bool {
	return strings.ToLower(c.Request.Header.Get("X-Requested-With")) == "xmlhttprequest"
}

