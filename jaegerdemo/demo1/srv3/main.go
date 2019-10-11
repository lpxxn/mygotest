package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	utils2 "github.com/lpxxn/gotest/app3_ginkgo/utils"
	"github.com/mygotest/jaegerdemo/demo1/common"
	"github.com/mygotest/jaegerdemo/demo1/middleware"
	"github.com/mygotest/jaegerdemo/demo1/utils"
	"github.com/opentracing/opentracing-go"
)

func main() {
	gin.SetMode(gin.DebugMode)
	tracer, closer := utils.NewJaegerTracer("srv3", common.JaegerHostPort)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	engine := gin.New()
	engine.Use(middleware.SetupRootTrace())

	engine.GET("remote3", UserInfo)

	if err := engine.Run(":9656"); err != nil {
		log.Fatal(err)
	}
}

func UserInfo(c *gin.Context) {
	utils.RandomSleep(0, 1)
	c.JSON(http.StatusOK, common.UserInfo{ID: utils.RandomInt(10, 100), Name: utils2.RandomStr(utils.RandomInt(3, 5))})
}
