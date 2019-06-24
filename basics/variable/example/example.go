// Sample program to show how to declare variables.
package main

import "fmt"

func main() {

	// Declare variables that are set to their zero value.
	var a bool
	var b int
	var c float32
	var d complex64
	var e string
	var f *int

	fmt.Printf("var a \t %T \t\t [%v]\n", a, a)
	fmt.Printf("var b \t %T \t\t [%v]\n", b, b)
	fmt.Printf("var c \t %T \t [%v]\n", c, c)
	fmt.Printf("var d \t %T \t [%v]\n", d, d)
	fmt.Printf("var e \t %T \t [%v]\n", e, e)
	fmt.Printf("var f \t %T \t\t [%v]\n\n", f, f)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	aa := 10
	bb := "hello"
	cc := 3.14159
	dd := true

	fmt.Printf("aa := 10 \t\t %T \t\t [%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T \t [%v]\n", bb, bb)
	fmt.Printf("cc := 3.14159 \t %T \t [%v]\n", cc, cc)
	fmt.Printf("dd := true \t\t %T \t\t [%v]\n\n", dd, dd)

	// Specify type and perform a conversion
	aaa := int32(10)

	fmt.Printf("aaa := int32(10) %T\t\t [%v]\n", aaa, aaa)
}
