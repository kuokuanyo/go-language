package main

import (
	"io"
	"net/http"
)

type hotdog int

/*type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
*/
func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(res, "doggy doggy doggy")
	case "/cat":
		io.WriteString(res, "kitty kitty kitty")
	}
}

func main() {
	var d hotdog
	//func ListenAndServe(addr string, handler Handler) error
	http.ListenAndServe(":8080", d)
}
