package main
import (
	"fmt"
)
// Arrays in Go are fixed-size sequences of elements of the same type.
// They are defined with a specific length and type, and the size cannot be changed after declaration.
func main() {
	// Declare an array of integers with a fixed size of 5
	var arr [5]int
	// Initialize the array with zero values
	// In Go, arrays are automatically initialized with zero values for their type.
	// For integers, the zero value is 0.
	// The array will have 5 elements, all initialized to 0.
	// This is done implicitly when the array is declared.
	// You can also initialize the array with specific values at the time of declaration.
	// For example: var arr = [5]int{1, 2, 3, 4, 5}
	// But here we are initializing it with zero values.
	// Print the initial state of the array
	fmt.Println("Initial array:", arr) // Output: [0 0 0 0 0]

	// Assign values to the array
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	fmt.Println("Array after assignment:", arr) // Output: [1 2 3 4 5]

	// Access elements in the array
	fmt.Println("First element:", arr[0]) // Output: 1
	fmt.Println("Second element:", arr[1]) // Output: 2

	// Length of the array
	fmt.Println("Length of the array:", len(arr)) // Output: 5

	// Iterate over the array using a for loop
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Element at index %d: %d\n", i, arr[i])
	}
	// Iterate over the array using a range loop
	for index, value := range arr {
		fmt.Printf("Element at index %d: %d\n", index, value)
	}
	// Arrays are value types in Go, meaning when you assign an array to another variable,
	// a copy of the array is made.
	arr2 := arr // Copying the array
	arr2[0] = 10 // Modifying the copied array
	fmt.Println("Original array after copying:", arr) // Output: [1 2 3 4 5]
	fmt.Println("Copied array after modification:", arr2) // Output: [	10 2 3 4 5]
	// This shows that modifying arr2 does not affect arr, as they are separate copies.
	// Arrays can also be multidimensional, such as a 2D array.
	var matrix [2][3]int // A 2D array with 2 rows and 3 columns
	// Initialize the 2D array
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6
	// Print the 2D array
	fmt.Println("2D Array (Matrix):")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println() // New line after each row
	}
	// Output:
	// 2D Array (Matrix):
	// 1 2 3
	// 4 5 6
	// Note: In Go, arrays are fixed in size, and their size is part of the type.
	// This means that [5]int and [10]int are considered different types.
	// If you need a dynamically sized array, you can use slices, which are more flexible
	// and can grow or shrink in size. Slices are built on top of arrays and
	// provide a more convenient way to work with sequences of elements.
	// Slices are a more flexible alternative to arrays in Go.
}