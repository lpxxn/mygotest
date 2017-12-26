package main

import (
	"os"
	"bufio"
	"strings"
	"sync"
	"fmt"
)

// 这是一个错误案例
// go run --race main.go  *.txt
func main() {
	var wg sync.WaitGroup

	w := newWords()
	for _, f := range os.Args[1:] {
		wg.Add(1)
		go func(file string) {
			if err := tallWords(file, w); err != nil {
				fmt.Println(err)
			}
			wg.Done()
		}(f)
	}
	wg.Wait()
	fmt.Println("words that appear more than once:")
	for word, count := range w.found {
		fmt.Println(word, ": ", count)
	}
}


type words struct {
	found map[string]int
}


func newWords() *words {
	//return &words{found: map[string]int {}}
	return &words{found: make(map[string]int)}
}


func (w *words) add(word string, n int) {
	count, ok := w.found[word]
	if !ok {
		w.found[word] = n
	} else {
		w.found[word] = count + 1
	}
}

func tallWords(filename string, dict *words) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		dict.add(word, 1)
	}
	return scanner.Err()
}

