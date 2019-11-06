package main

import (
	"log"
	"os"
	"text/template"
)

//struct
type course struct {
	Number, Name, Units string
}

//semester
type semester struct {
	Term    string
	Courses []course
}

//struct
type year struct {
	Fall, Spring, Summer semester
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
	//ParseGlob解析路徑檔案
	//func ParseGlob(pattern string) (*Template, error)
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{"CSCI-40", "Introduction to Programming in Go", "4"},
				course{"CSCI-130", "Introduction to Web Programming in Go", "4"},
				course{"CSCI-140", "Mobile", "4"},
			},
		},
		Spring: semester{
			Term: "Fall",
			Courses: []course{
				course{"CSCI-50", "Advanced Go", "5"},
				course{"CSCI-190", "Advanced Web Programming in Go", "5"},
				course{"CSCI-191", "Advanced Mobile", "5"},
			},
		},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", y)
	if err != nil {
		log.Fatalln(err)
	}
}
