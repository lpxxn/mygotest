package server

import (
	"fmt"
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
	r.GET("pets", listPets)
}

func apiRouter(r *gin.Engine) {
	api := r.Group("api")
	api.GET("pingapi", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"apiMsg": "Hello Api!"})
	})
}

// swagger:parameters listPets
type PetQueryFlags struct {
	// Status
	Status string `json:"status"`
	Id     string `json:"id"`
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    int32 `json:"code"`
		Message error `json:"message"`
	} `json:"body"`
}

// GetPets swagger:route GET /pets pets listPets
//
// Lists the pets known to the store.
//
// By default it will only lists pets that are available for sale.
// This can be changed with the status flag.
//
// Responses:
// 		default: genericError
// 		    200: []pingResponse
func listPets(c *gin.Context) {
	strpar := c.DefaultQuery("status", "1")
	// param in path
	strname := c.Param("id")
	fmt.Println(strpar)
	fmt.Println(strname)
	resp := make([]models.RspPing, 0)
	resp = append(resp, models.RspPing{Msg: "abcdef"})
	c.JSON(http.StatusOK, resp)

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
