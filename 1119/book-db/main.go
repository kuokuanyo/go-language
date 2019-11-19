package main

import (
	"book-db/controllers"
	"book-db/driver"
	"book-db/model"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var db *sql.DB

//資料庫連線
func init() {
	db = driver.ConnectDB()
}

func main() {
	//最後必須關閉
	defer db.Close()

	controller := controllers.Controller{}
	//router
	//func NewRouter() *Router
	r := mux.NewRouter()

	//func (r *Router) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route
	//func (r *Route) Methods(methods ...string) *Route
	r.HandleFunc("/books", controller.Getbooks(db)).Methods("GET")
	r.HandleFunc("/books/{id}", getbook).Methods("GET")
	r.HandleFunc("/books", addbook).Methods("POST")
	r.HandleFunc("/books", updatebook).Methods("PUT")
	r.HandleFunc("/books/{year}", removebook).Methods("DELETE")

	//connect
	http.ListenAndServe(":8080", r)
}

//get one data
func getbook(w http.ResponseWriter, req *http.Request) {
	var book model.Book
	//return map
	//func Vars(r *http.Request) map[string]string
	params := mux.Vars(req)

	rows := db.QueryRow("select * from books where id=?", params["id"])
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	if err != nil {
		log.Panicln(err)
	}
	//encode
	json.NewEncoder(w).Encode(book)

}

func addbook(w http.ResponseWriter, req *http.Request) {
	var book model.Book
	//decode(r.Body)
	//必須使用pointer
	json.NewDecoder(req.Body).Decode(&book)

	_, err := db.Exec("insert into books (id, title, author, year) values(?, ?, ?, ?);",
		book.ID, book.Title, book.Author, book.Year)
	if err != nil {
		log.Println(err)
	}
	//encode
	json.NewEncoder(w).Encode(book.ID)
}

func updatebook(w http.ResponseWriter, req *http.Request) {
	var book model.Book
	//decode(req.Body)
	//pointer
	json.NewDecoder(req.Body).Decode(&book)

	_, err := db.Exec("update books set title=?, author=?, year=? where id=?",
		book.Title, book.Author, book.Year, book.ID)
	if err != nil {
		log.Println(err)
	}
	//encode
	json.NewEncoder(w).Encode(book)
}

func removebook(w http.ResponseWriter, req *http.Request) {
	//return map
	//func Vars(r *http.Request) map[string]string
	params := mux.Vars(req)

	_, err := db.Exec("delete from books where year=?", params["year"])
	if err != nil {
		log.Println(err)
	}

	//encode
	json.NewEncoder(w).Encode("successful")
}
