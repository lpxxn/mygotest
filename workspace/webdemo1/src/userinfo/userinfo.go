package userinfo

import (

	//"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

type User struct {
	Name       string
	Age        int
	RequestUrl string
}

// Get User
func GetUserInfoById(c *gin.Context) {
	strpar := c.DefaultQuery("id", "1")
	// param in path
	strname := c.Param("name")
	id, _ := strconv.Atoi(strpar)
	users := []User{User{Name: "li", Age: 10}, {Name: strname, Age: 8, RequestUrl: c.Request.URL.RawQuery}}
	if id > 2 {
		users = append(users, User{Name: "Na", Age: 10, RequestUrl: c.Request.URL.RawQuery})
	}
	c.JSON(200, users)
	//rvjson, _ := json.Marshal(users)
	//c.JSON(200, string(rvjson))
}

func RequestOauth(c *gin.Context) {
	client := new(http.Client)
	//url := "http://localhost:4444/oauth2/auth"
	url := "http://localhost:4444/oauth2/auth?client_id=some-consumer&redirect_uri=http%3A%2F%2Flocalhost%3A9065%2Fcallback&response_type=code&scope=openid+offline+hydra.clients&state=lotatsztwtavexmwrjvjroxs&nonce=drfwevzauxnkkaqbkkebjcehv"
	res, _ := client.Get(url)
	defer res.Body.Close()

	reader := res.Body
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}

func CallBack(c *gin.Context) {
	err := c.Request.URL.Query().Get("error")
	if err != "" {
		myHtmlStr := fmt.Sprintf(
				`<html>
							<body>
								<h1>Error</h1>
								%s
							</body>
						</html>
						`, err)
			c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write([]byte(myHtmlStr))
		return
	}

	c.JSON(200, gin.H{"rowQuery": c.Request.URL.RawQuery, "url": c.Request.URL})
}
