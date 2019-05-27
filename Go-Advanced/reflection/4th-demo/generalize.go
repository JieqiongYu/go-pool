package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordID      int
	customerID int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

// Since the `createQuery` funciton should work with any struct, it takes a `interface{}` as argument.
func createQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		// Get the name of the struct for its `reflect.Type` using the `Name()` method.
		// `order`, `employee`
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return
	}
	fmt.Println("unsupported type")
}

func main() {
	o := order{
		ordID:      456,
		customerID: 56,
	}
	createQuery(o)

	e := employee{
		name:    "MerJQ",
		id:      565,
		address: "Roppongi, Tokyo",
		salary:  10000,
		country: "Japan",
	}
	createQuery(e)

	i := 90
	createQuery(i)
}
