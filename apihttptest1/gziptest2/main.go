package main

import (
	"net/http"
	"io"
	"io/ioutil"
	"fmt"
	"bytes"
	"compress/gzip"
	"github.com/fatih/stopwatch"
)

func main() {
	client := new(http.Client)

	url := "http://13.124.215.155:9002/TestManagerApi"

	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("accept-encoding", "gzip")

	s := stopwatch.Start(0)
	response, _ := client.Do(request)
	duration := s.ElapsedTime()
	defer  response.Body.Close()
	fmt.Println("request use time: ", duration)


	var reader io.Reader = response.Body

	body2, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Printf("read response error: %v", err)
		return
	}
	duration = s.ElapsedTime()
	fmt.Println("read body time: ", duration)
	fmt.Printf("body length: %d \n content : %s \n", len(body2), string(body2))


	var zipReader *bytes.Reader = bytes.NewReader(body2)
	reader, _ = gzip.NewReader(zipReader)
	body2, err = ioutil.ReadAll(reader)
	duration = s.ElapsedTime()
	fmt.Println("zip read body time: ", duration)
	fmt.Printf("body length: %d \n content : %s \n", len(body2), string(body2))


}
