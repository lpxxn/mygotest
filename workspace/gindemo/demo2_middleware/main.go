package main

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(globalMiddleware())

	router.GET("/rest/n/api/*some", mid1(), mid2(), handler)
	router.GET("redirect", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "http://www.so.com?a=b")
	})

	router.GET("redirect2", func(context *gin.Context) {
		context.Redirect(http.StatusCreated, "/redirect?a=b")
	})

	router.GET("redirect3", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/redirect?a=b")
	})

	router.GET("render", func(c *gin.Context) {
		RenderJson(c, &User{Name: "lipeng"})
	})

	srv := &http.Server{
		Addr:         ":8898",
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  5 * time.Minute,
	}
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
	// router.Run(":8898")
}

func globalMiddleware() gin.HandlerFunc {
	fmt.Println("globalMiddleware...1")

	return func(c *gin.Context) {
		fmt.Println("globalMiddleware...2")
		c.Next()
		fmt.Println("globalMiddleware...3")
	}
}

func handler(c *gin.Context) {
	c.JSON(http.StatusOK, struct {
		Name string `json:"name"`
	}{Name: "lipeng"})
	fmt.Println("exec handler.")
}

func mid1() gin.HandlerFunc {
	fmt.Println("mid1...1")

	return func(c *gin.Context) {

		fmt.Println("mid1...2")
		c.Next()
		fmt.Println("mid1...3")
	}
}

func mid2() gin.HandlerFunc {
	fmt.Println("mid2...1")

	return func(c *gin.Context) {
		fmt.Println("mid2...2")
		c.Next()
		fmt.Println("mid2...3")
	}
}

type User struct {
	Name string `json:"name"`
}

func (u *User) BeforeRender() {
	u.Name += ""
}

var _ BeforeRender = &User{}

func RenderJson(c *gin.Context, data interface{}) {
	customerRender(data)
	c.JSON(http.StatusOK, data)
}

func customerRender(data interface{}) {
	refVal := reflect.ValueOf(data)
	if refVal.Kind() == reflect.Ptr && refVal.IsNil() {
		return
	}
	if cr, ok := refVal.Interface().(BeforeRender); ok {
		cr.BeforeRender()
	}
}

type BeforeRender interface {
	BeforeRender()
}
