package main

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func main() {

	//第一個參數為位址，第二個參數為執行函式
	//func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	http.HandleFunc("/", foo)
	//func Handle(pattern string, handler Handler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	//連接
	http.ListenAndServe(":8080", nil)
}

//function
func foo(w http.ResponseWriter, req *http.Request) {
	//尋找cookie
	//func (r *Request) Cookie(name string) (*Cookie, error)
	cookie, err := req.Cookie("session")
	if err != nil {
		//返回隨機生成的uuid
		//func NewV4() (UUID, error)
		id, _ := uuid.NewV4()
		//設定cookie
		//因SetCookie需要使用指標
		cookie = &http.Cookie{
			Name: "session",
			//func (u UUID) String() string
			Value:    id.String(),
			HttpOnly: true,
		}
		//建立設定cookie
		//func SetCookie(w ResponseWriter, cookie *Cookie)
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}
