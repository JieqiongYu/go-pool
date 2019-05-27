# Reflection

## Knowledge about Reflection

### What is reflection?

Reflection is the ability of a program to inspect its variables and values at run time and find their type. 

###  What is the need to inspect a variable and find its type? 

To answer this question, we'll look at the following demos: 

*  First, `simple.go`

And then, from here, we are going to introduce a problem. 

* Second, `struct-demo.go`

We will create a SQL insert Query using the struct in the prvious demo `struct-demo.go`

* Third, `createQuery.go`

If we want the same `createQuery` function to help create different structs, it's where we need to examine the type of the struct argument passed to it at run time, find its fields and then create the query. 

* Fourth, `generalize.go`

### reflect package

The `reflect` package implements run-time reflection in Go.

* `reflect.Type`
* `reflect.Value`

Demo: `reflectTypeValue.go`

* `reflect.Kind`

Demo: `reflectKind.go`

* `NumField()` and `Field()` methods

Demo: `numberField-field.go`

* `Int()` and `String()`

Demo: `intAndString.go`

### Should reflection be used?

`Bob Pike`'s proverb: 

>Clear is better than clever. Reflection is never clear.

Reflection is a very powerful and advanced concept in Go and it should be used with care. It is very difficult to write clear and maintainable code using reflection. It should be avoided wherever possible and should be used only when absolutely necessary. 

## Reference
* https://golangbot.com/reflection/