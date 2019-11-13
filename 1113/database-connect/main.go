package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//資料庫
var db *sql.DB
var err error

func main() {
	//連接資料庫
	//完整的資料格式: [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	//開啟資料庫連線(sql.Open只是初始化sql.DB物件)
	//func Open(driverName, dataSourceName string) (*DB, error)
	//第一個參數為驅動名稱，第二個參數為資料庫的連結
	db, err = sql.Open("mysql", "root:asdf4440@tcp(127.0.0.1:3306)/test")
	check(err)
	//結束後關閉資料庫
	defer db.Close()

	//立即檢查資料庫連線是否可用
	//func (db *DB) Ping() error
	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	//連接
	http.ListenAndServe(":8080", nil)
}

//function index
func index(w http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(w, "Successfully completed")
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
