package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB
var err error

type Customers struct {
	CustomerNumber         string `db:"customerNumber"`
	CustomerName           string `json:"customerName"`
	ContactLastName        string `json:"contactLastName"`
	ContactFirstName       string `json:"contactFirstName"`
	Phone                  string `json:"phone"`
	AddressLine1           string `json:"addressLine1"`
	AddressLine2           sql.NullString
	City                   string `json:"city"`
	State                  sql.NullString
	PostalCode             string `json:"postalCode"`
	Country                string `json:"country"`
	SalesRepEmployeeNumber sql.NullString
	CreditLimit            string `json:"creditLimit"`
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
		fmt.Printf("%#v\n", post)
	}
}
