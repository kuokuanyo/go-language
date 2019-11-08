package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL) //印在終端機上
	fmt.Fprintln(w, "go look at your terminal")
}
