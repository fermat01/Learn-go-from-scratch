
package main

import "fmt"

func main() {

    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2 // multiple variable declaration
	// b and c are both of type int
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "apple"
    fmt.Println(f)
    /* An overflow operation in Go occurs when an arithmetic operation 
    (such as addition, multiplication, or subtraction) produces a value 
    that exceeds the maximum or minimum value that can be stored by the variable’s 
     data type. For example, adding 1 to the maximum value of a byte (which is 255 for an unsigned byte) 
     results in overflow, causing the value to "wrap around" to 0 for unsigned integers 
    */
    // Example of overflow
    var g byte = 255 // maximum value for byte
    fmt.Println(g)
    g = g + 1 // this will cause overflow
    fmt.Println(g) // g will wrap around to 0
}