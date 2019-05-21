# Gorilla JSON RPC Server

## Installation 

```bash
go get github.com/gorilla/rpc
```

## Test

```bash
curl -X POST \
   http://localhost:1234/rpc \
   -H 'cache-control: no-cache' \
   -H 'content-type: application/json' \
   -d '{
   "method": "JSONServer.GiveBookDetail",
   "params": [{
   "ID": "1234"
   }],
   "id": "1"
}'
```