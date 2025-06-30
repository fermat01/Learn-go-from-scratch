package main

import "fmt"



func main() {
	// An if statement in Go is used to execute a block of code conditionally.
	// The syntax is similar to C-style if statements.
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is not greater than 5")
	}

	// You can also use an if statement with an initialization statement.
	if y := 20; y < x {
		fmt.Println("y is less than x")
	} else {
		fmt.Println("y is not less than x")
	}
}