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
	tracer, closer := utils.NewJaegerTracer("srv2", common.JaegerHostPort)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	engine := gin.New()
	engine.Use(middleware.SetupRootTrace())

	engine.GET("remoteUserInfo", UserInfo)

	if err := engine.Run(":9655"); err != nil {
		log.Fatal(err)
	}
}

func UserInfo(c *gin.Context) {
	if _, err := utils.GetRequest("http://127.0.0.1:9656/remote3", func(r *http.Request) {
		span, found := middleware.GetSpan(c)
		if found {
			utils.InjectTraceID(span.Context(), r.Header)
		}
	}); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{"Err": err})
		return
	}
	utils.RandomSleep(1, 3)
	c.JSON(http.StatusOK, common.UserInfo{ID: utils.RandomInt(10, 100), Name: utils2.RandomStr(utils.RandomInt(3, 5))})
}
