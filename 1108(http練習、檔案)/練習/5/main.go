package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
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

		scanner := bufio.NewScanner(conn)
		//func (s *Scanner) Scan() bool
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
		}
		defer conn.Close()

		//沒離開迴圈，因此不會執行下列兩行
		fmt.Println("code got here.")
		io.WriteString(conn, "connected.")
	}
}
