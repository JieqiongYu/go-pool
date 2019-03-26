package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-pool/Go-Data-Structures/type-and-interface/simpleshape"
)

func main() {
	myRectangle := simpleshape.NewRectangle(4, 5)
	myTriangle := simpleshape.NewTriangle(2, 7)
	shapesSlice := []interface{}{myRectangle, myTriangle}

	for index, shape := range shapesSlice {
		fmt.Printf("Shape in index, %d, of the shapesSlice is a %T\n", index, shape)
	}

	aRandomShape := giveMeARandomShape()
	fmt.Printf("The type of aRandomShape is %T\n", aRandomShape)

	switch t := aRandomShape.(type) {
	case *simpleshape.Rectangle:
		fmt.Println("I got back a rectangle with an area equal to", t.Area())
	case *simpleshape.Triangle:
		fmt.Println("I got back a triangle with an area equal to", t.Area())
	default:
		fmt.Println("I don't recognize what I got back!")
	}
}

func giveMeARandomShape() interface{} {

	var shape interface{}
	var randomShapesSlice = make([]interface{}, 2)

	// Seed the random generator
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	// Pick a random number, either 1 or 0
	randomNumber := r.Intn(2)
	fmt.Println("Random Number: ", randomNumber)

	// Let's make a new rectangle
	rectangle := simpleshape.NewRectangle(3, 6)

	// Let's make a new triangle
	triangle := simpleshape.NewTriangle(9, 18)

	// Let's store the shapes into a slice data structure
	randomShapesSlice[0] = rectangle
	randomShapesSlice[1] = triangle

	shape = randomShapesSlice[randomNumber]

	return shape
}
