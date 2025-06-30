package main

// This is a simple Go program that prints "hello world" to the console.

import (
	"fmt"
	"net/http"
)
// The main function is the entry point of the program.
// It sets up an HTTP server that listens on port 8080 and responds with "Hello, World!" to any request made to the root URL ("/").
func main() {
    fmt.Println("hello world")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, coder2f!")
    })
    http.ListenAndServe(":8080", nil)
}