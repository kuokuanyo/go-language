package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

//function index
func index(w http.ResponseWriter, req *http.Request) {
	//get cookie
	c := getCookie(w, req)
	//execute template
	tpl.ExecuteTemplate(w, "index.gohtml", c.Value)
}

//function getCookie
func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	//find cookie
	c, err := req.Cookie("session")
	//if no
	if err != nil {
		//create uuid
		sID, _ := uuid.NewV4()
		//set cookie value
		c = &http.Cookie{
			Name: "session",
			//convert string
			Value: sID.String(),
		}
		//set new cookie
		http.SetCookie(w, c)
	}
	return c
}
