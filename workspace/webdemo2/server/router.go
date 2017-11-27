package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	newRouter := gin.Default()

	newRouter.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{"code": "NOT_FOUND", "msg": "sorry"})
	})

	userRouter(newRouter)
	apiRouter(newRouter)
	return newRouter
}

func userRouter(r *gin.Engine) {
	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "Hello World !"})
	})
}

func apiRouter(r *gin.Engine) {
	api := r.Group("api")
	api.GET("pingapi", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"apiMsg": "Hello Api!"})
	})
}
