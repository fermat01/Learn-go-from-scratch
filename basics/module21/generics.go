package main	
import (
	"fmt"
)

// Generics in Go allow you to write functions and data structures that can work with any data type.
// This is achieved using type parameters, which are specified in square brackets after the function name.	
// Generic function to sum the elements of a slice
// A generic function that can sum either ints or floats
func Sum[T int | float64](nums []T) T {
    var total T
    for _, num := range nums {
        total += num
    }
    return total
}

// Generic function to reverse a slice
func Reverse[T any](s []T) []T {
	n := len(s)
	reversed := make([]T, n)
	for i := 0; i < n; i++ {
		reversed[i] = s[n-1-i]
	}
	return reversed
}
// Generic function to filter elements in a slice based on a predicate function
func Filter[T any](s []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range s {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}
// Generic function to map elements in a slice using a transformation function
func Map[T any, U any](s []T, transform func(T) U) []U {
	var result []U
	for _, v := range s {
		result = append(result, transform(v))
	}
	return result
}

// Generic function to check if a slice contains a specific element	
func Contains[T comparable](s []T, element T) bool {
	for _, v := range s {
		if v == element {
			return true
		}
	}
	return false
}
// Generic function to concatenate two slices
func Concat[T any](s1, s2 []T) []T {
	result := make([]T, len(s1)+len(s2))
	copy(result, s1) // Copy elements from the first slice
	copy(result[len(s1):], s2) // Copy elements from the second slice
	return result
}
//  min function demonstrates the use of generic functions in Go.
func main() {

	// Create a slice of integers
	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Integer Slice:")
	fmt.Println("Sum:", Sum(intSlice))
	fmt.Println("Reversed:", Reverse(intSlice))
	
	// Create a slice of strings
	stringSlice := []string{"apple", "banana", "cherry"}
	fmt.Println("\nString Slice:")
	
	// Filter strings that contain the letter 'a'
	filtered := Filter(stringSlice, func(s string) bool {
		return len(s) > 5
	})
	fmt.Println("Filtered (length > 5):", filtered)
	
	// Map strings to their lengths
	mapped := Map(stringSlice, func(s string) int {
		return len(s)
	})
	fmt.Println("Mapped (lengths):", mapped)
	// Check if the integer slice contains a specific element
	fmt.Println("Contains 3?", Contains(intSlice, 3))
	
	// Concatenate two slices
	concatResult := Concat(intSlice, []int{6, 7, 8})
	fmt.Println("Concatenated Slice:", concatResult)
}