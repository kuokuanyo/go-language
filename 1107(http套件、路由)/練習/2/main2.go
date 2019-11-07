package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func main() {

	//func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", bar)
	http.HandleFunc("/me/", myname)

	//func ListenAndServe(addr string, handler Handler) error
	http.ListenAndServe(":8080", nil)
}

//function
func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo ran")
}

//function
func bar(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "bar ran")
}

//function
func myname(w http.ResponseWriter, r *http.Request) {

	//解析檔案
	tpl, err := template.ParseFiles("index.gohtml")
	//check
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	//執行模板
	err = tpl.ExecuteTemplate(w, "index.gohtml", "kuo")
	//check
	if err != nil {
		log.Fatalln("error parsing template", err)
	}
}
