package main

import (
	"fmt"
)	

// functions in Go are defined using the `func` keyword.
// Functions can take parameters and return values.
func main() {
//	// Define a simple function that adds two integers
	add := func(a int, b int) int {
		return a + b
	}	
	// Call the function and print the result
	result := add(5, 10)
	fmt.Println("The sum is:", result)
	// Define a function that takes a variable number of arguments
	sum := func(numbers ...int) int {
		total := 0
		for _, number := range numbers {
			total += number
		}
		return total
	}
	// Call the function with different numbers of arguments
	fmt.Println("Sum of 1, 2, 3:", sum(1, 2, 3))
	fmt.Println("Sum of 4, 5, 6, 7:", sum(4, 5, 6, 7))
	// Define a function that returns multiple values
	multiplyAndDivide := func(a, b int) (int, float64) {
		return a * b, float64(a) / float64(b)
	}
	// Call the function and capture the returned values
	product, quotient := multiplyAndDivide(10, 2)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)
	// Define a function that takes a function as an argument
	execute := func(fn func(int, int) int, x, y int) int {
		return fn(x, y)
	}
	// Call the function with the add function as an argument
	result = execute(add, 3, 4)
	fmt.Println("Result of execute with add:", result)
	// Define a function that returns another function
	makeMultiplier := func(factor int) func(int) int {
		return func(x int) int {
			return x * factor
		}
	}
	// Create a multiplier function that multiplies by 3
	multiplier := makeMultiplier(3)
	// Call the multiplier function
	fmt.Println("3 times 5 is:", multiplier(5))
	// Define a function that uses a closure
	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}
	// Create a new counter function
	countFunc := counter()
	// Call the counter function multiple times
	 fmt.Println("Counter:", countFunc())
	 fmt.Println("Counter:", countFunc())
	 fmt.Println("Counter:", countFunc())
	
	q, r := divide(10, 3)
	fmt.Println("Quotient:", q, "Remainder:", r)
}

func divide(a int, b int) (int, int) {
	// divide function takes two integers and returns their quotient and remainder
	if b == 0 {
		panic("division by zero")
	}
	quotient := a / b
	remainder := a % b
	// Return both quotient and remainder
	return quotient, remainder
}