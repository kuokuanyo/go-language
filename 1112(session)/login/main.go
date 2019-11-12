package main

import (
	"net/http"
	"text/template"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

//struct
type user struct {
	UserName string
	Password []byte
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
	//密碼加密
	//func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	bs, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	//加入資料
	dbUsers["a167829435@gmail.com"] = user{"a167829435@gmail.com", bs, "kuo", "yu"}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	//connect
	http.ListenAndServe(":8080", nil)
}

//function index
func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	//execute
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	//假設未登入
	if !alreadyLoggedIn(req) {
		//重新導向回http://localhost:8080/
		//func Redirect(w ResponseWriter, r *Request, url string, code int)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	//執行模板
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

//function signup
func signup(w http.ResponseWriter, req *http.Request) {
	//已經登入
	if alreadyLoggedIn(req) {
		//重新導向回http://localhost:8080/
		//func Redirect(w ResponseWriter, r *Request, url string, code int)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		//func (r *Request) FormValue(key string) string
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

		//username被使用過
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		//隨機生成新的uuid
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name: "session",
			//轉換成字串
			Value: sID.String(),
		}
		//設定新cookie
		http.SetCookie(w, c)

		//密碼處理
		//func GenerateFromPassword(password []byte, cost int) ([]byte, error)
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		//check
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		//add value
		dbSessions[c.Value] = un
		u := user{un, bs, f, l}
		dbUsers[un] = u

		//重新導向回http://localhost:8080/
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	//execute
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

//function getUser 取得user(struct)
func getUser(req *http.Request) user {
	var u user
	//尋找cookie(session)
	c, err := req.Cookie("session")
	//沒有session
	if err != nil {
		return u
	}

	//假設user已經存在
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

//function alreadyLoggedIn
func alreadyLoggedIn(req *http.Request) bool {
	//尋找cookie
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}

//function login
func login(w http.ResponseWriter, req *http.Request) {
	//假設已經登入
	//重新導向
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u user
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")

		//是否有此帳號
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		//檢查密碼是否符合
		//func CompareHashAndPassword(hashedPassword, password []byte) error
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		//創建uuid
		sID, err := uuid.NewV4()
		//設定cookie
		c := &http.Cookie{
			Name: "session",
			//轉換字串
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = un
		//重新導向回http://localhost:8080/
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "login.gohtml", u)
}
