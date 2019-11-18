package main

import (
	"conn"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Year   string
}

//設定資料庫資訊
var user = conn.MySqlUser{
	Host:     "127.0.0.1", //主機
	MaxIdle:  10,          //閒置的連接數
	MaxOpen:  10,          //最大連接數
	User:     "root",      //用戶名
	Password: "asdf4440",  //密碼
	Database: "book",      //資料庫名稱
	Port:     3306,        //端口
}

var books []Book
var db *conn.DB

func main() {
	//建立初始化連線
	connect_db := user.Init()

	//回傳指標
	db := conn.NewDB(connect_db)

	//最後必須關閉
	defer db.Close()

	//router
	//func NewRouter() *Router
	r := mux.NewRouter()

	r.HandleFunc("/books", getbooks).Methods("GET")
	r.HandleFunc("/books/{id}", getbook).Methods("GET")
	r.HandleFunc("/books", addbook).Methods("POST")
	r.HandleFunc("/books", updatebook).Methods("PUT")
	r.HandleFunc("/books/{id}", removebook).Methods("DELETE")

	//connect
	http.ListenAndServe(":8080", r)
}

func getbooks(w http.ResponseWriter, req *http.Request) {
	//var book Book
	//books := []Book{}

	rows, err := db.Query("select * from books")
	if err != nil {
		log.Panicln(err)
	}

	defer rows.Close()
	/*
		for rows.Next() {
			err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
			if err != nil {
				log.Panicln(err)
			}
			books = append(books, book)
		}

		//encode
		json.NewEncoder(w).Encode(books)
	*/
}

func getbook(w http.ResponseWriter, req *http.Request) {

}

func addbook(w http.ResponseWriter, req *http.Request) {

}

func updatebook(w http.ResponseWriter, req *http.Request) {

}

func removebook(w http.ResponseWriter, req *http.Request) {

}
