package main

import (
	"flag"
	"fmt"
)

func main() {
	var gopherName string
	// &gopherName is the address of the variable gopherName
	// "gophoername" is the name of the flag
	// "Gopher" is the default value of "gophername" flag
	// "The name of the Gohper" to document what the "gophername" flag is for
	flag.StringVar(&gopherName, "gophername", "Gopher", "The name of the Gohper")
	flag.Parse()
	fmt.Println("Hello " + gopherName + "!")
}
