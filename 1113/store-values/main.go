package main

import (
	"net/http"
	"strings"
	"text/template"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/", index)
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	//get cookie
	c := getCookie(w, req)
	//appendvalue
	c = appendValue(w, c)
	//split string(c.Value)
	//slice
	xs := strings.Split(c.Value, "|")
	//execute template
	tpl.ExecuteTemplate(w, "index.gohtml", xs)
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	//find cookie
	c, err := req.Cookie("session")
	//if no
	if err != nil {
		//create uuid
		sID, _ := uuid.NewV4()
		//set cookie
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

func appendValue(w http.ResponseWriter, c *http.Cookie) *http.Cookie {
	//values
	p1 := "disneyland.jpg"
	p2 := "atbeach.jpg"
	p3 := "hollywood.jpg"
	//append
	s := c.Value
	//check contains
	if !strings.Contains(s, p2) {
		s += "|" + p1
	}
	if !strings.Contains(s, p2) {
		s += "|" + p2
	}
	if !strings.Contains(s, p3) {
		s += "|" + p3
	}
	c.Value = s
	//set cookie
	http.SetCookie(w, c)
	return c
}
