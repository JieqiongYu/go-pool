package main

import "fmt"

func main() {

	// We use the make() built-in function to create a new slice of length 5.
	// Here we make a slice of integers of length 5.
	var mySlice = make([]int, 5)
	fmt.Printf("Contents of mySlice: %v\n\n", mySlice)

	populateIntegerSlice(mySlice)
	fmt.Printf("Contents of mySlice: %v\n", mySlice)
	// We can use the len() built-in function to return the length of the slice
	fmt.Println("The length of mySlice is: ", len(mySlice))
	// We can use the cap() built-in funciton to return the capacity of the slice
	fmt.Println("The capacity of mySlice is: ", cap(mySlice))

	// Add a new element to mySlice, and notice the changes to the length and capacity of the slice
	fmt.Println("After adding a new element to mySlice...")
	mySlice = append(mySlice, 18)
	fmt.Printf("Contents of mySlice: %v\n", mySlice)
	fmt.Println("The length of mySlice is: ", len(mySlice))
	fmt.Println("The capacity of mySlice is: ", cap(mySlice))

	// This short hand notation allows us to get the element from index 1 up to
	// (but not including) index 4.
	s := mySlice[1:4]
	fmt.Println("mySlice[1:4]", s)

	// When you slice a slice, you get a reference back to that slice.
	// Any changes you make to the subslice will be reflected in the original slice.
	s[0] = 7 // this will cause mySlice[1] to be equal to 7
	fmt.Println("mySlice:", mySlice)

	// All elements in mySlice up to (not including) the element at index 4
	t := mySlice[:4]
	fmt.Println("mySlice[:4]", t)

	// The elements from (and including) the elements at index 1
	u := mySlice[1:]
	fmt.Println("mySlice[1:]", u)

	changingLineupExample()
}

// Slices get passed by reference into functions, meaning, if make changes to
// a slice within a function, our change will be reflected to the slice that
// was passed into the function.
func populateIntegerSlice(input []int) {

	// We set values in a slice, just like we do with arrays using the square brackets
	// notation [] with the element index enclosed within the square brackets.
	input[0] = 3
	input[1] = 6
	input[2] = 9
	input[3] = 12
	input[4] = 15
}

func changingLineupExample() {

	// Declaring and initializing a slice representing the original band members of
	// rock band, INXS using a slice literal
	// (notice it looks just like an array literal except without the element count)
	fmt.Println("The original INXS lineup:")
	inxs := []string{"Michael", "Andrew", "John", "Tim", "Kirk", "Garry"}
	fmt.Println(inxs)

	// Michael left the band in 1997
	fmt.Println("The lineup without Michael:")
	// inxs[:0]: an empty slice; inxs[1:]... means from the 2nd element to the end of the slice
	inxs = append(inxs[:0], inxs[1:]...)
	fmt.Println(inxs)

	// JD joins the band in 2005
	fmt.Println("The lineup with a new member, JD:")
	inxs = append(inxs, "JD")
	fmt.Println(inxs)
}
