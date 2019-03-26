# Empty Interface Demo

This is a demo to demonstrate the in which situations we might find an empty interface `interface{}` be useful. 

Every type in Go implements the empty interface.

`interface{}`

Some Use Cases

1. A function that returns a value of `interface{}` can return any type.
2. We can store heterogeneous values in an array, slice, or map using the `interface{}` type.