package main

import (
	"log"
	"net/http"
)

func main() {
	//func Fatal(v ...interface{})
	//func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error
	//FileServer返回以根目錄為文件內容
	//func FileServer(root FileSystem) Handler
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
