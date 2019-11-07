package main

import (
	"log"
	"net/http"
	"text/template"
)

type hotdog int

/*
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
*/
func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//解析表格
	//func (r *Request) ParseForm() error
	err := req.ParseForm()
	//check
	if err != nil {
		log.Fatalln(err)
	}

	//執行模板
	//func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

/*模板
type Template struct {
    *parse.Tree
    // contains filtered or unexported fields
}*/
var tpl *template.Template

//初始化
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
	//func ListenAndServe(addr string, handler Handler) error
	http.ListenAndServe(":8080", d)
}
