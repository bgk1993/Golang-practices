package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

func main() {
	s := square{
		sideLength: 33,
	}

	t := triangle{
		base:   13,
		height: 21,
	}

	fmt.Println("Square Area")
	printArea(s)
	fmt.Println("Triangle Area")
	printArea(t)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
