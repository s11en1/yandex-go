package main

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.width * rectangle.height
}

type Circle struct {
	radius float64
}

func (circle Circle) Area() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}
