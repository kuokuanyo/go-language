package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	//創建新模板，並從命名文件中取得模板定義
	//func ParseFiles(filenames ...string) (*Template, error)
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	//func (t *Template) Execute(wr io.Writer, data interface{}) error
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
