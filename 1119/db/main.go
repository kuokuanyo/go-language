package main

import (
	"conn"
	"encoding/json"
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

var db *conn.DB

//資料庫連線
func init() {
	//建立初始化連線
	connect_db := user.Init()

	//回傳指標
	db = conn.NewDB(connect_db)
}

func main() {
	//最後必須關閉
	defer db.Close()

	//router
	//func NewRouter() *Router
	r := mux.NewRouter()

	//func (r *Route) HandlerFunc(f func(http.ResponseWriter, *http.Request)) *Route
	//func (r *Route) Methods(methods ...string) *Route
	r.HandleFunc("/books", getbooks).Methods("GET")
	r.HandleFunc("/books/{id}", getbook).Methods("GET")
	r.HandleFunc("/books", addbook).Methods("POST")
	r.HandleFunc("/books", updatebook).Methods("PUT")
	r.HandleFunc("/books/{year}", removebook).Methods("DELETE")

	//connect
	http.ListenAndServe(":8080", r)
}

//get all datas
func getbooks(w http.ResponseWriter, req *http.Request) {
	var book Book
	var books []Book

	//尋找data
	rows, err := db.Query("select * from books")
	if err != nil {
		log.Println(err)
	}
	//最後要關閉
	defer rows.Close()
	//處理每一行
	//Next method 迭代查詢資料，回傳bool
	//func (rs *Rows) Next() bool
	for rows.Next() {
		//Scan method方法用來讀取每一列的值
		//func (rs *Rows) Scan(dest ...interface{}) error
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if err != nil {
			log.Panicln(err)
		}
		//add slice
		books = append(books, book)
	}
	//encode
	json.NewEncoder(w).Encode(books)
}

//get one data
func getbook(w http.ResponseWriter, req *http.Request) {
	var book Book
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
	var book Book
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
	var book Book
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
