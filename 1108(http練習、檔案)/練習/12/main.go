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
		go serve(conn)
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

	switch {
	case rMethod == "GET" && rURL == "/":
		index(c)
	case rMethod == "GET" && rURL == "/apply":
		apply(c)
	case rMethod == "POST" && rURL == "/apply":
		applypost(c)
	default:
		d(c)
	}
}

func index(c net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang = "en">
	<head>
			<meta charset="UTF-8">
			<title>get index</title>
	</head>
	<body>
			<h1>"GET INDEX"</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
	</body>
	</html>
	`

	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}

func apply(c net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang = "en">
	<head>
			<meta charset="UTF-8">
			<title>get dog</title>
	</head>
	<body>
			<h1>"GET APPLY"</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
			<form action="/apply" method="POST">
			<input type="hidden" value="in my good death">
			<input type="submit" value="submit">
			</form>
	</body>
	</html>
	`

	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}

func applypost(c net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang = "en">
	<head>
			<meta charset="UTF-8">
			<title>post apply</title>
	</head>
	<body>
			<h1>"POST APPLY"</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
	</body>
	</html>
	`

	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}

func d(c net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang = "en">
	<head>
			<meta charset="UTF-8">
			<title>default</title>
	</head>
	<body>
			<h1>"default"</h1>
	</body>
	</html>
	`

	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}
