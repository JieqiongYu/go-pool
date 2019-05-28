package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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
		log.Fatal("error connecting to database: ", connectionError)
	}
}

func createRecord(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	name, ok := vals["name"]
	if ok {
		log.Print("going ot insert record in database for name: ", name[0])
		stmt, err := db.Prepare("INSERT employee SET name=?")
		if err != nil {
			log.Print("error preparing query :: ", err)
			return
		}
		result, err := stmt.Exec(name[0])
		if err != nil {
			log.Print("error executing query :: ", err)
			return
		}
		id, err := result.LastInsertId()
		fmt.Fprintf(w, "Last Inserted Record Id is :: %s", strconv.FormatInt(id, 10))
	} else {
		fmt.Fprintf(w, "Error occurred while creating record in database for name :: %s", name[0])
	}

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/employee/create", createRecord).Methods("POST")
	defer db.Close()
	err := http.ListenAndServe(connHost+":"+connPort, router)
	if err != nil {
		log.Fatal("error starting http server :", err)
		return
	}
}
