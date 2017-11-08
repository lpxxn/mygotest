package builder

import (
	"strconv"
)

type Color string
type Make string
type Model string

const (
	BLUE Color = "blue"
	Red			= "red"
)

type Car interface {
	Drive() string
	Stop() string
}

type CarBuilder interface {
	TopSpeed(int) CarBuilder
	Paint(Color) CarBuilder
	Build() Car
}

type CarBuildOne struct {
	SpeedOption int
	Color Color
}

type CarOne struct {
	TopSpeed int
	Color Color
}

func (c *CarOne) Drive() string {
	return "Driving at sppd: " + strconv.Itoa(c.TopSpeed)
}

func (c *CarOne) Stop() string {
	return "Stop a " + string(c.Color) + " car"
}

func (cb *CarBuildOne) TopSpeed(speed int) CarBuilder {
	cb.SpeedOption = speed
	return cb
}

func (cb *CarBuildOne) Paint(color Color) CarBuilder {
	cb.Color = color
	return cb
}

func (cb *CarBuildOne) Build() Car {
	return &CarOne{
		TopSpeed: cb.SpeedOption,
		Color:cb.Color,
	}
}


func New() CarBuilder {
	return &CarBuildOne{}
}


