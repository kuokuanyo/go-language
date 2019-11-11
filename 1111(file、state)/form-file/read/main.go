package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {

	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	//連接
	http.ListenAndServe(":8080", nil)
}

//function
func foo(w http.ResponseWriter, req *http.Request) {
	var s string
	fmt.Println(req.Method)

	// MethodPost    = "POST"
	if req.Method == http.MethodPost {

		//open
		//返回提供表單的文件
		//func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
		f, h, err := req.FormFile("q")

		//check
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		//information
		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)

		//read
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<br>`+s)
}
