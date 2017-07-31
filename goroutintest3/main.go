package main

import (
	"fmt"
	"sync"
)

type Result struct {
	Field   string
	Content string
}

type Info struct {
	Name string
	Desc string
}

func main() {

	// results := make(chan *Result, 3)
	results := make(chan *Result)
	var waitGrout sync.WaitGroup

	feeds1 := []Info{
		Info{Name: "li", Desc: "liliiiiii"},
		Info{Name: "peng", Desc: "penggggggg"},
		Info{Name: "yi", Desc: "yyyyyyyyyyyyy"},
		Info{Name: "er", Desc: "errrrrrrrrrr"},
		Info{Name: "san", Desc: "sansssssssss"}}

	feeds2 := []Info{
		Info{Name: "si", Desc: "si"},
		Info{Name: "wu", Desc: "Wu"},
		Info{Name: "liu", Desc: "Liu"}}

	feeds := [][]Info{feeds1, feeds2}
	fmt.Println(feeds, results)

	waitGrout.Add(len(feeds))
	for _, feed := range feeds {
		fmt.Println(feed)
		go func(feed []Info) {
			Match(feed, results)
			waitGrout.Done()
		}(feed)
	}

	go func() {
		waitGrout.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("channel len ", len(results), "cap :", cap(results), " -> result.Field : ", result.Field, " Content : ", result.Content)
	}
}

func Match(feeds []Info, results chan<- *Result) {
	for _, result := range feeds {
		results <- &Result{Field: result.Name, Content: result.Desc}
		fmt.Println("channel <- result Name: ", result.Name, " Desc: ", result.Desc)
	}
}
