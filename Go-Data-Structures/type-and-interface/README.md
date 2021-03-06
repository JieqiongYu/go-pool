# struct

Code snippet to demonstrate how struct works in Golang Language.

## Types Declaration

* Type Method Declaration Using A Point Receiver
We Declare Methods Using A Pointer Receiver

```go
type typeName struct {
    field1 typeOfField1
    field2 typeOfField2
}

func (r *typeName) methodName() returnType {
}
```

## Interface Declaration

Declaring An Interface

```go
type InterfaceName interface {
    methodName1() returnType1
    methodName2() returnType2
}
```

**Note:** A type implements an interface when it has defined all the methods in the method list of an interface.

## Embedding Interfaces

We can use pieces of an implementation by embedding types, either within a struct or an interface. 

```go
type Reader interface {
    Read(p []byte)(n int, err error)
}

type Writer interface {
    Write(p []byte)(n int, err error)
}

type ReadWriter interface {
    Reader
    Writer
}
```
