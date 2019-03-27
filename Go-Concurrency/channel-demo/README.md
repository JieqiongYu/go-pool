# Channel Demo

Code snippet to deomnstarte how `channel` is used in golang language.

* Channels are conduits (pipes) that you can use to pass values of a particular type from one goroutine to another
* Channels allow goroutines to share memory by communicating
* We use the channel operator, `<-`, to send and receive values [Note: The data flows in the direction of the arrow]
* We can create a channel using the built-in make function:

```go
ch := make(chan type)
```