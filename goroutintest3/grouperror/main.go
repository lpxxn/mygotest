package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)


func main() {
	ErrGroupTest1()
}

func ErrGroupTest1() {
	var g errgroup.Group
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.abogadfle.com/",
		"http://www.somestupidname.com/",
	}

	for _, uri := range urls {
		currentUri := uri
		g.Go(func() error {
			if resp, err := http.Get(currentUri); err != nil {
				return err
			} else {
				resp.Body.Close()
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("test1 end")
}