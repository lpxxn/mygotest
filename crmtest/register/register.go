package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	//"time"
)

func main() {
	//url := "http://localhost:5000/act/register"
	//url := "http://localhost:5000/act/register"

	go func() {
		for i := 0; i < 50; i++ {
			go GetTest(i)
		}
	}()
	go func() {
		for i := 0; i < 300; i++ {
			//time.Sleep(time.Duration(time.Second))
			//go RegisterUser(url, i)
			//go GetTest(i)
			go GetUsers(i)
		}
	}()
	stopSingal := make(chan os.Signal)
	signal.Notify(stopSingal, os.Interrupt)
	<-stopSingal
}

func GetUsers(i int) {
	data := strings.NewReader("{\"PageIndex\": 0, \"PageSize\": 100 }")
	url := "http://www.crmpro.com:8077/crmv2/userlists"
	req, _ := http.NewRequest("POST", url, data)
	req.Header.Add("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("res error:", err)
		return
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("index: ", i, "  body : ", string(body))

}

func GetTest(i int) {
	//data := strings.NewReader("{\"uid\": 1000, \"grp\": \"AgmTest\\\\group1\" }")
	url := "http://www.crmpro.com:8077/TestManagerApi"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("res error:", err)
		return
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("index: ", i, "  body : ", string(body))

}
func RegisterUser(url string, i int) {
	data := strings.NewReader("{\"uid\": 1000, \"grp\": \"AgmTest\\\\group1\" }")
	req, _ := http.NewRequest("POST", url, data)
	req.Header.Add("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("res error:", err)
		return
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("index: ", i, "  body : ", string(body))

}
