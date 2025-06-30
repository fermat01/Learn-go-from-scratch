package main

import (
	"fmt"
)
// Methods in Go are functions that are associated with a specific type, allowing you to define behavior for that type.
// They are similar to methods in object-oriented programming languages, but Go does not have classes. Instead, you define methods on structs or other types.
// A method is defined by specifying a receiver type before the function name.
// The receiver is a special parameter that allows the method to access the fields of the struct or type it is associated with.
// Here, we define a struct named `Rectangle` with two fields: `Width` and `Height`.	

type Rectangle struct {
	Width  float64 // Width of the rectangle
	Height float64 // Height of the rectangle
}
// The `Area` method calculates the area of the rectangle.
// It is defined with a receiver of type `Rectangle`, allowing it to access the `Width

// and `Height` fields of the rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// The `Perimeter` method calculates the perimeter of the rectangle.
// It is also defined with a receiver of type `Rectangle`, allowing it to access the `
// `Width` and `Height` fields of the rectangle.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
// The `Scale` method scales the rectangle by a given factor.
// It modifies the `Width` and `Height` fields of the rectangle by multiplying them by
// the provided factor. Note that this method does not return a new rectangle; it modifies
// the existing rectangle.
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}



func main() {
	// Create a new rectangle with width 5 and height 10
	rect := Rectangle{Width: 5, Height: 10}

	// Calculate the area of the rectangle using the Area method
	area := rect.Area()
	fmt.Printf("Area of rectangle: %.2f\n", area)

	// Calculate the perimeter of the rectangle using the Perimeter method
	perimeter := rect.Perimeter()
	fmt.Printf("Perimeter of rectangle: %.2f\n", perimeter)

	// Scale the rectangle by a factor of 2 using the Scale method
	rect.Scale(2)
	fmt.Printf("Scaled rectangle: Width=%.2f, Height=%.2f\n", rect.Width, rect.Height)
}
