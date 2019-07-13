// Sample program to show how to declare and initialize struct types.
package main

import "fmt"

// example represents a type with different fields.
type example struct {
	flag    bool    // 1 byte
	counter int16   // 2 bytes
	pi      float32 // 4 bytes
}

func main() {

	// Declare a variable of type example set to its zero value.
	var e1 example

	// Display the value.
	fmt.Printf("%+v\n", e1)

	// Declare a variable of type example and init using a struct literal
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Display the field values
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
}
