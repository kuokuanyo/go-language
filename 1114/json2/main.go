package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type User struct {
	Name   string
	Gender string
	Age    int
	Id     string
}

func main() {
	//New return a new initialized router
	//func New() *Router
	r := httprouter.New()
	//func (r *Router) GET(path string, handle Handle)
	r.GET("/", index)
	r.GET("/user/:id", getUser)
	//connect
	http.ListenAndServe("localhost:8080", r)
}

//function index
//type Handle func(http.ResponseWriter, *http.Request, Params)
func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
			<html lang="en">
			<head>
			<meta charset="UTF-8">
			<title>Index</title>
			</head>
			<body>
			<a href="/user/9872309847">GO TO: http://localhost:8080/user/9872309847
			</body>
			</html>`
	//Set the header
	//func (h Header) Set(key, value string)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	//WriteHeader(statusCode int)
	w.WriteHeader(http.StatusOK)
	//write
	//Write([]byte) (int, error)
	w.Write([]byte(s))
}

//function getUser
func getUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := User{
		Name:   "James Bond",
		Gender: "male",
		Age:    32,
		//return the match value
		//func (ps Params) ByName(name string) string
		Id: p.ByName("id"),
	}

	//encode(marshal)
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s\n", uj)
}
