package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	/*func Listen(network, address string) (Listener, error)
	network必須是"tcp", "tcp4", "tcp6", "unix" or "unixpacket"
	*/
	li, err := net.Listen("tcp", ":8080")
	/*type Listener interface {
	    等待並回傳下一個連接
	    Accept() (Conn, error)
		結束一定要關閉
	    Close() error
	    // Addr returns the listener's network address.
	    Addr() Addr
		}*/

	if err != nil {
		//func Fatalln(v ...interface{})
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		//等待並回傳下一個連接
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		//goroutine
		go handle(conn)
	}
}

//function
func handle(conn net.Conn) {
	//最後必須關閉
	defer conn.Close()

	//讀取請求
	request(conn)
}

//function
func request(conn net.Conn) {
	i := 0
	//從r中讀取拆分成心的Scanner
	//func NewScanner(r io.Reader) *Scanner
	scanner := bufio.NewScanner(conn)
	//func (s *Scanner) Scan() bool
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 { //第一次請求
			mux(conn, ln)
		}
		if ln == "" {
			//讀取請求完畢
			break
		}
		i++
	}
}

//function
func mux(conn net.Conn, ln string) {
	//請求
	//Fields()以一個或多個空格拆解成字串
	//func Fields(s string) []string
	m := strings.Fields(ln)[0]
	u := strings.Fields(ln)[1]
	fmt.Println("***METHOD", m)
	fmt.Println("***URL", u)

	if m == "GET" && u == "/" {
		index(conn)
	}
	if m == "GET" && u == "/about" {
		about(conn)
	}
	if m == "GET" && u == "/contact" {
		contact(conn)
	}
	if m == "GET" && u == "/apply" {
		apply(conn)
	}
	if m == "POST" && u == "/apply" {
		applyProcess(conn)
	}
}

//function
func index(conn net.Conn) {
	body := `<!DOCTYPE html><html lang = "en"><head><meta
	charset="UTF-8"><title></title></head><body><strong>INDEX</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`

	//寫入conn
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

//function
func about(conn net.Conn) {
	body := `<!DOCTYPE html><html lang = "en"><head><meta
	charset="UTF-8"><title></title></head><body><strong>ABOUT</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`

	//寫入conn
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

//function
func contact(conn net.Conn) {
	body := `<!DOCTYPE html><html lang = "en"><head><meta
	charset="UTF-8"><title></title></head><body><strong>CONTACT</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`

	//寫入conn
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

//function
func apply(conn net.Conn) {
	body := `<!DOCTYPE html><html lang = "en"><head><meta
	charset="UTF-8"><title></title></head><body><strong>APPLY</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	<form method="post" action="/apply">
	<input type="submit" value="apply">
	</form>
	</body></html>`

	//寫入conn
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

//function
func applyProcess(conn net.Conn) {
	body := `<!DOCTYPE html><html lang = "en"><head><meta
	charset="UTF-8"><title></title></head><body><strong>APPLY PROCESS</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`

	//寫入conn
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
