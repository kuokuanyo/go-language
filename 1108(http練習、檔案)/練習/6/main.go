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
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)

			//結束迴圈
			if ln == "" {
				fmt.Println("This is the end of the http request headers.")
				break
			}
		}
		defer conn.Close()

		//會執行下列兩行
		fmt.Println("code got here.")
		io.WriteString(conn, "connected.")
	}
}
