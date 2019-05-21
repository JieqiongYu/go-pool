# Query-based matching 

Query parameters are those that get passed along with the URL. This is what we commonly see in a REST GET request. 

```bash
http://localhost:8000/articles/?id=123&category=books
```

## Test

```bash
curl -X GET http://localhost:8000/articles/?id=123&category=books
```