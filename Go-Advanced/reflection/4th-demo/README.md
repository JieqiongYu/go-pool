# Generalize Query Creator 

Generalize Query Creator and make it work on any struct. 

For example, if we pass the struct below, 

```go
o := order {
    orderID: 1233,
    customerID: 567
}
```

our `createQuery` function should return 

```sql
insert into order values(1234, 567)
```

Similarly if we pass 

```go
e := employee {
    name: "MerJQ",
    id: 565,
    address: "Roppongi, Tokyo",
    salary: 10000,
    country: "Japan",
}
```

It should return, 

```sql
insert into employee values ("MerJQ", 565, "Roppongi, Tokyo", 10000, "Singapore")
```

Since the `createQuery` funciton should work with any struct, it takes a `interface{}` as argument.

The `createQuery` function should work on any struct. The only way to write this function is to examine the type of the struct argument passed to it at run time, find its fields and then create the query. 

This is where reflection is useful. 

##   `reflect` package

The `reflect` package helps to identify the underlying concrete type and the value of a `interface{}` variable. 

### `reflect.Type` and `reflect.Value`

The concrete type of `interface{}` is represented by `reflect.Type` and the underlying value is represented by `reflect.Value`. 

There are two functions 

* `reflect.TypeOf()`
    * return `reflect.Type`
* `reflect.ValueOf()`
    * return `reflect.Value`



