package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

//初始化
//func Must(t *Template, err error) *Template
//func ParseFiles(filenames ...string) (*Template, error)
func init() {

	//Must用於變量的初始化
	//如出現錯誤會出現panic
	tpl = template.Must(template.ParseFiles("tpl2.gohtml"))
}

func main() {

	//slice
	sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}

	//func (t *Template) Execute(wr io.Writer, data interface{}) error
	//也可以err := tpl.Execute(os.Stdout 42)
	//func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	err := tpl.ExecuteTemplate(os.Stdout, "tpl2.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}
