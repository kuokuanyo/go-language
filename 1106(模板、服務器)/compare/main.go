package main

import (
	"log"
	"os"
	"text/template"
)

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
	//func (t *Template) ParseFiles(filenames ...string) (*Template, error)
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	//struct
	g1 := struct {
		Score1 int
		Score2 int
	}{
		7,
		9,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", g1)
	if err != nil {
		log.Fatalln(err)
	}
}
