package main

import (
	"fmt"
)

// Function that returns multiple values
 func main() {		
	// Define a function that returns multiple values
	multiplyAndDivide := func(a, b int) (int, float64) {
		p := a * b
		q := float64(a) / float64(b)
		// Return the product and quotient
		return  p,q
	}
	// Call the function and capture the returned values
	product, quotient := multiplyAndDivide(10, 2)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)			
	// Define a function that returns multiple values with named return values
	multiplyAndDivideNamed := func(a, b int) (product int, quotient float64) {
		product = a * b
		quotient = float64(a) / float64(b)
		return // Named return values can be returned without specifying them again
	}
	// Call the function and capture the returned values
	productNamed, quotientNamed := multiplyAndDivideNamed(20, 4)
	fmt.Println("Product (named):", productNamed)
	fmt.Println("Quotient (named):", quotientNamed)
	// Define a function that returns multiple values of different types					

	mixedReturn := func(a, b int) (int, string, bool) {
		return a + b, "Result", true
	}
	// Call the function and capture the returned values
	sum, message, success := mixedReturn(5, 10)
	fmt.Println("Sum:", sum)
	fmt.Println("Message:", message)
	fmt.Println("Success:", success)
	// Define a function that returns multiple values in a struct
	type Result struct {
		Product  int
		Quotient float64
	}
	resultStruct := func(a, b int) Result {
		return Result{
			Product:  a * b,
			Quotient: float64(a) / float64(b),
		}
	}
	// Call the function and capture the returned struct
	result := resultStruct(15, 3)
	fmt.Println("Product (struct):", result.Product)
	fmt.Println("Quotient (struct):", result.Quotient)	
	// Define a function that returns multiple values in a slice
	sliceReturn := func(a, b int) []interface{} {
		return []interface{}{a + b, "Sum", true}
	}
	// Call the function and capture the returned slice
	sliceResult := sliceReturn(8, 12)
	fmt.Println("Slice Result:", sliceResult)
	fmt.Println("Sum from slice:", sliceResult[0])
	fmt.Println("Message from slice:", sliceResult[1])
 }				