package main

import (
	"io"
	"net/http"
)

func main() {

	//func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	http.HandleFunc("/", dog)

	//func ListenAndServe(addr string, handler Handler) error
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<!--not serving from our server-->
	<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/golde33443.jpg">
	`)
}
