package main

import (
	"io"
	"net/http"
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
	io.WriteString(w, "hello kuo")
}
