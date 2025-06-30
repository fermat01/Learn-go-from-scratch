package main
import (
	"fmt"
	 
)	

// Strings and Runes in Go
// Strings are a sequence of characters, while runes are a representation of Unicode code points.
// Runes are used to handle characters that may not fit into a single byte, such as
// characters from non-Latin scripts or special symbols.
func main() {
	// Declare and initialize a string
	str := "Hello, 世界" // Contains both ASCII and non-ASCII characters
	fmt.Println("String:", str)

	// Iterate over the string using range to get runes
	for i, char := range str {
		fmt.Printf("Character at index %d: %c (Unicode: %U)\n", i, char, char)
	}

	// Convert string to rune slice
	runes := []rune(str)
	fmt.Println("Rune slice:", runes)

	// Accessing individual runes
	fmt.Println("First rune:", string(runes[0])) // Output: H
	fmt.Println("Second rune:", string(runes[1])) // Output: e
}
//// This code demonstrates how to work with strings and runes in Go.
// It shows how to declare a string, iterate over it to access individual characters (runes),
// convert a string to a slice of runes, and access individual runes.
// The `range` keyword is used to iterate over the string, which allows us to handle both ASCII and non-ASCII characters correctly.
// The output includes the index of each character, its Unicode representation, and the rune slice.
// This is useful for handling text that includes characters from various languages or special symbols,	