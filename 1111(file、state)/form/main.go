package main

import (
	"log"
	"net/http"
	"text/template"
)

//模板
var tpl *template.Template

//初始化
func init() {
	//Must初始化模板
	//ParseFiles解析檔案並創建模板
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

//new type
type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	//ListenAndServe連接
	http.ListenAndServe(":8080", nil)
}

//function
func foo(w http.ResponseWriter, req *http.Request) {
	//FormValue返回查詢組件的第一個值
	f := req.FormValue("first")
	l := req.FormValue("last")
	s := req.FormValue("subscribe") == "on"

	//執行模板
	err := tpl.ExecuteTemplate(w, "index.gohtml", person{f, l, s})
	if err != nil {
		//使用特定的代碼回覆請求
		//func Error(w ResponseWriter, error string, code int)
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
