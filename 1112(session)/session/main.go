package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

//struct
type user struct {
	UserName string
	First    string
	Last     string
}

//模板
var tpl *template.Template

//user id, user
var dbUsers = map[string]user{}

//session id, user id
var dbSessions = map[string]string{}

//初始化
func init() {
	//Must模板初始化
	//ParseGlob解析路徑並創建模板
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	//connect
	http.ListenAndServe(":8080", nil)
}

//function
func foo(w http.ResponseWriter, req *http.Request) {

	//尋找cookie
	c, err := req.Cookie("session")
	//check
	if err != nil {
		//創建一個uuid
		////func NewV4() (UUID, error)
		sID, _ := uuid.NewV4()
		//設定cookie
		c = &http.Cookie{
			Name: "session",
			//轉換字串
			//func (u UUID) String() string
			Value: sID.String(),
		}
		//設定新cookie
		//func SetCookie(w ResponseWriter, cookie *Cookie)
		http.SetCookie(w, c)
	}

	//假設user存在，取得user
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}

	if req.Method == http.MethodPost {
		//查詢命名組件
		//func (r *Request) FormValue(key string) string
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		u = user{un, f, l}

		//新增map value
		dbSessions[c.Value] = un
		dbUsers[un] = u
	}

	//執行模板
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

//function
func bar(w http.ResponseWriter, req *http.Request) {
	//尋找cookie
	c, err := req.Cookie("session")
	if err != nil {
		//重新導向回到http://localhost:8080/
		//func Redirect(w ResponseWriter, r *Request, url string, code int)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
