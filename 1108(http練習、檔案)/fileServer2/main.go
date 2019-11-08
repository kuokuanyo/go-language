package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", dog)
	//FileServer返回根目錄的文件系統(.為目前位置)
	//func FileServer(root FileSystem) Handler

	// func StripPrefix(prefix string, h Handler) Handler
	//StripPrefix將第二個參數的路徑以第一個參數位址查詢
	//使用第一個參數請求URL
	http.Handle("/resourses/", http.StripPrefix("/resourses", http.FileServer(http.Dir("./asset"))))

	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//使使用StripPrefix後的路徑
	io.WriteString(w, `<img src="/resourses/dog.jpg">`)
}
