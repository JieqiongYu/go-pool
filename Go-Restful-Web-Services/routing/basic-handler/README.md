# Understanding Go's net/http package

It's a small Go program called basicHandler.go that defines the route and a function handler. 

## Find the port usage and kill the <PID>

```bash
lsof -i :8000
kill -9 <PID>
```

## Test

```bash
curl -X GET http://localhost:8000/hello
```