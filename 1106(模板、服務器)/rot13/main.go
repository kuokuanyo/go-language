package main

import (
	"bufio"
	"fmt"
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
		panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}
}

//function
func handle(conn net.Conn) {
	//從r中讀取拆分成心的Scanner
	//func NewScanner(r io.Reader) *Scanner
	scanner := bufio.NewScanner(conn)

	//func (s *Scanner) Scan() bool
	for scanner.Scan() {
		//小寫
		ln := strings.ToLower(scanner.Text())
		//slice
		bs := []byte(ln)
		r := rot13(bs)

		fmt.Fprintf(conn, "%s - %s\n\n", ln, r)
	}
	//最後必須關閉
	defer conn.Close()
}

//function
func rot13(bs []byte) []byte {
	var r13 = make([]byte, len(bs))
	for i, v := range bs {
		if v <= 109 {
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}
	}
	return r13
}
