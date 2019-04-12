# Cloud Native Go

In this demo, it demonstrates things as the following:

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