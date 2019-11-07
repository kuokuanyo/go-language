package main

import (
	"fmt"
	"net/http"
)

type hotdog int

/*
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
*/
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	var d hotdog
	//func ListenAndServe(addr string, handler Handler) error
	http.ListenAndServe(":8080", d)
}
