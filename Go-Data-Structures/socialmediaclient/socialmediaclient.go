package main

import (
	"fmt"

	"github.com/go-pool/Go-Data-Structures/socialmedia"
)

func main() {
	myPost := socialmedia.NewPost("MerJQ", socialmedia.Moods["thrilled"], "Go is awesome!", "Check out the Go web site!", "https://golang.org", "", "", []string{"go", "golang", "programming language"})
	fmt.Printf("myPost: %+v\n", myPost)
}
