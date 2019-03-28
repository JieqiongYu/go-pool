package main

import "fmt"

var greetings string
var howdyDone chan bool

// Normal command: go run racedemo.go
// Race condition command: go run -race race-demo.go
func main() {
	howdyDone = make(chan bool)
	go howdyGreetings()
	greetings = "Hello Gohper!"
	fmt.Println(greetings)
	<-howdyDone
}

func howdyGreetings() {
	greetings = "Howdy Gopher!"
	howdyDone <- true
}
