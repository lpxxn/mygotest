package main

import (
	"github.com/gin-gonic/gin"
	"strings"
	"net/http"
	"fmt"
	gerr "github.com/go-errors/errors"
)

func main() {
	r := gin.Default()
	r.Use(globalRecover)

	r.GET("hi", func(context *gin.Context) {
		panic("error")
	})
	if err := http.ListenAndServe(":9100", r); err != nil {
		panic(err)
	}
	fmt.Println("server run port :", "9100")

	select {}
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

