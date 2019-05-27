package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 56
	x := reflect.ValueOf(a).Int()
	fmt.Printf("type: %T value: %v\n", x, x)

	b := "MerJQ"
	y := reflect.ValueOf(b).String()
	fmt.Printf("type: %T value: %v\n", y, y)
}
