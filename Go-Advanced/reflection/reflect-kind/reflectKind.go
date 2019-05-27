package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordID      int
	customerID int
}

func createQuery(q interface{}) {
	t := reflect.TypeOf(q)
	k := t.Kind()
	fmt.Println("Type", t)
	fmt.Println("Kind", k)

}

func main() {
	o := order{
		ordID:      456,
		customerID: 56,
	}
	createQuery(o)
}
