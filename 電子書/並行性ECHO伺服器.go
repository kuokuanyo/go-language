package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second) //並行
	}
	c.Close()
}

func main() {
	//Listen建構監聽網路連線的物件
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		//監聽的Accept會阻斷直到有連線請求進來
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) //例如放棄連線
			continue
		}
		//handleConn處理用戶端連線，在迴圈中輸出時間給用戶端
		go handleConn(conn) //並行處理連線
	}
}
