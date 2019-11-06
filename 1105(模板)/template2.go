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

	//建立檔案
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	//要關閉檔案
	defer nf.Close()

	//寫入nf
	//func (t *Template) Execute(wr io.Writer, data interface{}) error
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
