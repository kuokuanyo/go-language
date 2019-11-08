package main

import (
	"io"
	"log"
	"net"
)

func main() {

	//func Listen(network, address string) (Listener, error)
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	//最後要關閉
	defer li.Close()

	for {
		/*type Listener interface {
		    Accept() (Conn, error)
		    Close() error
		    Addr() Addr
		}*/
		//等待下一個請求
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}

		//func WriteString(w Writer, s string) (n int, err error)
		io.WriteString(conn, "I see you connected.")

		//最後必須關閉
		conn.Close()
	}
}
