package main

import "fmt"

type order struct {
	ordID      int
	customerID int
}

func main() {
	o := order{
		ordID:      1234,
		customerID: 567,
	}
	fmt.Println(o)
}
