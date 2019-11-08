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
	//Must初始化
	//解析檔案並創建模板
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	http.HandleFunc("/", dog)
	//FileServer返回該路徑的目錄及文件內容
	fs := http.FileServer(http.Dir("public"))
	//StripPrefix利用第二個參數路徑為第一個參數檔名
	http.Handle("/public/", http.StripPrefix("/public", fs))
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	//執行模板
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
