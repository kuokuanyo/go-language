package driver

import (
	"conn"
	"database/sql"
)

var db *sql.DB

//設定資料庫資訊
var user = conn.MySqlUser{
	Host:     "127.0.0.1", //主機
	MaxIdle:  10,          //閒置的連接數
	MaxOpen:  10,          //最大連接數
	User:     "root",      //用戶名
	Password: "asdf4440",  //密碼
	Database: "user",      //資料庫名稱
	Port:     3306,        //端口
}

func Connect() *sql.DB {
	//建立初始化連線
	db = user.Init()
	return db
}
