package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

type Post struct {
	customerNumber         string `json:"customerNumber"`
	customerName           string `json:"customerName"`
	contactLastName        string `json:"contactLastName"`
	contactFirstName       string `json:"contactFirstName"`
	phone                  string `json:"phone"`
	addressLine1           string `json:"addressLine1"`
	addressLine2           string `json:"addressLine2"`
	city                   string `json:"city"`
	state                  string `json:"state"`
	postalCode             string `json:"postalCode"`
	country                string `json:"country"`
	salesRepEmployeeNumber string `json:"salesRepEmployeeNumber"`
	creditLimit            string `json:"creditLimit"`
}

func main() {
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/classicmodels")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/posts", getPosts).Methods("GET")

	http.ListenAndServe(":5555", router)
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var posts []Post

	result, err := db.Query("SELECT customerNumber,customerName, contactFirstName from customers")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var post Post
		err := result.Scan(&post.customerNumber, &post.customerName, &post.contactFirstName, &post.contactLastName, &post.phone, &post.addressLine1, &post.addressLine2, &post.city, &post.state, &post.postalCode, &post.country, &post.salesRepEmployeeNumber, &post.creditLimit)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}

	json.NewEncoder(w).Encode(posts)
}
