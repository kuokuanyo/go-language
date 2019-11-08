package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpg", kuo)

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo run")
}

func dog(w http.ResponseWriter, req *http.Request) {
	//解析文件並創建模板
	tpl, err := template.ParseFiles("dog.html")
	if err != nil {
		log.Fatalln(err)
	}
	//執行模板
	tpl.ExecuteTemplate(w, "dog.html", nil)
}

func kuo(w http.ResponseWriter, req *http.Request) {
	//ServeFile回覆文件或目錄的請求
	//func ServeFile(w ResponseWriter, r *Request, name string)
	http.ServeFile(w, req, "dog.jpg")
}
