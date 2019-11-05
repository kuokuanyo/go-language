package main

import "fmt"

func main() {
	//channel
	c := make(chan int)

	//send並且匿名函式
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	//receive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit.")
}
