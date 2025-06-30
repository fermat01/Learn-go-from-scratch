package main
import (
	"fmt"
)
// Pointers in Go
// Pointers are variables that store the memory address of another variable.
// They allow you to reference and manipulate the original variable directly, rather than working with a copy.
// This can be useful for performance optimization and when you want to modify the original variable from a function.
func main() {
	// Declare an integer variable
	var x int = 10
	fmt.Println("Value of x:", x) // Output: 10

	// Declare a pointer to the integer variable
	var p *int = &x
	fmt.Println("Address of x:", p) // Output: Address of x in memory

	// Dereference the pointer to get the value it points to
	fmt.Println("Value pointed by p:", *p) // Output: 10

	// Modify the value using the pointer
	*p = 20
	fmt.Println("New value of x:", x) // Output: 20

	// Call modifyValue to change x via pointer
	modifyValue(&x)
	fmt.Println("Value of x after modification:", x) // Output: 30
}
// Pointers can also be used with functions to modify the original variable
func modifyValue(val *int) {
	*val = 30 // Modify the value at the address pointed by val
	
	fmt.Println("Value modified inside function:", *val) // Output: 30
	
}
	
