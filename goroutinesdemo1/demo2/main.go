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

	test1 := make(chan *int)
	// 指针channel 提前关闭后写入报错
	// 读取的时候指针为空，取值会报错
	close(test1)
	go func() {
		defer func() {
			if e := recover(); e != nil {
				fmt.Println(e)
			}
		}()
		var v  int= 10
		test1 <- &v
	}()
	v1, ok := <-test1
	fmt.Println(v1, ok)
}
