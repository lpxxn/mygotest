package main

import (
	"net/http"
	"io"
	"io/ioutil"
	"fmt"
	"bytes"
	"compress/gzip"
)

func main() {
	client := new(http.Client)

	url := "http://13.124.215.155:9002/TestManagerApi"

	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("accept-encoding", "gzip")

	response, _ := client.Do(request)
	defer  response.Body.Close()

	var reader io.Reader = response.Body
	body2, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Printf("read response error: %v", err)
		return
	}
	fmt.Printf("body length: %d \n content : %s \n", len(body2), string(body2))


	var zipReader *bytes.Reader = bytes.NewReader(body2)
	reader, _ = gzip.NewReader(zipReader)
	body2, err = ioutil.ReadAll(reader)

	fmt.Printf("body length: %d \n content : %s \n", len(body2), string(body2))


}
