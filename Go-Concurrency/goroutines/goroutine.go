package main

import "fmt"

func main() {
	go printGreetings("goroutine")
	printGreetings("main function")
}

func printGreetings(source string) {
	for i := 0; i < 9; i++ {
		fmt.Println("Hello Gopher!", i, source)
	}
}
