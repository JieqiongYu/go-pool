package main

import (
	"fmt"
	"strconv"
)

func main() {
	userID := "1234563789"

	result, err := strconv.ParseInt(userID,10, 64)
	if err != nil || result > 999999999 || result < 100000000 {
		fmt.Println("Not valid when convert or for length")
	}
}
