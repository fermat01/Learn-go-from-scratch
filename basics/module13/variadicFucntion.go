package main	
import (
	"fmt"
)	

// variadic functions
func main() {
	// Define a variadic function that takes a variable number of integers
	sum := func(numbers ...int) int {
		total := 0
		for _, number := range numbers {
			total += number
		}
		return total
	}

	// Call the variadic function with different numbers of arguments
	fmt.Println("Sum of 1, 2, 3:", sum(1, 2, 3))
	fmt.Println("Sum of 4, 5, 6, 7:", sum(4, 5, 6, 7))
	fmt.Println("Sum of no arguments:", sum())
	
	// Define a variadic function that takes a variable number of strings
	printNames := func(names ...string) {
		for _, name := range names {
			fmt.Println("Hey hi", name)
		}
	}

	// Call the variadic function with different numbers of arguments
	printNames("Alice", "Bob", "Charlie","Brown")
	printNames("Dave")
}

