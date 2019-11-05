//template.Must輔助程式讓錯誤處理更方便
package main

import (
		"github"
	 	"os"
	 	"log"
	 	"text/template"
	 	"time"
	 	)

//樣板
const templ = `{{.TotalCount}} issues:
{{range .Items}}-------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}`

//來自text/template的New函式建構並回傳一個樣板
var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}