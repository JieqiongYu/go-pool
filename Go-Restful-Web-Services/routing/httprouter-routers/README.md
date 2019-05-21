# httprouter 

The program is trying to implement a REST service using `httprouter`. We are defining two routes here: 
 
* /api/v1/go-version
* /api/v1/show-file/:name

## Installation 

```bash
go get github.com/julienschmidt/httprouter
```

## Test

```bash
curl -X GET http://localhost:8000/api/v1/go-version
curl -X GET http://localhost:8000/api/v1/show-file/text.txt
```

## Note

This example can be extended to any system information.