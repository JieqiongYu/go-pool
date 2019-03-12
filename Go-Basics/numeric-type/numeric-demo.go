package main

import "fmt"

func main() {

	// Declare a 32 bit integer without initializing it
	// Remember: variables declared without an explicit initial value are given their zero value
	var myInteger int32


	fmt.Println("Value of myInteger: ", myInteger)

	// Declare and Initialize 2 float64's and thens how their sum
	var myFloatingPointNumber float64 = 36.333
	var myOtherFloatingPointNumber float64 = 306.96969696
	fmt.Println("Sum of my floating point numebrs: ", myFloatingPointNumber+myOtherFloatingPointNumber)

	// Go will figure out what the type is
	x, y, z := 0, 1, 2
	fmt.Printf("x: %d\ty: %d\t z: %d\t", x, y, z)

	// Example of a complex number
	myComplexNumber := 5 + 24i
	fmt.Println("Value of myComplexNumber: ", myComplexNumber)

	// Example of grouping constant declaration/initializations
	const(
		sppedOfLight 		= 299792458 // unit: m/s
		pi 					= 3.14
		myFavoriteNumber 	= 108
	)

	// Example of grouping variable declarations/initializations
	var (
		a int 	= 0 		// an integer
		b 		= 1.8 + 3i // a complex number
		c 		= 2.7 		// a floating-point number
	)
	fmt.Printf("a: %v\tb: %v\t c: %v\t", a, b, c)
}
