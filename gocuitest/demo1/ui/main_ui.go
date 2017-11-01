package ui

import (
	"github.com/jroimartin/gocui"
	"log"
	"fmt"
)

var (
	ViewaArr = []string{"side", "sendmsg", "receivemsg", "sendmsgeidter"}
	active   = 0
)

type MainUi struct {
	MainGui *gocui.Gui
}

func (ui *MainUi)Init(){
	g := ui.MainGui
	fmt.Println(g.Size())


	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorGreen

	g.SetManagerFunc(Layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panic(err)
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		log.Panic(err)
	}
}

func (ui *MainUi) SetSideUserInfo(users []string) error {
	v, err := ui.MainGui.View(ViewaArr[0])
	if err != nil {
		log.Panic(err)
		return err
	}
	v.Clear()
	for item, _ := range users{
		fmt.Fprintln(v, item)
	}
	//ui.MainGui.Flush()
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func nextView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (active + 1) % len(ViewaArr)
	name := ViewaArr[nextIndex]

	if _, err := SetCurrentViewOnTop(g, name); err != nil {
		return err
	}
	active = nextIndex

	return nil
}

func Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	sideWidth := int(0.2 * float32(maxX))
	if v, err := g.SetView(ViewaArr[0], -1, -1, sideWidth, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "用户"
		v.Wrap = true
		//v.Autoscroll = true
		fmt.Fprintln(v, "aaa")
		fmt.Fprintln(v, "bbbb")
	}
	remainWidth := maxX - sideWidth
	halfRemainWidth := int(0.5 * float32(remainWidth))
	msgMaxY := maxY - 10
	msg1EndX := sideWidth + halfRemainWidth
	if v, err := g.SetView(ViewaArr[1], sideWidth, -1, msg1EndX, msgMaxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Send Msg"
		v.Wrap = true
		v.Autoscroll = true
		if _, err := SetCurrentViewOnTop(g, ViewaArr[0]); err != nil {
			log.Panic(err)
		}
	}
	if v, err := g.SetView(ViewaArr[2], msg1EndX, -1, maxX, msgMaxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Receive Msg"
		v.Wrap = true
		v.Autoscroll = true
	}

	if v, err := g.SetView(ViewaArr[3], sideWidth, msgMaxY, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Edit Msg"
		v.Wrap = true
		//v.Autoscroll = true
		v.Editable = true
	}
	return nil
}

func SetCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

