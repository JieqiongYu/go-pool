package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("GOOGLE_CLOUD_PROJECT", "kouzoh-p-jieqiong-yu")
	fmt.Println("GOOGLE_CLOUD_PROJECT:", os.Getenv("GOOGLE_CLOUD_PROJECT"))
}
