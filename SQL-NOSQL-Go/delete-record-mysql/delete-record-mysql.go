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

func deleteRecord(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	name, ok := vals["name"]
	if ok {
		log.Print("going to delete record in database for name :: ", name[0])
		stmt, err := db.Prepare("DELETE FROM employee where name=?")
		if err != nil {
			log.Print("error occurred while preparing query :: ", err)
			return
		}
		result, err := stmt.Exec(name[0])
		if err != nil {
			log.Print("error occurred while executing query :: ", err)
			return
		}
		rowsAffected, err := result.RowsAffected()
		fmt.Fprintf(w, "Number of rows deleted in databases are :: %d", rowsAffected)
	} else {
		fmt.Fprintf(w, "Error occurred while deleting record in database for name %s", name[0])
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/employee/delete", deleteRecord).Methods("DELETE")
	defer db.Close()
	err := http.ListenAndServe(connHost+":"+connPort, router)
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}
}
