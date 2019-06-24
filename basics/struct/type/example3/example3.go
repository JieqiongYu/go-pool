// Sample program to show how variables of an unnamed type can
// be assigned to variables of a named type, when they are identical.
// And it also shows how 2 named type can be assigned in intent way.
package main

import "fmt"

type bill struct {
	flag    bool
	counter int16
	pi      float32
}

type alice struct {
	flag    bool
	counter int16
	pi      float32
}

// example represents a type with different fields.
type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {

	// Show how 2 named type can be assigned in intent way.
	var b bill
	var a alice
	// b = a is not allowed, using the following way to achieve this
	b = bill(a) // tell the compiler the intent explicitly
	fmt.Println(b, a)
	fmt.Println()

	// Declare a variable of an anonymous type and init
	// using a struct literal.
	e := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Create a value of type example.
	var ex example

	// Assign the value of the unnamed struct type
	// to the named struct type value.
	ex = e

	// Display the values.
	fmt.Printf("%+v\n", ex)
	fmt.Printf("%+v\n\n", e)
	fmt.Println("Flag", e.flag)
	fmt.Println("Counter", e.counter)
	fmt.Println("Pi", e.pi)
}
