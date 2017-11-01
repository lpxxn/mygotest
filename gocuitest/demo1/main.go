package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
	"github.com/mygotest/gocuitest/demo1/ui"
	"time"
)


func main() {

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panic(err)
	}
	mainUi := &ui.MainUi{MainGui:g}
	mainUi.Init()
	fmt.Println(g.Size())

	users := []string{"li", "peng", "zhang", "san"}
	go func() {
		time.Sleep(200)
		mainUi.SetSideUserInfo(users)
	}()
	defer g.Close()
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panic(err)
	}


}
