package main

import "fmt"

type order struct {
	ordID      int
	customerID int
}

func createQuery(o order) string {
	i := fmt.Sprintf("insert into order values(%d, %d)", o.ordID, o.customerID)
	return i
}

func main() {
	o := order{
		ordID:      1234,
		customerID: 567,
	}
	fmt.Println(createQuery(o))
}
