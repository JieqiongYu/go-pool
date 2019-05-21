# Chaining with Alice

## Installation 

```bash
go get github.com/justinas/alice
```

## Test

```bash
# Correct JSON
curl -i -H "Content-Type: application/json" -X POST http://localhost:8000/city -d '{"name":"Boston", "area":89}'
# Not JSON media type
curl -i -X POST http://localhost:8000/city -d '{"name":"New York", "area":304}' 
```