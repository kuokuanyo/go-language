package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog doggy")
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	var d hotdog
	var c hotcat

	//分配並回傳新的ServeMux
	//func NewServeMux() *ServeMux
	mux := http.NewServeMux()

	//func (mux *ServeMux) Handle(pattern string, handler Handler)
	//第一個參數為位址
	//第二個參數為該網址的處理方式
	mux.Handle("/dog/", d) //dog之後還可以加入字
	mux.Handle("/cat", c)  //cat結尾

	http.ListenAndServe(":8080", mux)
}
