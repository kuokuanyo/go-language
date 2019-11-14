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
	//New return a new initialized router(pointer)
	//func New() *Router
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", getUser)
	//add route
	//func (r *Router) POST(path string, handle Handle)
	r.POST("/user", createUser)
	//func (r *Router) DELETE(path string, handle Handle)
	r.DELETE("/user/:id", deleteUser)
	//connect
	http.ListenAndServe("localhost:8080", r)
}

//function index
//type Handle func(http.ResponseWriter, *http.Request, Params)
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
	w.WriteHeader(http.StatusOK) //200
	//write
	//Write([]byte) (int, error)
	w.Write([]byte(s))
}

//function getUser
//type Handle func(http.ResponseWriter, *http.Request, Params)
func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
	w.WriteHeader(http.StatusOK) //200
	fmt.Fprintf(w, "%s\n", uj)
}

//function createUser
//type Handle func(http.ResponseWriter, *http.Request, Params)
func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := User{}

	//decode
	//NewDecoder return a new decoder(pointer)
	//func NewDecoder(r io.Reader) *Decoder
	//Decode的參數必須為指標，否則會出現錯誤
	//decode and store the result in v(pointer)
	//func (dec *Decoder) Decode(v interface{}) error
	json.NewDecoder(r.Body).Decode(&u)

	//change Id
	u.Id = "007"

	//encode(marshal)
	//func Marshal(v interface{}) ([]byte, error)
	uj, _ := json.Marshal(u)

	//write content-type, statuscode, payload
	//set header
	//func (h Header) Set(key, value string)
	w.Header().Set("Content-Type", "application/json")
	//WriteHeader(statusCode int)
	w.WriteHeader(http.StatusCreated) //201
	fmt.Fprintf(w, "%s\n", uj)
}

//function deleteUser
//type Handle func(http.ResponseWriter, *http.Request, Params)
func deleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK) //200
	fmt.Fprintf(w, "Write code to delete user\n")
}
