package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

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

func main() {
	var u = make(map[string]User)
	//return a new router(pointer)
	r := httprouter.New()
	uc := NewUserController(u)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

//return pointer
//function  NewUserController
func NewUserController(m map[string]User) *UserController {
	return &UserController{m}
}

func StoreUsers(m map[string]User) {
	//create file
	f, err := os.Create("data")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	//encode
	//func NewEncoder(w io.Writer) *Encoder
	//func (enc *Encoder) Encode(v interface{}) error
	json.NewEncoder(f).Encode(m)
}

func LoadUsers() map[string]User {
	m := make(map[string]User)
	//open file
	f, err := os.Open("data")
	if err != nil {
		fmt.Println(err)
		return m
	}
	defer f.Close()
	//decode
	//Decode參數必須使用指標，否則會出現錯誤
	err = json.NewDecoder(f).Decode(&m)
	if err != nil {
		fmt.Println(err)
	}
	return m
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	u := uc.session[id]
	//encode(marshal)
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) //200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := User{}
	//decode
	//Decode參數必須是指標
	json.NewDecoder(r.Body).Decode(&u)
	//create id
	i, _ := uuid.NewV4()
	u.ID = i.String()
	uc.session[u.ID] = u
	//store
	StoreUsers(uc.session)
	//encode(marshal)
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) //201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	delete(uc.session, id)
	//store
	StoreUsers(uc.session)
	w.WriteHeader(http.StatusOK) //200
	fmt.Fprintf(w, "Delete user", id, "\n")
}
