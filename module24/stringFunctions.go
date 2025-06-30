package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// String functions in Go provide a variety of operations for manipulating and analyzing strings.
// The `strings` package in Go provides many useful functions for working with strings.
// Here are some common string functions and their usage:
func main() {
	// Example string
	str := "Hello, World!"
 
	num := 42
	fmt.Printf("Number: %d\n", num)
	str3 := strconv.Itoa(num)
	fmt.Printf("Converted to string: %s\n", str3)
	// Length of the string
	length := len(str)
	fmt.Printf("Length of the string: %d\n", length)

	// Convert to uppercase
	upper := strings.ToUpper(str)
	fmt.Printf("Uppercase: %s\n", upper)

	// Convert to lowercase
	lower := strings.ToLower(str)
	fmt.Printf("Lowercase: %s\n", lower)

	// Check if the string contains a substring
	contains := strings.Contains(str, "World")
	fmt.Printf("Contains 'World': %t\n", contains)

	// Split the string into a slice of substrings
	split := strings.Split(str, ", ")
	fmt.Printf("Split: %v\n", split)

	// Join a slice of strings into a single string
	joined := strings.Join(split, " - ")
	fmt.Printf("Joined: %s\n", joined)
	// Replace a substring with another substring
	replaced := strings.ReplaceAll(str, "World", "Go")
	fmt.Printf("Replaced: %s\n", replaced)
	// Trim whitespace from the string
	trimmed := strings.TrimSpace("   Hello, Go!   ")
	fmt.Printf("Trimmed: '%s'\n", trimmed)
	// Check if the string starts with a prefix
	startsWith := strings.HasPrefix(str, "Hello")
	fmt.Printf("Starts with 'Hello': %t\n", startsWith)
	// Check if the string ends with a suffix
	endsWith := strings.HasSuffix(str, "!")
	fmt.Printf("Ends with '!': %t\n", endsWith)	
	// Count occurrences of a substring
	count := strings.Count(str, "o")
	fmt.Printf("Count of 'o': %d\n", count)
	// Find the index of a substring
	index := strings.Index(str, "World")
	fmt.Printf("Index of 'World': %d\n", index)
	// Find the last index of a substring
	lastIndex := strings.LastIndex(str, "o")
	fmt.Printf("Last index of 'o': %d\n", lastIndex)
	// Repeat the string multiple times
	repeated := strings.Repeat(str, 2)
	fmt.Printf("Repeated: %s\n", repeated)
	// Check if the string is empty
	isEmpty := str == ""
	fmt.Printf("Is empty: %t\n", isEmpty)
	// Convert string to rune slice (for Unicode characters)
	runes := []rune(str)
	fmt.Printf("Runes: %v\n", runes)
	// Convert rune slice back to string
	strFromRunes := string(runes)
	fmt.Printf("String from runes: %s\n", strFromRunes)
	// Check if the string is a valid UTF-8 encoded string
	isValidUTF8 := utf8.ValidString(str)
	fmt.Printf("Is valid UTF-8: %t\n", isValidUTF8)
}	
