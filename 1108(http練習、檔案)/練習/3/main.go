package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func main() {

	/*三種很相似(陷阱)
	type HandlerFunc func(ResponseWriter, *Request)
	func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	*/

	//func Handle(pattern string, handler Handler)
	//第一種應為
	//轉換函式類型(使用上列的第一種)
	http.Handle("/", http.HandlerFunc(foo))
	http.Handle("/dog/", http.HandlerFunc(bar))
	http.Handle("/me/", http.HandlerFunc(myname))

	/*
	   如果要用上面的第二種
	   func HandleFunc("/", (foo)
	   func HandleFunc("/dog/", bar)
	   func HandleFunc("/me/", myname)
	*/

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
