package main	
import (
	"fmt"
)

func main() {
	// Example usage of recursive functions
	fmt.Println("Factorial of 5:", factorial(5)) // Output: 120
	fmt.Println("Fibonacci of 6:", fibonacci(6)) // Output: 8
	fmt.Println("Sum of digits of 1234:", sumOfDigits(1234)) // Output: 10
	fmt.Println("GCD of 48 and 18:", gcd(48, 18)) // Output: 6
}
// This code demonstrates the use of recursive functions in Go.
// It includes functions to calculate factorial, Fibonacci numbers, sum of digits, and greatest common divisor
// using recursion. Each function has a base case to terminate the recursion and a recursive case to
// call itself with modified parameters until the base case is reached.
// The main function provides examples of how to call these recursive functions and print their results.
// recursion concept in Go
// Factorial function using recursion
func factorial(n int) int {
	if n == 0 {
		return 1 // Base case: factorial of 0 is 1
	}
	return n * factorial(n-1) // Recursive case
}
// Fibonacci function using recursion
func fibonacci(n int) int {
	if n <= 1 {
		return n // Base case: fibonacci(0) = 0, fibonacci(1) = 1
	}
	return fibonacci(n-1) + fibonacci(n-2) // Recursive case
}
// Function to calculate the sum of digits of a number using recursion
func sumOfDigits(n int) int {
	if n == 0 {
		return 0 // Base case: sum of digits of 0 is 0
	}
	return n%10 + sumOfDigits(n/10) // Recursive case: last digit + sum of remaining digits
}
// Function to calculate the greatest common divisor (GCD) using recursion
func gcd(a, b int) int {
	if b == 0 {
		return a // Base case: GCD(a, 0) = a
	}
	return gcd(b, a%b) // Recursive case: GCD(b, a % b)
}