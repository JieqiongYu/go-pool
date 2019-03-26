# Arrays

Code snippet to demonstrate the usage of array in Golang.

It shows these characteristics of arrays in Golang:

* If we do not inialize a value for the array, they will default to the zero value of the type. In the case of integers, we epxect the zero value to be 0.
* Arrays are poassed as value to functions, meaning a copy of the array is passed and not a reference (pointer) to the array.
* Since arrays are values, equality comparisons of two arrays are done value for value. Even though 2 arrays have different memory addresses, that part is not considered as a criteria for the comparison.