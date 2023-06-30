package main

func main() {
	circle1 := Circle{
		Shape: Shape{
			X:     1,
			Y:     1,
			Color: "Red",
		},
		Radius: 15,
	}

	circle1.Print()

	circle1Copy := circle1.Clone()

	circle1Copy.Print()
}
