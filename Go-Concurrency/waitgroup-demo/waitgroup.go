package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var urlList = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.youtube.com/",
	}

	// Loop through the list of URLs
	for _, url := range urlList {
		// Increment the WaitGroup counter
		wg.Add(1)
		// Call the fetchURL function as a goroutine to fetch the URL
		go fetchURL(url, &wg)
	}

	// Wait for the goroutines that are part of the WaitGroup to finish
	wg.Wait()
}

func fetchURL(url string, wg *sync.WaitGroup) {

	// Decrement the WaitGroup counter once we've fetched the URL
	// Defer is used to ensure that a function call is performed later in a program's execution,
	// usually for purpose of cleanup. defer is often used where e.g. ensure and finally would be used
	// in other languages.
	defer wg.Done()

	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Failed to fetch the URL,", url, "and encountered this error:", err)
	} else {
		fmt.Println("Contents of url,", url, "is:")
		contents, err := ioutil.ReadAll(response.Body)

		response.Body.Close()
		if err != nil {
			log.Fatal("Failed to read the response body and encountered this error:", err)
		}
		fmt.Println(string(contents))
	}
}
