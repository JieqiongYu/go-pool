# Goroutines

Code snippet to demonstrate how goroutine works in Golang language.

* Lightweight threads that are managed by the Go runtime
* Place the keyword `go` before a function call to execute it as a goroutine

## Executing a Function as a Goroutine

* Executing a function as a goroutine
* `go funcName()`
* Executing an anonymous function as a goroutine

```go
go func()
{
    // function implementation goes here
}()
```