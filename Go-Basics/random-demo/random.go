package main

import (
	"fmt"
	"math/rand"
	"time"
)

var tasks = []string{"Learning Japanese", "Coding", "Watch TV", "Work Out"}

func main() {
	generateRandomTasks()
}

func generateRandomTasks() {

	rand.Seed(time.Now().UnixNano())

	randomTask := tasks[rand.Intn(len(tasks))]
	fmt.Println("Random task is", randomTask)
}
