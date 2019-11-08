package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/facicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	//FormValue回傳命名組件的第一個值作為查詢的動作
	//func (r *Request) FormValue(key string) string
	v := req.FormValue("q")
	io.WriteString(w, "Do my search: "+v)
}

//http://localhost:8080/?q=查詢字串
