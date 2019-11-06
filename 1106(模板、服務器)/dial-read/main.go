package main

import (
	"fmt"
	"net"
)

func main() {

	/*連接到該位址
	func Dial(network, address string) (Conn, error)*/
	conn, err := net.Dial("tcp", "localhost:8080")
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
	defer conn.Close()

	fmt.Fprintln(conn, "I dialed you.")
}
