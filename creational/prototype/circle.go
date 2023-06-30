package main

import "fmt"

type Circle struct {
	Shape

	Radius int
}

func (c *Circle) Print() {
	fmt.Printf("Circle: x = %d, y = %d, color = %s, radius = %d", c.X, c.Y, c.Color, c.Radius)
}

func (c *Circle) Clone() IShape {
	return &Circle{
		Shape:  c.Shape,
		Radius: c.Radius,
	}
}
