# Gorilla's Handler Middleware for Logging 

The most important middleware for common tasks: 

* `LoggingHandler`: For logging in Apache Common Log Format
* `ComprehenssionHandler`: For zipping the responses
* `RecoveryHandler`: For recovering from unexpected panics

## Installation 

```bash
go get "github.com/gorilla/handlers"
```

## Format of the log

```
IP-Date-Method:Endpoint-ResponseStatus
```

## Test

```bash
# Not use `localhost` this time as the IP part of the log would be empty. 
curl -X GET "http://127.0.0.1:8000"
```