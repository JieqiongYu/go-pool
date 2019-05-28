package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	connHost   = "localhost"
	connPort   = "8080"
	mongoDbURL = "127.0.0.1"
)

var session *mgo.Session
var connectionError error

func init() {
	session, connectionError = mgo.Dial(mongoDbURL)
	if connectionError != nil {
		log.Fatal("error connecting to database :: ", connectionError)
	}
	session.SetMode(mgo.Monotonic, true)
}

type Employee struct {
	Id   int    `json:"id`
	Name string `json:"name"`
}

func readDocuments(w http.ResponseWriter, r *http.Request) {
	log.Print("reading documents from database")
	var employees []Employee
	collection := session.DB("mydb").C("employee")
	err := collection.Find(bson.M{}).All(&employees)
	if err != nil {
		log.Print("error occurred while reading documents from database :: ", err)
		return
	}
	json.NewEncoder(w).Encode(employees)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/employees", readDocuments).Methods("GET")
	defer session.Close()
	err := http.ListenAndServe(connHost+":"+connPort, router)
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}
}
