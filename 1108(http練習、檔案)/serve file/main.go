package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", dog)
	http.HandleFunc("/dog.jpg", dogpic)

	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="dog.jpg">`)
}

func dogpic(w http.ResponseWriter, req *http.Request) {
	//ServeFile使用命名的文件或目錄回答請求
	//func ServeFile(w ResponseWriter, r *Request, name string)
	http.ServeFile(w, req, "dog.jpg")
}
