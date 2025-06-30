package main

import "fmt"

// Custom error type that implements the error interface
// This type can be used to provide more context about an error condition.
// It includes fields for an error code and a message.
type MyError struct {
    Code    int
    Message string
}

// Implement the Error() method to satisfy the error interface
func (e *MyError) Error() string {
    return fmt.Sprintf("Code %d: %s", e.Code, e.Message)
}

func doSomething() error {
    // Simulate an error condition
    return &MyError{
        Code:    404,
        Message: "Resource not found",
    }
}

func main() {
    err := doSomething()
    if err != nil {
        // Print the error message
        fmt.Println("Error:", err)

        // Check if the error is of type MyError
        if myErr, ok := err.(*MyError); ok {
            fmt.Printf("Custom error details: Code %d, Message: %s\n", myErr.Code, myErr.Message)
        }
    }
}
