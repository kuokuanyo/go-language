package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//廣播程序
//clients紀錄連線中的用戶端
type client chan<- string //發送訊息的channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) //所有用戶端訊息
)

func broadcaster() {
	clients := make(map[client]bool) //所有連線中的用戶端
	for {
		select {
		case msg := <-messages:
			//廣播訊息到所有用戶端的發送訊息channel
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

//handleConn為用戶端建構一個新的發送訊息channel並透過entering對廣播程序宣布其抵達
func handleConn(conn net.Conn) {
	ch := make(chan string) //發送用戶端訊息
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + "has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- ch
	messages <- who + "has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
