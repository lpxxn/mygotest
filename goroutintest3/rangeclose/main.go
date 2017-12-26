package main

import "fmt"
// it is possible to close a non-empty channel but still have the remaining values be received.
func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	var f = fmt.Println

	for elem := range queue {
		f(elem)
	}
}
