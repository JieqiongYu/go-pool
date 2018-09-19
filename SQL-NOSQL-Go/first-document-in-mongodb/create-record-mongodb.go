package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	mgo "gopkg.in/mgo.v2"
)

const (
	connHost   = "localhost"
	connPort   = "8080"
	mongoDbUrl = "127.0.0.1"
)

var session *mgo.Session
var connectionError error

type Employee struct {
	Id   int    `json:"uid"`
	Name string `json:"name"`
}

func init() {
	session, connectionError = mgo.Dial(mongoDbUrl)
	if connectionError != nil {
		log.Fatal("error connecting to database :: ", connectionError)
	}
	session.SetMode(mgo.Monotonic, true)
}

func createDocument(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	name, nameOk := vals["name"]
	id, idOk := vals["id"]
	if nameOk && idOk {
		employeeId, err := strconv.Atoi(id[0])
		if err != nil {
			log.Print("error converting string id to int :: ", err)
			return
		}
		log.Print("going to insert document in database for name :: ", name[0])
		collection := session.DB("mydb").C("employee")
		err = collection.Insert(&Employee{employeeId, name[0]})
		if err != nil {
			log.Print("error occurred while inserting document in database :: ", err)
			return
		}
		fmt.Fprintf(w, "Last created document id is :: %s", id[0])
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/employee/create", createDocument).Methods("POST")
	defer session.Close()
	err := http.ListenAndServe(connHost+":"+connPort, router)
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}
}
