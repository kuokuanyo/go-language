package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		serve(conn)
	}
}

//function
func serve(c net.Conn) {
	defer c.Close()
	scanner := bufio.NewScanner(c)
	var i int
	var rMethod, rURL string
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		//開頭
		if i == 0 {
			//以一個或多個空格拆分
			xs := strings.Fields(ln)
			rMethod = xs[0]
			rURL = xs[1]
			fmt.Println("METHOD:", rMethod)
			fmt.Println("URL:", rURL)
		}

		//結束迴圈
		if ln == "" {
			fmt.Println("This is the end of the http request headers.")
			break
		}
		i++
	}
	body := `
	<!DOCTYPE html>
	<html lang = "en">
	<head>
			<meta charset="UTF-8">
			<title>code gangsta</title>
	</head>
	<body>
			<h1>"holy cow this is low level"</h1>
	</body>
	</html>
	`

	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}
