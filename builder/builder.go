package builder

import (
	"fmt"
)

type Speed int
type Color string
type Wheel string

const (
	RED Color = "red"
	BLUE      = "blue"
	GREEN     = "green"
)

const (
	NormalWheel Wheel = "normal"
	SportsWheel       = "sports"
)

type car struct {
	topSpeed Speed
	wheel Wheel
	color Color
}

type Builder interface {
	Wheel(Wheel) Builder
	TopSpeed(int) Builder
	Paint(Color) Builder
	Build() Car
}

type Car interface {
	Drive() string
	Stop() string
}

func New() Builder {
	return &car{}
}

func (c *car) Wheel (wheel Wheel) Builder {
	c.wheel = Wheel(wheel)
	return c
}

func (c *car) TopSpeed (speed int) Builder {
	c.topSpeed = Speed(speed)
	return c
}

func (c *car) Paint (color Color) Builder {
	c.color = Color(color)
	return c
}

func (c *car) Build () Car {
	return &car{
		c.topSpeed,
		c.wheel,
		c.color,
	}
}

func (c *car) Drive () string {
	return fmt.Sprintf("%s Color + %s Wheel Driving at speed: %d",
		string(c.color),
		string(c.wheel),
		int(c.topSpeed),
	)
}

func (c *car) Stop() string {
	return fmt.Sprintf("Stopping a")
}