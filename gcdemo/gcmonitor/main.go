package main

import (
	"runtime/pprof"
	"flag"
	"os"
	"log"
	"runtime"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		Testgorouting();
		defer pprof.StopCPUProfile()
	}

	// ... rest of the program ...


	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics



		Testgorouting();

		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}

	time.Sleep(time.Second * 3)
}

func Testgorouting() {
	c1 := make(chan int, 20)
	for i := 0; i < 10; i++ {
		c1 <- i
	}

}

/*
// go run main.go -cpuprofile cpu.prof -memprofile mem.prof
              app name | file name
go tool pprof gcmonitor cpu.prof
go tool pprof gcmonitor mem.prof

go tool pprof -alloc_space mem.prof
go tool pprof -alloc_objects mem.prof



(pprof) list testgorouting
top 10
web

 go tool pprof -web gcmonitor ./mem.prof
*/