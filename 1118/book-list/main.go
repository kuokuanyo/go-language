package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	_ "database/sql"

	_ "github.com/lib/pq"
	_ "github.com/subosito/gotenv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Year   string
}

var books []Book

func main() {
	//router
	//func NewRouter() *Router
	r := mux.NewRouter()

	books = append(books, Book{1, "Golang pointers", "Mr. Golang", "2010"})
	books = append(books, Book{2, "Goroutines", "Mr. Goroutine", "2011"})
	books = append(books, Book{3, "Golang routers", "Mr. Router", "2012"})
	books = append(books, Book{4, "Golang concurrency", "Mr. Currency", "2013"})
	books = append(books, Book{5, "Golang good parts", "Mr. Good", "2014"})

	r.HandleFunc("/books", getbooks).Methods("GET")
	r.HandleFunc("/books/{id}", getbook).Methods("GET")
	r.HandleFunc("/books", addbook).Methods("POST")
	r.HandleFunc("/books", updatebook).Methods("PUT")
	r.HandleFunc("/books/{id}", removebook).Methods("DELETE")

	//connect
	http.ListenAndServe(":8080", r)
}

func getbooks(w http.ResponseWriter, req *http.Request) {
	//encode
	json.NewEncoder(w).Encode(books)
}

func getbook(w http.ResponseWriter, req *http.Request) {
	//return map
	//func Vars(r *http.Request) map[string]string
	params := mux.Vars(req)

	//type convert
	i, _ := strconv.Atoi(params["id"])

	for _, book := range books {
		if book.ID == i {
			//encode
			json.NewEncoder(w).Encode(&book)
		}
	}
}

func addbook(w http.ResponseWriter, req *http.Request) {
	var book Book
	//decode
	_ = json.NewDecoder(req.Body).Decode(&book)

	books = append(books, book)

	//encode
	json.NewEncoder(w).Encode(books)
}

func updatebook(w http.ResponseWriter, req *http.Request) {
	var book Book
	//decode
	_ = json.NewDecoder(req.Body).Decode(&book)

	//update
	for i, item := range books {
		if item.ID == book.ID {
			books[i] = book
		}
	}

	//encode
	json.NewEncoder(w).Encode(books)
}

func removebook(w http.ResponseWriter, req *http.Request) {
	//return map
	//func Vars(r *http.Request) map[string]string
	params := mux.Vars(req)

	//convert
	id, _ := strconv.Atoi(params["id"])

	for i, item := range books {
		if item.ID == id {
			books = append(books[:i], books[i+1:]...)
		}
	}

	//encode
	json.NewEncoder(w).Encode(books)
}
