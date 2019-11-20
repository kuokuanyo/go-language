package main

import (
	"book-db/controllers"
	"book-db/driver"
	"database/sql"
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
	r.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	r.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	r.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	r.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")

	//connect
	http.ListenAndServe(":8080", r)
}
