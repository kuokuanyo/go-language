/*
建立兩個同值但不同型別的欄位
A為字串B為template.HTML
展示html/template套件自動對標題跳脫
*/
package main

import (
	"html/template"
	"os"
	"log"
)

func main() {
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))
	var data struct {
		A string 
		B template.HTML
	}
	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}