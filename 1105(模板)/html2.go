package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	//os.Args[0]為文件路徑
	//os.Args[1]為用戶輸入的參數
	name := os.Args[1]
	fmt.Println(os.Args[0]) //文件路徑
	fmt.Println(os.Args[1]) //用戶輸入參數

	str := fmt.Sprint(`
	<!DOCTYPE html>
	<html lang = "en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` +
		name +
		` </h1>
	</body>
	</html>
	`)

	//建立文件
	nf, err := os.Create("index1.html")
	if err != nil {
		log.Fatal("error creating file.")
	}
	defer nf.Close() //檔案關閉

	//複製字串
	io.Copy(nf, strings.NewReader(str))
}
