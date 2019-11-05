//週期性輸出時間的TCP伺服器
package main

import (
	"io"
	"log"
	"net"
	"time"
)

/*type Conn interface {Read(b []byte) (n int, err error)
Write(b []byte) (n int, err error)
Close() error
...
*/
func handleConn(c net.Conn) {
	defer c.Close() //延遲
	for {
		//func WriteString(w Writer, s string) (n int, err error)
		//time.時間.Format提供日期與時間資訊的格式化
		_, err := io.WriteString(c, time.Now().Format("10:34:00\n"))
		if err != nil {
			return //例如斷線
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	//Listen建構監聽網路連線的物件
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		//監聽的Accept直到有連線請求進來會被阻斷
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) //例如放棄連線
			continue
		}
		//handleConn處理用戶端連線，在迴圈中輸出時間給用戶端
		go handleConn(conn) //並行處理連線
	}
}
