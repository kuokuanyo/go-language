package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
)

//struct
type User struct {
	ID     string
	Name   string
	Gender string
	Age    int
}

type UserController struct {
	session map[string]User
}

//return pointer
//function  NewUserController
func NewUserController(m map[string]User) *UserController {
	return &UserController{m}
}

func main() {
	var u = make(map[string]User)
	//New return a new router(pointer)
	//func New() *Router
	r := httprouter.New()
	uc := NewUserController(u)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe("localhost:8080", r)
}

//GetUser
//type Handle func(http.ResponseWriter, *http.Request, Params)
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//Byname return the match value
	//func (ps Params) ByName(name string) string
	id := p.ByName("id")

	//get user
	u := uc.session[id]

	//encode(marshal)
	//func Marshal(v interface{}) ([]byte, error)
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	//set header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) //200
	fmt.Fprintf(w, "%s\n", uj)
}

//CreateUser
//type Handle func(http.ResponseWriter, *http.Request, Params)
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := User{}

	//decode
	//NewDecoder return a new decoder
	//func NewDecoder(r io.Reader) *Decoder
	//Decode and store the result in the value pointed to v
	//func (dec *Decoder) Decode(v interface{}) error
	json.NewDecoder(r.Body).Decode(&u)

	//create Id
	Id, _ := uuid.NewV4()
	u.ID = Id.String()

	//store the user
	uc.session[u.ID] = u

	//encode(marshal)
	//func Marshal(v interface{}) ([]byte, error)
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	//set header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) //201
	fmt.Fprintf(w, "%s\n", uj)
}

//DeleteUser
//type Handle func(http.ResponseWriter, *http.Request, Params)
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//Byname return the match value
	//func (ps Params) ByName(name string) string
	id := p.ByName("id")

	//delete
	delete(uc.session, id)

	w.WriteHeader(http.StatusOK) //200
	fmt.Fprintf(w, "Delete User", id, "\n")
}
