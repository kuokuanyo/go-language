package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

//struct
type sage struct {
	Name  string
	Motto string
}

//struct
type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

//type FuncMap map[string]interface{}
//uc是函式將在模板被呼叫
//ft也是函式
var fm = template.FuncMap{
	"uc": strings.ToUpper, //func ToUpper(s string) string
	"ft": firstThree,
}

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
func firstThree(s string) string {
	s = strings.TrimSpace(s) //TrimSpace將前後多餘空格移除
	s = s[:3]                //取前三個字串
	return s
}

func main() {
	//struct
	b := sage{
		"Buddha", "The belief of no beliefs",
	}
	g := sage{
		"Gandhi", "Be the change",
	}
	m := sage{
		"MLK", "Hatred never ceases with hatred but with love alone is healed.",
	}

	f := car{
		"Ford", "F150", 2,
	}
	c := car{
		"Toyota", "Corolla", 4,
	}

	sages := []sage{b, g, m}
	cars := []car{f, c}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
