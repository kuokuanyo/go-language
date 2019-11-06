package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
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
		log.Panic(err)
	}
	defer li.Close()

	for {
		//等待並回傳下一個連接
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		//goroutine
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
		ln := scanner.Text()
		fmt.Println(ln)
	}
	//最後必須關閉
	defer conn.Close()
}
