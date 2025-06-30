package main
import (
	"fmt"
)
// structs in Go are a way to group related data together. They are similar to classes in other object-oriented languages, but they do not support inheritance. Instead, structs can be composed of other structs or basic types.
// A struct is defined using the `type` keyword followed by the struct name and its fields.
// Here, we define a struct named `Person` with two fields: `Name` of type `string` and `Age` of type `int`.

type Person struct {
	Name string // Name of the person
	Age  int    // Age of the person
}


func main() {
	// Declare and initialize a struct
	// The struct is of type Person, which has two fields: Name and Age.
	person := Person{
		Name: "Alice",
		Age:  30,
	}

	// Accessing fields of the struct
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)

	// Modifying fields of the struct
	person.Age = 31
	// using sprintf to format the string
	fmt.Println("Updated Age:", person.Age)
	s:= fmt.Sprintf("%s is %d years old.", person.Name, person.Age)
	fmt.Println(s)

	

}	
