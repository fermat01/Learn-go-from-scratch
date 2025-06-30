package main

import "fmt"

func main() {
	// Example values for format verbs
	str := "hello"
	str2 := "world"
	num := 42
	pi := 3.14
	piPrecise := 3.14159
	flag := true
	char := 'A'
	hexNum := 255
	x := 10 // For pointer example
	binaryNum := 5
	octalNum := 8
	sciNum := 12345678
	floatNum := 123456.78
     message := "Hello" 

	// %v - Default format
	fmt.Printf("%%v: %v\n", str) // Output: %v: hello

	fmt.Printf("|%10s|\n", message) // Output: |   Hello|
	fmt.Printf("|%-10s|\n", message) // Output: |Hello   |
	// %s - String
	fmt.Printf("%%s: %s\n", str2) // Output: %s: world

	// %d - Integer (base 10)
	fmt.Printf("%%d: %d\n", num) // Output: %d: 42

	// %f - Float
	fmt.Printf("%%f: %f\n", pi) // Output: %f: 3.140000

	// %.2f - Float with 2 decimal places
	fmt.Printf("%%.2f: %.2f\n", piPrecise) // Output: %.2f: 3.14

	// %t - Boolean
	fmt.Printf("%%t: %t\n", flag) // Output: %t: true

	// %c - Character (rune)
	fmt.Printf("%%c: %c\n", char) // Output: %c: A

	// %x - Hex (lowercase)
	fmt.Printf("%%x: %x\n", hexNum) // Output: %x: ff

	// %X - Hex (uppercase)
	fmt.Printf("%%X: %X\n", hexNum) // Output: %X: FF

	// %p - Pointer address
	fmt.Printf("%%p: %p\n", &x) // Output: %p: 0x... (address will vary)

	// %q - Quoted string
	fmt.Printf("%%q: %q\n", "hi") // Output: %q: "hi"

	// %b - Binary
	fmt.Printf("%%b: %b\n", binaryNum) // Output: %b: 101

	// %o - Octal
	fmt.Printf("%%o: %o\n", octalNum) // Output: %o: 10

	// %e - Scientific notation
	fmt.Printf("%%e: %e\n", float64(sciNum)) // Output: %e: 1.234568e+07

	// %E - Scientific (uppercase)
	fmt.Printf("%%E: %E\n", float64(sciNum)) // Output: %E: 1.234568E+07

	// %g - Float, exponent as needed
	fmt.Printf("%%g: %g\n", floatNum) // Output: %g: 123456.78

	// %G - Float, exponent as needed
	fmt.Printf("%%G: %G\n", floatNum) // Output: %G: 123456.78
}
