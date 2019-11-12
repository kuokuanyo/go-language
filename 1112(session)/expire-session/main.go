package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

//struct
type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

//struct
type session struct {
	un           string
	lastActivity time.Time
}

//模板
var tpl *template.Template

//user id, user
var dbUsers = map[string]user{}

//session id, user id
var dbSessions = map[string]session{}
var dbSessionsCleaned time.Time

const sessionLength int = 30

//初始化
func init() {
	//Must模板初始化
	//ParseGlob解析路徑並創建模板
	tpl = template.Must(template.ParseGlob("templates/*"))
	dbSessionsCleaned = time.Now()
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	//connect
	http.ListenAndServe(":8080", nil)
}

//function index
func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	showSessions()
	//execute
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	//假設未登入
	if !alreadyLoggedIn(w, req) {
		//重新導向回http://localhost:8080/
		//func Redirect(w ResponseWriter, r *Request, url string, code int)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	showSessions()
	//執行模板
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

//function signup
func signup(w http.ResponseWriter, req *http.Request) {
	//已經登入
	if alreadyLoggedIn(w, req) {
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
		r := req.FormValue("role")

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
		c.MaxAge = sessionLength
		//設定新cookie
		http.SetCookie(w, c)

		//add value
		dbSessions[c.Value] = session{un, time.Now()}

		//密碼處理
		//func GenerateFromPassword(password []byte, cost int) ([]byte, error)
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		//check
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		//add value
		u := user{un, bs, f, l, r}
		dbUsers[un] = u

		//重新導向回http://localhost:8080/
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	showSessions()
	//execute
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

//function getUser 取得user(struct)
func getUser(w http.ResponseWriter, req *http.Request) user {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u user
	if s, ok := dbSessions[c.Value]; ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
		u = dbUsers[s.un]
	}
	return u
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := dbSessions[c.Value]
	if ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
	}
	_, ok = dbUsers[s.un]
	// refresh session
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	return ok
}

//function login
func login(w http.ResponseWriter, req *http.Request) {
	//假設已經登入
	//重新導向
	if alreadyLoggedIn(w, req) {
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
		c.MaxAge = sessionLength
		http.SetCookie(w, c)
		dbSessions[c.Value] = session{un, time.Now()}
		//重新導向回http://localhost:8080/
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w, "login.gohtml", u)
}

//function logout
func logout(w http.ResponseWriter, req *http.Request) {
	//假設為登入
	if !alreadyLoggedIn(w, req) {
		//重新導向
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	//尋找cookie
	c, _ := req.Cookie("session")
	//刪除session
	delete(dbSessions, c.Value)
	//刪除cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	//clean up
	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
		go cleanSessions()
	}
	http.Redirect(w, req, "/login", http.StatusSeeOther)
}

//func showSessions
func showSessions() {
	fmt.Println("***********")
	for k, v := range dbSessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}

//func cleanSessions
func cleanSessions() {
	fmt.Println("BEFORE CLEAN")
	showSessions()
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			//刪除
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN")
	showSessions()
}
