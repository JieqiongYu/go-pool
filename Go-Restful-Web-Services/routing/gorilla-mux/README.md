# Gorilla Mux 

For Gorilla Mux, it allows the following features:

* Path-based matching
* Query-based matching
* Domain-based matching
* Sub-domain based matching
* Reverse URL generation 

It differs from `httprouter` as it modifes the request object instead of using an additioanl argument to pass the URL parameters. 

## Installation 

```bash
go get -u github.com/gorilla/mux
```

## Test

```bash
curl http://localhost:8000/articles/books/123
```
