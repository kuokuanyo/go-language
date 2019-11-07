package main

import (
	"log"
	"net/http"
	"net/url"
	"text/template"
)

type hotdog int

/*
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
*/
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//解析表格
	//func (r *Request) ParseForm() error
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method        string
		URL           *url.URL
		Submissions   map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		r.Method,
		r.URL,
		r.Form,
		r.Header,
		r.Host,
		r.ContentLength,
	}

	//執行模板
	//func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	//Must用於變量的初始化
	//如出現錯誤會出現panic
	//func Must(t *Template, err error) *Template

	//New分配名稱給該模板
	//func New(name string) *Template

	//加入功能給模板
	//Funcs必須要解析模板之前使用
	//func (t *Template) Funcs(funcMap FuncMap) *Template

	//創建模板並解析模板定義(Funcs須在此函式之前使用)
	//ParseGlob解析路徑檔案
	//func ParseGlob(pattern string) (*Template, error)
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
