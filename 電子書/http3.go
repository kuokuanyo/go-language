//以ServeMux建構URL與相對應處理程序的關係
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
func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

//method
func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		//假設認不得路徑
		w.WriteHeader(http.StatusNotFound) //404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	//type ServeMux struct{ }
	//func NewServeMux() *ServeMux
	mux := http.NewServeMux() //function

	//HandlerFunc是個型別，符合Handler介面的方法函式型別
	//http.HandlerFunc(db.list)與http.HandlerFunc(db.price)是型別轉換而非呼叫函式

	//以下有兩種方法
	//1:
	//func (mux *ServeMux) Handle(pattern string, handler Handler)
	mux.Handle("/list", http.HandlerFunc(db.list))   //db.list: method
	mux.Handle("/price", http.HandlerFunc(db.price)) //db.price: method
	//2:
	//func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	mux.HandleFunc("/list", (db.list))
	mux.HandleFunc("/price", (db.price))

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

/*
package http
type HandlerFunc func(w ResponseWriter, req *Request) 符合Handler介面方法的韓是型別
func (f HandlerFunc) ServeHTTP(w ResponseWriter, req *Request) {
	f(w, r)
}
*/

/*應用
package http
type Handler interface { ServeHTTP(w ResponseWriter, req *Request) }
func ListenAndServe(address string, h Handler) error
*/
