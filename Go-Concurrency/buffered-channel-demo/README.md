# Buffered Channels

Code snippet to demonstrate what `buffered channel` is and how to use it.

* Buffered channels are asynchronous, sending and receiving messages through the channel will not block unless the channel is full 
* We can create a buffered channel using the built-in make function:

```go
ch := make(chan type, capacity)
```

Note: if passing the `capacity` as `1`, it's a regular channel.
