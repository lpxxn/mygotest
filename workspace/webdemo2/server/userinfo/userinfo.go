package userinfo

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mygotest/workspace/webdemo2/server/parammodes"
	"github.com/mygotest/workspace/webdemo2/utils"
	"net/http"
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
//     Security:
//       api_key:
//       oauth: read, write

//     Responses:
//       default: pingResponse
//       200: pingResponse
//        description: succession
//       422:
//        description: false
func SetUserInfo(c *gin.Context) {
	var param = &parammodes.UserParamInfo{}
	if c.BindJSON(param) != nil {
		c.JSON(http.StatusOK, gin.H{"Status": false})
		return
	}
	pjson, _ := json.Marshal(param)
	str, err := utils.Cluster.Set("gouser", pjson, 0).Result()
	fmt.Println(str, err)
	c.JSON(http.StatusOK, gin.H{"Status": true})
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
