package main

import (
	"fmt"
	"time"
)


func main() {
	// A switch statement in Go is used to execute one of several possible blocks of code based on the value of an expression.
	// It is similar to a series of if-else statements, but can be more concise and easier to read.
	// The switch statement evaluates an expression and compares it against multiple cases.
	// If a case matches, the corresponding block of code is executed.

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}