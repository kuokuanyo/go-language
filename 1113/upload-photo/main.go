package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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

func index(w http.ResponseWriter, req *http.Request) {
	//get cookie
	c := getCookie(w, req)
	if req.Method == http.MethodPost {
		//func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
		//mf is  file, fh is header
		mf, fh, err := req.FormFile("nf")
		if err != nil {
			fmt.Println(err)
		}
		//close file
		defer mf.Close()
		//create sha for file name
		//string split
		//take index 1(file type)
		ext := strings.Split(fh.Filename, ".")[1]
		//加密
		//func New() hash.Hash
		h := sha1.New()
		//copy from src(content) to dst
		//func Copy(dst Writer, src Reader) (written int64, err error)
		io.Copy(h, mf)
		//file name
		//%x十六進位
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext
		//create new file
		//Getwd returns root path
		//func Getwd() (dir string, err error)
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		//join fliepath
		//func Join(elem ...string) string
		path := filepath.Join(wd, "public", "pics", fname)
		//create new file
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		//close nf
		defer nf.Close()
		//copy
		//func (f *File) Seek(offset int64, whence int) (ret int64, err error)
		//whence is offset location
		mf.Seek(0, 0)
		//copy from mf(content) to nf
		io.Copy(nf, mf)
		//appendValue
		c = appendValue(w, c, fname)
	}
	//string split
	xs := strings.Split(c.Value, "|")
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
			Name:  "session",
			Value: sID.String(),
		}
		//set new cookie
		http.SetCookie(w, c)
	}
	return c
}

func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
	s := c.Value
	//check contain
	if !strings.Contains(s, fname) {
		s += "|" + fname
	}
	c.Value = s
	//set cookie
	http.SetCookie(w, c)
	return c
}
