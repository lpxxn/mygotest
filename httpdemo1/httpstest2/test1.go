package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var tr *http.Transport = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}
var client *http.Client = &http.Client{Transport: tr}

func getJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	// de := json.NewDecoder(resp.Body)
	// fmt.Println(de)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err2 := json.Unmarshal([]byte(body), target)
	return err2
	// bodyString := string(body)
	// fmt.Println(bodyString)

	//return de.Decode(target)
}

type MyDataUserDetail struct {
	Login int64  `json:"login"`
	Name  string `json:"name"`
	Id    string `json:"id"`
}
type MyDataDetail struct {
	Total int                `json:"total"`
	Users []MyDataUserDetail `json:"users"`
}

type MyData struct {
	State bool         `json:"state"`
	Data  MyDataDetail `json:"data"`
}

func main() {
	data := new(MyData) // or &MyData{}
	getJson("http://localhost:60377/TestManagerApi", data)
	fmt.Println(data)
}
