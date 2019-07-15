package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

	router.Run(":8898")
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
