package main

import (
	"fmt"
)

// Define an interface named Shape
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Define a struct named Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Implement the Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Implement the Perimeter method for Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Define a struct named Circle
type Circle struct {
	Radius float64
}

// Implement the Area method for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Implement the Perimeter method for Circle
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func main() {
	// Create instances of Rectangle and Circle
	rect := Rectangle{Width: 4, Height: 3}
	circle := Circle{Radius: 2}

	// Call methods on the instances
	printShapeInfo(rect)
	printShapeInfo(circle)
}

func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

