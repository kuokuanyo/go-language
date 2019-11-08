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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="get">
		<input type="text" name="q">
		<input type="submit">
	</form>
	<br>`+v)
}

//method=get會改變url(網址)
//method=post只會改變內容
//http://localhost:8080/?q=查詢字串
