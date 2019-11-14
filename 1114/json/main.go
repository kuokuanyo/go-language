package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//struct
type person struct {
	Fname string
	Lname string
	Items []string
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/mshl", mshl)
	http.HandleFunc("/encd", encd)
	//connect
	//func ListenAndServe(addr string, handler Handler) error
	http.ListenAndServe(":8080", nil)
}

//function foo
func foo(w http.ResponseWriter, req *http.Request) {
	//s is html
	s := `<!DOCTYPE html>
			<heml lang="en">
			<head>
			<meta charset="UTF-8">
			<title>FOO</title>
			</head>
			<body>
			You are at foo
			</body>
			</html>`
	//write
	w.Write([]byte(s))
}

//function mshl
func mshl(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		"Kuo",
		"Yu",
		[]string{"Suit", "Gun", "Wry sense of humor"},
	}
	//func Marshal(v interface{}) ([]byte, error)
	//encode json
	json, err := json.Marshal(p1)
	//if error
	if err != nil {
		log.Panicln(err)
	}
	//write
	w.Write(json)
}

//function encd
func encd(w http.ResponseWriter, req *http.Request) {
	//w.Header() is type of header
	//Set the key value of header
	//func (h Header) Set(key, value string)
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		"Kuo",
		"Yu",
		[]string{"Suit", "Gun", "Wry sense of humor"},
	}
	//func NewEncoder(w io.Writer) *Encoder
	//NewEncoder return new encoder that write to w
	//func (enc *Encoder) Encode(v interface{}) error
	//Encode json
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Panicln(err)
	}
}
