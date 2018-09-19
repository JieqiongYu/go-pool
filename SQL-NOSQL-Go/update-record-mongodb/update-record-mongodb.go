package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

func updateDocument(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	vals := r.URL.Query()
	name, ok := vals["name"]
	if ok {
		employeeId, err := strconv.Atoi(id)
		if err != nil {
			log.Print("error converting string id to int :: ", err)
			return
		}
		log.Print("going to update document in database for id :: ", id)
		collection := session.DB("mydb").C("employee")
		var changeInfo *mgo.ChangeInfo
		changeInfo, err = collection.Upsert(bson.M{"id": employeeId}, &Employee{employeeId, name[0]})
		if err != nil {
			log.Print("error occurred while updating record in database :: ", err)
			return
		}
		fmt.Fprintf(w, "Number of documents updated in database are :: %d", changeInfo.Updated)
	} else {
		fmt.Fprintf(w, "Error occurred while updating document in database for id :: %s", id)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/employee/update/{id}", updateDocument).Methods("PUT")
	defer session.Close()
	err := http.ListenAndServe(connHost+":"+connPort, router)
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}
}
