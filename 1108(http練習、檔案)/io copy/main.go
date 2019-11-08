package main

import (
	"io"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", dog)
	http.HandleFunc("/dog.jpg", dogpic)

	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-TYPE", "text/html; charset=utf-8")

	io.WriteString(w, `<img src="dog.jpg">`)
}

func dogpic(w http.ResponseWriter, req *http.Request) {
	//open file
	f, err := os.Open("dog.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}
