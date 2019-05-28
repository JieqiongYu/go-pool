package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

const (
	connHost       = "localhost"
	connPort       = "8080"
	driverName     = "mysql"
	dataSourceName = "root:password@/mydb"
)

var db *sql.DB
var connectionError error

func init() {
	db, connectionError = sql.Open(driverName, dataSourceName)
	if connectionError != nil {
		log.Fatal("error connecting to database :: ", connectionError)
	}
}

type Employee struct {
	Id   int    `json:"uid"`
	Name string `json:"name"`
}

func updateRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	vals := r.URL.Query()
	name, ok := vals["name"]
	if ok {
		log.Print("going to update record in database for id :: ", id)
		stmt, err := db.Prepare("UPDATE employee SET name=? where uid=?")
		if err != nil {
			log.Print("error occurred while preparing query :: ", err)
			return
		}
		result, err := stmt.Exec(name[0], id)
		if err != nil {
			log.Print("error occurred while executing query :: ", err)
			return
		}
		rowsAffected, err := result.RowsAffected()
		fmt.Fprintf(w, "Number of rows updated in database are :: %d", rowsAffected)
	} else {
		fmt.Fprintf(w, "Error occurred while updating record in database for id :: %s", id)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/employee/update/{id}", updateRecord).Methods("PUT")
	defer db.Close()
	err := http.ListenAndServe(connHost+":"+connPort, router)
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}
}
