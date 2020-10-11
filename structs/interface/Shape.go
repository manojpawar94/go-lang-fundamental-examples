package main

import (
	"math"
)

/*Shape interface*/
type Shape interface {
	area() float64
}

/*Circle struct to implement methods of Shape interface*/
type Circle struct {
	x, y, radius float64
}

/*Rectangle struct to implement methods of Shape interface*/
type Rectangle struct {
	x, y, width, height float64
}

/* define a method for circle (implementation of Shape.area())*/
func (circle Circle) area() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}

/* define a method for rectangle (implementation of Shape.area())*/
func (rectangle Rectangle) area() float64 {
	return rectangle.width * rectangle.height
}

/* define a method for shape */
func calcArea(shape Shape) float64 {
	return shape.area()
}
