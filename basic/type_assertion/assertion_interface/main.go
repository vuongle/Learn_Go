package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	var s Shape

	// Type assertion to call the Area method on the Rectangle struct
	s = Rectangle{width: 2, height: 3}
	fmt.Println(s.(Rectangle).Area())

	// Type assertion to call the Area method on the Circle struct
	s = Circle{radius: 1}
	fmt.Println(s.(Circle).Area())
}
