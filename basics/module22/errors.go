package main
import (
	"fmt"
	"errors"
)
// Custom error types in Go allow you to define specific error conditions and provide additional context.	
// Here, we define a custom error type named `DivisionError` that implements the `error` interface.
func sqrt (x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("cannot calculate square root of a negative number")
	}
	return x * x, nil
}
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}
func main() {
	// Example of using the custom error type
	result, err := sqrt(-4)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Square root:", result)
	}

	// Example of using the divide function with error handling
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division result:", result)
	}
}	