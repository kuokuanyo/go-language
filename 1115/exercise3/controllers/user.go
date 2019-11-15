package controllers

import (
	"net/http"
	"session"
	"time"

	"golang.org/x/crypto/bcrypt"

	"models"

	uuid "github.com/satori/go.uuid"
)

func (c Controller) SignUp(w http.ResponseWriter, req *http.Request) {
	//假設已經登入
	if session.AlreadyLoggedIn(w, req) {
		//重新導向至localhost:8080
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	var u models.User

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")

		//username已被使用
		if _, ok := session.Users[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		//create id
		sID, _ := uuid.NewV4()
		//set cookie
		c := &http.Cookie{
			Name: "session",
			//convert string
			Value: sID.String(),
		}
		c.MaxAge = session.Length
		//set new cookie
		http.SetCookie(w, c)
		session.Sessions[c.Value] = models.Session{un, time.Now()}

		//加密
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		u = models.User{un, bs, f, l, r}
		session.Users[un] = u
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	session.Show()
	c.tpl.ExecuteTemplate(w, "signup.gohtml", u)
}

func (c Controller) Login(w http.ResponseWriter, req *http.Request) {
	if session.AlreadyLoggedIn(w, req) {
		//重新導向至localhost:8080
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u models.User

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		//假設不符合資訊
		u, ok := session.Users[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		//check password
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		//create id
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = session.Length
		http.SetCookie(w, c)
		session.Sessions[c.Value] = models.Session{un, time.Now()}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	session.Show()
	c.tpl.ExecuteTemplate(w, "login.gohtml", u)
}

func (c Controller) Logout(w http.ResponseWriter, req *http.Request) {
	if !session.AlreadyLoggedIn(w, req) {
		//重新導向至localhost:8080
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	//find cookie
	ck, _ := req.Cookie("session")

	delete(session.Sessions, ck.Value)

	ck = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1, //-1 is delete
	}
	http.SetCookie(w, ck)

	if time.Now().Sub(session.LastCleaned) > (time.Second * 30) {
		go session.Clean()
	}

	http.Redirect(w, req, "/login", http.StatusSeeOther)
}
