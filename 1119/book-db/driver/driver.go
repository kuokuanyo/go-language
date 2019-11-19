package driver

import (
	"database/sql"
	"fmt"
	"log"
)

//用戶資料
type MySqlUser struct {
	Host string //主機
	//最大連接數
	MaxIdle  int
	MaxOpen  int
	User     string //用戶名
	Password string //密碼
	Database string //資料庫名稱
	Port     int    //端口
}

var db *sql.DB

//定義資料庫連線連線
//完整的資料格式: [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
//mehtod
func ConnectDB() *sql.DB {
	//設定資料庫資訊
	user := MySqlUser{
		Host:     "127.0.0.1", //主機
		MaxIdle:  10,          //閒置的連接數
		MaxOpen:  10,          //最大連接數
		User:     "root",      //用戶名
		Password: "asdf4440",  //密碼
		Database: "book",      //資料庫名稱
		Port:     3306,        //端口
	}
	//資料庫連結字串
	//func Sprintf(format string, a ...interface{}) string
	DataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		user.User,
		user.Password,
		user.Host,
		user.Port,
		user.Database)

	//開啟資料庫連線(sql.Open只是初始化sql.DB物件)
	//func Open(driverName, dataSourceName string) (*DB, error)
	//第一個參數為驅動名稱，第二個參數為資料庫的連結
	db, err := sql.Open("mysql", DataSourceName)
	//檢查錯誤
	if err != nil {
		log.Fatal(err)
	}

	//立即檢查資料庫連線是否可用
	//func (db *DB) Ping() error
	err = db.Ping()
	//檢查錯誤
	if err != nil {
		log.Fatal(err)
	}

	//設定最大連接數
	//SetMaxIdleConns設置閒置的連接數
	db.SetMaxIdleConns(user.MaxIdle)
	//SetMaxOpenConns設置最大打開的連接數，默認值為0代表沒有限制
	db.SetMaxOpenConns(user.MaxOpen)
	//無錯誤
	return db
}
