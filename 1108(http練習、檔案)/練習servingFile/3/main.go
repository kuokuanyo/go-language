package main

import (
	"html/template"
	"log"
	"net/http"
)

//模板
var tpl *template.Template

//初始化
func init() {
	//Must初始化
	//ParseFiles創建模板
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	//FileServer返回該路徑的目錄及文件內容
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/pics/", fs)
	http.HandleFunc("/", dog)

	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	//執行模板
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
