package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	response, err := client.Get("http://localhost:60377/TestManagerApi")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
		fmt.Println(response.Body)
		body, _ := ioutil.ReadAll(response.Body)
		bodyString := string(body)
		fmt.Println(bodyString)
	}
	defer response.Body.Close()
}
