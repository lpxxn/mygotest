package main

import "fmt"

var battle = make(chan string)

func warrior(name string, done chan struct{}) {
	defer func() {
		if re := recover(); re != nil {
			fmt.Println(re)
		}
	}()
	select {
	case opponent := <-battle:
		fmt.Printf("%s beat %s \n", name, opponent)
	case battle <- name:
		fmt.Println(battle, "  ", name)

	}
	done <- struct{}{}
}

func main() {
	done := make(chan struct{})
	langs := []string{"Go", "C", "C++", "Java", "Perl", "Python"}

	go func() {
		done <- struct{}{}
	}()
	// 关闭后不可以写panic 可以读
	//close(done)
	//done <- struct{}{}
	x, ok := <-done
	fmt.Println(ok, "  ", x)
	//fmt.Println(<-done)
	for _, l := range langs{
		go warrior(l, done)
	}
	for _ = range langs {
		<-done
	}
}
