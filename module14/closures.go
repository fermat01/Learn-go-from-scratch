package main

import "fmt"

func fibonacci() func() int {
    a, b := 0, 1
    return func() int {
        a, b = b, a+b
        return a
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
// This code defines a closure that generates Fibonacci numbers.
// The `fibonacci` function returns a function that, when called, returns the next Fibonacci number.
// The closure captures the variables `a` and `b`, allowing the inner function to maintain state across calls.
// In the `main` function, we call the closure multiple times to print the first 10 Fibonacci numbers.	