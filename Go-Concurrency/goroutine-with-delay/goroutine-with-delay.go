package main

import (
	"fmt"
	"time"
)

func main() {
	go printGreetings("goroutine")
	printGreetings("main function")
}

func printGreetings(source string) {
	for i := 0; i < 9; i++ {
		fmt.Println("Hello Gohper!", i, source)
	}
	time.Sleep(time.Millisecond * 1)
}
