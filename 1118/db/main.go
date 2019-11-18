package main

import (
	"conn"
)

//設定資料庫資訊
var user = conn.MySqlUser{
	Host:     "127.0.0.1", //主機
	MaxIdle:  10,          //閒置的連接數
	MaxOpen:  10,          //最大連接數
	User:     "root",      //用戶名
	Password: "asdf4440",  //密碼
	Database: "test",      //資料庫名稱
	Port:     3306,        //端口
}

/*
//建立查詢欄位
var (
	名稱	int
	名稱	string
	名稱	bool
)
*/

//上面查詢欄位名稱等於此[]string{}的變數名稱
//須為字串
//var s = []string{上列設定名稱}

func main() {

	//建立初始化連線
	connect_db := user.Init()

	//回傳指標
	db := conn.NewDB(connect_db)

	//最後必須關閉
	defer db.Close()

	//建立資料庫
	db.CreateDb("book")

	//使用資料庫
	db.Use_Db("book")

	//建立資料表
	db.CreateTable("books", "id", "int", "title", "varchar(50)", "author", "varchar(50)", "year", "varchar(50)")

	//插入數值
	db.Insert("books", "id", "1", "title", "Golang is great", "author", "Mr.Great", "year", "2012")
	db.Insert("books", "id", "2", "title", "c++ is greatest", "author", "Mr.C+", "year", "2015")

	/*
		//更改數值
		db.Update_db(資料庫名稱, 設定欄位名稱, 設定新數值, 更改的欄位, 更改欄位的數值)

		//刪除資料庫
		db.Delete_Db(資料庫名稱)

		//刪除資料表
		db.Delete_Tb(資料庫名稱)

		//讀取資料
		//第三個與後面參數長度必須相同
		db.Read(資料庫名稱, s, 設定的變數(var))
	*/

}
