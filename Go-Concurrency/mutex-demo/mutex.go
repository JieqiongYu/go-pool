package main

import (
	"fmt"
	"sync"
)

var greetings string
var howdyDone chan bool
var mutex = &sync.Mutex{}

func main() {
	howdyDone = make(chan bool)
	go howdyGreetings()
	// Identify the place where the race might happen and use mutex.Lock() and mutex.Unlock()
	mutex.Lock()
	greetings = "Hello Gopher!"
	fmt.Println(greetings)
	mutex.Unlock()
	<-howdyDone
}

func howdyGreetings() {
	// Identify the place where the race might happen and use mutex.Lock() and mutex.Unlock()
	mutex.Lock()
	greetings = "Howdy Gopher!"
	mutex.Unlock()
	howdyDone <- true
}
