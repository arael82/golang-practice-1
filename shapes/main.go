package main

import (
	"fmt"
	"math"
)

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	t := triangle{
		height: 2,
		base:   3,
	}

	s := square{
		sideLength: 4,
	}

	printArea(t)
	printArea(s)
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
}

func (s square) getArea() float64 {
	return (math.Pow(s.sideLength, 2))
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
