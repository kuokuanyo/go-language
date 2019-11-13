package main

import (
	"html/template"
	"net/http"
)

//模板
var tpl *template.Template

//初始化
func init() {
	//Must初始化模板
	//ParseGlob對路徑解析並創建模板
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	//func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	http.HandleFunc("/", index)
	//func Handle(pattern string, handler Handler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	//connect
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	//執行模板
	tpl.ExecuteTemplate(w, "index.html", nil)
}
