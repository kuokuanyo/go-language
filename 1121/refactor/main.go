package main

import (
	"database/sql"
	"log"
	"net/http"
	"refactor/controllers"
	"refactor/driver"

	"github.com/gorilla/mux"
)

var db *sql.DB

func main() {
	//connect db
	db = driver.Connect()

	//type controller
	controller := controllers.Controller{}

	//create router
	//func NewRouter() *Router
	router := mux.NewRouter() //*Router

	//func (r *Router) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *Route
	//func (r *Router) Methods(methods ...string) *Route
	router.HandleFunc("/signup", controller.Signup(db)).Methods("POST")
	router.HandleFunc("/login", controller.Login(db)).Methods("POST")
	//被保護端點
	router.HandleFunc("/protected", controller.TokenVerifyMiddleWare(controller.ProtectedEndpoint())).Methods("GET")

	//connect server
	//log.Fatal record error
	log.Fatal(http.ListenAndServe(":8080", router))
}
