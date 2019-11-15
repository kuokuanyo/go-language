package session

import (
	"fmt"
	"models"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

const Length int = 30

var Users = map[string]models.User{}
var Sessions = map[string]models.Session{}
var LastCleaned time.Time = time.Now()

func GetUser(w http.ResponseWriter, req *http.Request) models.User {
	//get cookie
	c, err := req.Cookie("session")
	if err != nil {
		//create id
		sID, _ := uuid.NewV4()
		//set cookie
		c = &http.Cookie{
			Name: "session",
			//convert string
			Value: sID.String(),
		}
	}
	c.MaxAge = Length
	//set new cookie
	http.SetCookie(w, c)

	var u models.User
	if s, ok := Sessions[c.Value]; ok {
		s.LastActivity = time.Now()
		Sessions[c.Value] = s
		u = Users[s.UserName]
	}
	return u
}

func AlreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	// not log
	if err != nil {
		return false
	}

	s, ok := Sessions[c.Value]
	if ok {
		s.LastActivity = time.Now()
		Sessions[c.Value] = s
	}

	_, ok = Users[s.UserName]
	c.MaxAge = Length
	http.SetCookie(w, c)
	return ok
}

func Clean() {
	fmt.Println("BEFORE CLEAN")
	Show()
	for k, v := range Sessions {
		if time.Now().Sub(v.LastActivity) > (time.Second * 30) {
			//delete
			delete(Sessions, k)
		}
	}
	LastCleaned = time.Now()
	fmt.Println("AFTER CLEAN")
	Show()
}

func Show() {
	fmt.Println("**********")
	for k, v := range Sessions {
		fmt.Println(k, v.UserName)
	}
	fmt.Println("")
}
