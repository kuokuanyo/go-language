//http1不論請求的URL為何，都會列出所有的庫存
//定義一個伺服器，會依照不同的URL觸發不同的行為
package main

import (
	"fmt"
	"log"
	"net/http"
)

//價格
type dollars float32

//method
func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

//map
type database map[string]dollars

//method
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	//列出所有產品
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			//假設認不得路徑
			w.WriteHeader(http.StatusNotFound) //404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound) //404
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
