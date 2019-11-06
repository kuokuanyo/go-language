package main

import (
	"log"
	"os"
	"text/template"
)

type user struct {
	Name  string
	Motto string
	Admin bool
}

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

	u1 := user{"Buddha", "The belief of no beliefs", false}
	u2 := user{"Gandhi", "Be the change", true}
	u3 := user{"", "Nobody", true}

	users := []user{u1, u2, u3}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", users)
	if err != nil {
		log.Fatalln(err)
	}
}
