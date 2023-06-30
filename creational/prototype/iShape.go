package main

type IShape interface {
	Print()
	Clone() IShape
}

type Shape struct {
	X     int
	Y     int
	Color string
}
