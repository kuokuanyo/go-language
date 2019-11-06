package main

import (
	"fmt"
	"io"
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

		//寫入字串
		//func WriteString(w Writer, s string) (n int, err error)
		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!")

		conn.Close()
	}
}
