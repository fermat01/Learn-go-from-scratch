package main
import (
	"fmt"
)
// slices are a flexible and powerful way to work with sequences of data in Go.
func main() {
	// Creating a slice using a slice literal
	fruits := []string{"apple", "banana", "cherry"}
	fmt.Println("Fruits:", fruits)

	// Creating a slice using the make function
	numbers := make([]int, 5) // creates a slice of int with length 5 and capacity 5
	fmt.Println("Numbers:", numbers)
	// Accessing elements in a slice
	fmt.Println("First fruit:", fruits[0])
	fmt.Println("Second fruit:", fruits[1])

	// Modifying an element in a slice
	fruits[1] = "blueberry"
	fmt.Println("Modified fruits:", fruits)

	// Appending to a slice
	fruits = append(fruits, "date")
	fmt.Println("After appending:", fruits)

	// Slicing a slice
	subFruits := fruits[1:3]
	fmt.Println("Sub-slice:", subFruits)
	// Length and capacity of a slice
	fmt.Println("Length of fruits:", len(fruits))
	fmt.Println("Capacity of fruits:", cap(fruits))	
	// Iterating over a slice using a for loop
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Fruit at index %d: %s\n", i, fruits[i])
	}
	// Iterating over a slice using a range loop
	for index, value := range fruits {
		fmt.Printf("Fruit at index %d: %s\n", index, value)
	}
	// Slices can be multidimensional, such as a 2D slice.
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2D Slice (Matrix):")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println() // New line after each row
	}
	// Output:
	// 2D Slice (Matrix):
	// 1 2 3
	// 4 5 6
	// Note: Slices are more flexible than arrays in Go. They can grow and shrink
	// dynamically, and they are built on top of arrays. When you append to a slice
	// and the underlying array is full, Go automatically allocates a new array with
	// more capacity and copies the elements to the new array. This allows slices to
	// be more efficient and easier to work with than fixed-size arrays.
	// Slices are reference types, meaning when you assign a slice to another variable,
	// both variables refer to the same underlying array. Modifying one slice will
	// affect the other. If you want to create a copy of a slice, you can
	// use the `copy` function.
	copyOfFruits := make([]string, len(fruits))
	copy(copyOfFruits, fruits) // Copying the slice
	copyOfFruits[0] = "kiwi" // Modifying the copied slice
	fmt.Println("Original fruits after copying:", fruits) // Output: [blueberry cherry date]
	fmt.Println("Copied fruits after modification:", copyOfFruits) // Output: [kiwi blueberry cherry date]
	// This shows that modifying `copyOfFruits` does not affect `fruits`, as they are separate slices.
	// Slices can also be used to create dynamic data structures like stacks and queues.
	// For example, you can use a slice to implement a stack by using the `append` function to push elements onto the stack
	// and the `pop` operation by slicing the slice to remove the last element.
	stack := []int{1, 2, 3}
	fmt.Println("Stack before pop:", stack)
	stack = append(stack, 4) // Push 4 onto the stack
	fmt.Println("Stack after push:", stack)
	stack = stack[:len(stack)-1] // Pop the last element
	fmt.Println("Stack after pop:", stack) // Output: [1 2 3]
	// Slices can also be used to implement queues by using the `append` function to
	// add elements to the end of the slice and slicing the slice to remove elements from the
	// front of the slice.
	queue := []int{1, 2, 3}
	fmt.Println("Queue before dequeue:", queue)
	queue = append(queue, 4) // Enqueue 4
	fmt.Println("Queue after enqueue:", queue)
	queue = queue[1:] // Dequeue the first element
	fmt.Println("Queue after dequeue:", queue) // Output: [2 3 4]
	// Slices are a powerful feature in Go that allows you to work with sequences of data
										
}		