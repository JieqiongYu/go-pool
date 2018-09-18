package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

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
		log.Fatal("error connecting to database :: ", connectionError)
	}
}

type Employee struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func readRecords(w http.ResponseWriter, r *http.Request) {
	log.Print("reading records from database")
	rows, err := db.Query("SELECT * FROM employee")
	if err != nil {
		log.Print("error occurred while executing select query :: ", err)
		return
	}
	employees := []Employee{}
	for rows.Next() {
		var uid int
		var name string
		err = rows.Scan(&uid, &name)
		employee := Employee{Id: uid, Name: name}
		employees = append(employees, employee)
	}
	json.NewEncoder(w).Encode(employees)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/employees", readRecords).Methods("GET")
	defer db.Close()
	err := http.ListenAndServe(connHost+":"+connPort, router)
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}
}
