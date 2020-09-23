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

type Employees struct {
	EmployeeNumber string         `json:"employeeNumber" db:"employeeNumber"`
	LastName       string         `json:"lastName" db:"lastName"`
	FirstName      string         `json:"firstName" db:"firstName"`
	Extension      string         `json:"extension" db:"extension"`
	Email          string         `json:"email" db:"email"`
	OfficeCode     string         `json:"officeCode" db:"officeCode"`
	ReportsTo      sql.NullString `json:"reportsTo" db:"reportsTo"`
	JobTitle       string         `json:"jobTitle" db:"jobTitle"`
}

func main() {
	db, err = sqlx.Open("mysql", "root:#@tcp(127.0.0.1:3306)/classicmodels")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/employees", getEmployees).Methods("GET")

	http.ListenAndServe(":4444", router)
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
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

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	employee := Employees{}
	rows, err := db.Queryx("SELECT * FROM employees")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.StructScan(&employee)
		if err != nil {
			log.Fatalln(err)
		}
		json.NewEncoder(w).Encode(employee)

	}
}
