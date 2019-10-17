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
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

/*應用
package http
type Handler interface { ServeHTTP(w ResponseWriter, req *Request) }
func ListenAndServe(address string, h Handler) error
*/

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
