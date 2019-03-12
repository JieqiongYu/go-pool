package main

import (
	"fmt"
	"github.com/go-pool/Go-Basics/greeting-package"
)

func main() {
	greeting_package.PrintGreetings()
	greeting_package.GopherGreetings()

	fmt.Println("The value of the Magic Number is:", greeting_package.MagicNumber)
}
