package userinfo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mygotest/workspace/webdemo2/models"
	"github.com/mygotest/workspace/webdemo2/server/parammodes"
	"github.com/mygotest/workspace/webdemo2/utils"
	"golang.org/x/time/rate"
)

// swagger:route Post /setuserinfo users upUserParam
//
// 修改用户 by some parameters.
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
//     Responses:
//       default: operateRev
//       200: operateRev
//        description: succession
//       422:
//        description: false
func SetUserInfo(c *gin.Context) {
	var param = &parammodes.UserParamInfo{}
	if c.BindJSON(param) != nil {
		c.JSON(http.StatusOK, gin.H{"Status": false})
		return
	}
	fmt.Println(param)
	pjson, _ := json.Marshal(param)
	fmt.Println(string(pjson))
	str, err := utils.Cluster.Set("gouser", pjson, 0).Result()
	fmt.Println(str, err)
	rev := parammodes.OperateResult{RevValueBase: parammodes.RevValueBase{Status: true}}
	c.JSON(http.StatusOK, rev)
}

func GetUserInfo(c *gin.Context) {
	str, err := utils.Cluster.Get("gouser").Result()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Status": false})
		return
	}
	u := &parammodes.UserParamInfo{}
	json.Unmarshal([]byte(str), u)
	c.JSON(http.StatusOK, gin.H{"Status": true, "Data": u})
}

// test  private
type VP struct {
	A models.ResPingBody
}

func TestP() {
	var t = models.GetTestPrivate()
	t.Name()
	rate.NewLimiter()
}
