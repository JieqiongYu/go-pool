# Cloud Native Go

In this demo, it demonstrates things as the following

[Cloud Native Go](#cloud-native-go) 
- [Simple Go HTTP Server Implementation](#simple-go-http-server-implementation)
- [Docker Image for Go Microservice](#docker-image-for-go-microservice)
  - [1. Build `Dockerfile`](#1-build-dockerfile)
  - [2. Push customized image to Docker Hub](#2-push-customized-image-to-docker-hub)

## Simple Go HTTP Server Implementation

* Simple HTTP server implementation in go
  * Using the Go `net/http` package

    ```go
    import "net/http"
    ```

  * Implementing and start a simple HTTP server

    ```go
    func main() {
        http.HandleFunc("/", index)
        http.HandleFunc("/api/echo", echo)
        http.ListenAndServe(port(), nil)
    }
    ```

  * Defining simple handler functions

    ```go
    func index(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "Hello Cloud Native Go.")
    }
    ```

* JSON Marshalling/Unmarshalling of Go Structs
  * Using the `Go JSON` package

    ```go
    import "encoding/json"
    ```

  * JSON marshalling/unmarshalling of Go structs

    ```go
    /*TOJSON is the marshalling method*/
    func (b Book) TOJSON() []byte {
        TOJSON, err := json.Marshal(b)
        if err != nil {
            panic(err)
        }
        return TOJSON
    }
    ```

    ```go
    /*FromJSON is the unmarshalling method*/
    func FromJSON(data []byte) Book {
        book := Book{}
        err := json.Unmarshal(data, &book)
        if err != nil {
            panic(err)
        }
        return book
    }
    ```

  * Defining Go structs with additional JSON metadata

    ```go
    /*Book struct*/
    type Book struct {
        // define the book
        Title       string `json:"title"`
        Author      string `json:"author"`
        ISBN        string `json:"isbn"`
        Description string `json:"description,omitempty"`
    }
    ```

  * Adding simple REST endpoint with JSON response
* Simple REST API implementation
  * Using the request path and query parameters
  
  ```go
  isbn := r.URL.Path[len("/api/books/"):]
  ```
  
  * Using HTTP status code
  * Using HTTP verbs: GET, PUT, POST, DELETE

## Docker Image for Go Microservice

The base image for this is `golang:1.12.3-alpine3.9`, and the command to get the docker image is:

```bash
docker pull golang:1.12.3-alpine3.9
```

### 1. Build `Dockerfile`

```bash
cd /path/to/Cloud-Native-Go/
docker build -t cloud-native-go:1.0.0 .
```

Remove one image

```bash
docker rmi -f <IMAGE_ID>
```

### 2. Push customized image to Docker Hub

Use `tag` command

```bash
## docker tag [DOCKER_REPOSITORY_NAME:TAG] [DOCKER_ACCOUNT/[DOCKER_REPOSITORY_NAME:TAG]]
docker tag cloud-native-go:1.0.0 jieqiong/cloud-native-go:1.0.0
```

Login docker hub and do the push thing

```bash
docker login
docker push jieqiong/cloud-native-go:1.0.0
```

Run cloud-native-go container

```bash
docker run -it -p 8080:8080 cloud-native-go:1.0.0
```

How to test if it works?

```bash
# in local terminal
curl http://localhost:8080/api/books
```

Try a different port:

```bash
# -e means for `environment`
docker run -it -e "PORT=9090" -p 9090:9090 cloud-native-go:1.0.0
```

How to test if it works?

```bash
# in local terminal
curl http://localhost:9090/api/books
```