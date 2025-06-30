package main
import "fmt"
func main() {

	// A for loop in Go can be used to iterate over a range of values.
	// The syntax is similar to C-style for loops.
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// You can also use a for loop to iterate over a slice or array.
	fruits := []string{"apple", "banana", "cherry"}
	for index, fruit := range fruits {
		fmt.Printf("Index: %d, Fruit: %s\n", index, fruit)
	}

	// A while-like loop can be created using a for loop with no initialization or post statement.
	j := 0
	for j < 3 {
		fmt.Println("While-like loop iteration:", j)
		j++
	}
}
