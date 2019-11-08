package main

import (
	"io"
	"net/http"
)

func main() {

	//FileServer返回根目錄的文件系統(.為目前位置)
	//func FileServer(root FileSystem) Handler
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog.jpg", dog)

	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="dog.jpg">`)
}
