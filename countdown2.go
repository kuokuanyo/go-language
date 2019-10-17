//執行火箭發射倒數計時
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) //讀取一個位元組
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second) //channel
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
