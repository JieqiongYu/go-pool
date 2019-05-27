package main

import (
	"fmt"
	"reflect"
)

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}
}

func main() {
	e := employee{
		name:    "MerJQ",
		id:      565,
		address: "Roppongi, Tokyo",
		salary:  10000,
		country: "Japan",
	}
	createQuery(e)
}
