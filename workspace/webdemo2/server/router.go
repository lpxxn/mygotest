package server

import (
	"github.com/mygotest/workspace/webdemo2/models"
	"github.com/mygotest/workspace/webdemo2/server/userinfo"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	newRouter := gin.Default()
	newRouter.Use(cors.Default())
	newRouter.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{"code": "NOT_FOUND", "msg": "sorry"})
	})

	userRouter(newRouter)
	apiRouter(newRouter)
	return newRouter
}

func userRouter(r *gin.Engine) {
	r.GET("ping", ping)
	r.POST("setuserinfo", userinfo.SetUserInfo)
	r.GET("getuserinfo", userinfo.GetUserInfo)
}

func apiRouter(r *gin.Engine) {
	api := r.Group("api")
	api.GET("pingapi", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"apiMsg": "Hello Api!"})
	})
}

// swagger:route GET /ping pets users listPets
//
// Lists pets filtered by some parameters.
//
// This will show all available pets by default.
// You can get the pets that are out of stock
//
//     Consumes:
//     - application/json
//     - application/x-protobuf
//
//     Produces:
//     - application/json
//     - application/x-protobuf
//
//     Schemes: http, https, ws, wss
//
//     Security:
//       api_key:
//       oauth: read, write
//
//     Responses:
//       default: pingResponse
//       200: pingResponse
//        description: succession
//       422:
//        description: false
func ping(c *gin.Context) {
	pingModel := models.RspPing{Msg: "Hello World!", CurrentTime: time.Now()}
	c.JSON(http.StatusOK, pingModel)
}
