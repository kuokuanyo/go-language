package main

import (
	"log"
	"math"
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
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

//function
func double(x int) int {
	return x + x
}

//function square
func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

//function sqroot
func sqroot(x float64) float64 {
	return math.Sqrt(x)
}

//type FuncMap map[string]interface{}
//給Funcs函式使用
var fm = template.FuncMap{
	"fdbl":  double,
	"fsq":   square,
	"fsqrt": sqroot,
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 3)
	if err != nil {
		log.Fatalln(err)
	}
}
