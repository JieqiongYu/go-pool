// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

import "fmt"

// user represents a user in the system.
type user struct {
	name  string
	email string
	age   int
}

func main() {

	// Declare variable of type user and init using a struct literal.
	mer := user{
		name:  "Mer",
		email: "mer@mer.com",
		age:   3,
	}

	// Display the field values.
	fmt.Println("Name", mer.name)
	fmt.Println("Email", mer.email)
	fmt.Println("Age", mer.age)

	fmt.Println()

	// Declare a variable using an anonymous struct.
	jq := struct {
		name  string
		email string
		age   int
	}{
		name:  "JQ",
		email: "jq@jq.com",
		age:   4,
	}

	// Display the field values.
	fmt.Println("Name", jq.name)
	fmt.Println("Email", jq.email)
	fmt.Println("Age", jq.age)
}
