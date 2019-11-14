package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	//New return a new initialized router
	//func New() *Router
	r := httprouter.New()
	//func (r *Router) GET(path string, handle Handle)
	r.GET("/", index)
	//connect
	http.ListenAndServe("localhost:8080", r)
}

//type Handle func(http.ResponseWriter, *http.Request, Params)
func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
