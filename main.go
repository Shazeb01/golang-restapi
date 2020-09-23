package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB
var err error

type Customers struct {
	CustomerNumber         string         `json:"customerNumber" db:"customerNumber"`
	CustomerName           string         `json:"customerName" db:"customerName"`
	ContactLastName        string         `json:"contactLastName" db:"contactLastName"`
	ContactFirstName       string         `json:"contactFirstName" db:"contactFirstName"`
	Phone                  string         `json:"phone" db:"phone"`
	AddressLine1           string         `json:"addressLine1" db:"addressLine1"`
	AddressLine2           sql.NullString `db:"addressLine2"`
	City                   string         `json:"city" db:"city"`
	State                  sql.NullString `db:"state"`
	PostalCode             sql.NullString `json:"postalCode" db:"postalCode"`
	Country                string         `json:"country" db:"country"`
	SalesRepEmployeeNumber sql.NullString `db:"salesRepEmployeeNumber"`
	CreditLimit            string         `json:"creditLimit" db:"creditLimit"`
}

func main() {
	db, err = sqlx.Open("mysql", "root:@tcp(127.0.0.1:3306)/classicmodels")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/data", getPosts).Methods("GET")

	http.ListenAndServe(":4444", router)
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	post := Customers{}
	rows, err := db.Queryx("SELECT * FROM customers")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.StructScan(&post)
		if err != nil {
			log.Fatalln(err)
		}
		json.NewEncoder(w).Encode(post)
	}
}
