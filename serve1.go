package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//連接路徑以/開頭的URL請求處理函式，並啟動伺服器傾聽8000的請求
	http.HandleFunc("/", handler) //每個請求呼叫處理程序
	log.Fatal(http.ListenAndServe("localhost:8000", nil)) 
}

//處理程序回應所請求URL的Path元件
//創建函式
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}