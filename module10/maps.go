package main


import (
	"fmt"
)

//  Maps in Go are a built-in data type that associates keys with values.
// They are similar to dictionaries in Python or hash tables in other languages.		

func main() {
	// Declare and initialize a map
	// The map is of type map[string]int, meaning it maps strings to integers.
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}

	// Accessing a value in the map using its key
	fmt.Println("Alice's age:", ages["Alice"])

	// Adding a new key-value pair to the map
	ages["Charlie"] = 35
	fmt.Println("Charlie's age:", ages["Charlie"])

	// Iterating over the map
	for name, age := range ages {
		fmt.Printf("%s is %d years old.\n", name, age)
	}

	// Checking if a key exists in the map
	if age, exists := ages["David"]; exists {
		fmt.Println("David's age:", age)
	} else {
		fmt.Println("David not found in the map.")
	}
}
