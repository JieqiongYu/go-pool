package main

import (
	"fmt"
	"reflect"
)

type order struct {
	orderID    int
	customerID int
}

func createQuery(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	fmt.Println("Type", t)
	fmt.Println("Value", v)
}

func main() {
	o := order{
		orderID:    456,
		customerID: 56,
	}
	createQuery(o)
}
