package main

import (
	"time"
	"fmt"
	"context"
	"net/http"
	"github.com/go-errors/errors"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)

	// Make a request, that will call the google homepage
	req, _ := http.NewRequest(http.MethodGet, "https://baidu.com", nil)
	//req, _ := http.NewRequest(http.MethodGet, "https://google.com", nil)
	// Associate the cancellable context we just created to the request
	req = req.WithContext(ctx)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Reqeust failed:", err)
		return
	}
	defer  cancel()
	fmt.Println("response recevied , status code: ", res.StatusCode)

}



