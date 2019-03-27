package main

import "fmt"

var done = make(chan bool)

func main() {
	go printGreetings("goroutine")
	printGreetings("main function")
	<-done
}

func printGreetings(source string) {
	for i := 0; i < 9; i++ {
		fmt.Println("Hello Gopher!", i, source)
	}

	if source == "goroutine" {
		done <- true
	}
}
