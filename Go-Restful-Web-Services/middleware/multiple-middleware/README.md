# Multiple middleware and chaining.

Create a city API for saving city details. For simplicity's sake, the API will have one POST method, and the body consists of two fields: 
* city name 
* city area

Requirements: 
* Only allow JSON media type
* Send the server time in UTC back to the client for every request. 

## no - middleware

### Test

```bash
curl -H "Content-Type: application/json" -X POST http://localhost:8000/city -d '{"name":"New York", "area":304}'
curl -H "Content-Type: application/json" -X POST http://localhost:8000/city -d '{"name":"Boston", "area":89}'
```

## Multiple Middleware

### Test

```bash
# Correct JSON
curl -i -H "Content-Type: application/json" -X POST http://localhost:8000/city -d '{"name":"Boston", "area":89}'
# Not JSON media type
curl -i -X POST http://localhost:8000/city -d '{"name":"New York", "area":304}' 
```
