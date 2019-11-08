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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="dog.jpg">`)
}

func dogpic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("dog.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	//Stat返回FileInfo結構
	//func (f *File) Stat() (FileInfo, error)
	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	//func ServeContent(w ResponseWriter, req *Request, name string, modtime time.Time, content io.ReadSeeker)
	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}
