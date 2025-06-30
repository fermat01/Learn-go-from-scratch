package main

import (
	"fmt"
)

// interface in Go is a type that specifies a contract or a set of methods that a type must implement.
// It allows you to define behavior without specifying the concrete type.	
// An interface is defined using the `type` keyword followed by the interface name and its methods.



type Shape interface {
	Area() float64 // Method to calculate the area of the shape
	Perimeter() float64 // Method to calculate the perimeter of the shape
}
// Circle is a struct that represents a circle with a radius.
type Circle struct {
	Radius float64 // Radius of the circle
}
// Rectangle is a struct that represents a rectangle with a width and height.
type Rectangle struct {
	Width  float64 // Width of the rectangle
	Height float64 // Height of the rectangle
}
// Area method calculates the area of the circle.
// It is defined with a receiver of type `Circle`, allowing it to access the `Radius
// field of the circle.
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius // Area = π * r^2
}
// Perimeter method calculates the perimeter of the circle.
// It is also defined with a receiver of type `Circle`, allowing it to access the `

// Radius` field of the circle.
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius // Perimeter = 2 * π * r
}
// Area method calculates the area of the rectangle.
// It is defined with a receiver of type `Rectangle`, allowing it to access the `Width
// and `Height` fields of the rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height // Area = width * height
}
// Perimeter method calculates the perimeter of the rectangle.
// It is also defined with a receiver of type `Rectangle`, allowing it to access the `
// Width` and `Height` fields of the rectangle.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height) // Perimeter = 2 * (width + height)
}		

// main function demonstrates the use of interfaces in Go.
// It creates instances of `Circle` and `Rectangle`, stores them in a slice of `Shape`,
// and iterates over the shapes to print their area and perimeter.
// It also shows how to use type assertions to check the concrete type of the shape.
// The main function is the entry point of the program.
// It is where the program starts executing.
func main() {
	// Create a circle and a rectangle
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}

	// Create a slice of Shape interface to hold different shapes
	shapes := []Shape{circle, rectangle}

	// Iterate over the shapes and print their area and perimeter
	for _, shape := range shapes {
		fmt.Printf("Area: %.2f, Perimeter: %.2f\n", shape.Area(), shape.Perimeter())
	}

	// You can also use type assertions to check the concrete type of the shape
	for _, shape := range shapes {
		switch s := shape.(type) {
		case Circle:
			fmt.Printf("Circle with radius %.2f\n", s.Radius)
		case Rectangle:
			fmt.Printf("Rectangle with width %.2f and height %.2f\n", s.Width, s.Height)
		default:
			fmt.Println("Unknown shape")
		}
	}	
}



