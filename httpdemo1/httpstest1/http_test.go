package main_test

import (
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

type Say struct{}

func (s *Say) Anything(c *gin.Context) {
	log.Print("Received Say.Anything API request")
	log.Println(c.Request.UserAgent())
	log.Println("remote addr", c.Request.RemoteAddr)
	ip := c.Request.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = strings.Split(c.Request.RemoteAddr, ":")[0]
	}
	log.Println("x forwarded ip: ", ip)
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}

func TestHtt2Client(t *testing.T) {
	say := new(Say)

	router := gin.Default()
	router.GET("/greeter", say.Anything)
	if err := http.ListenAndServe(":5110", router); err != nil {
		panic(err)
	}
}
