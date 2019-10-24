package main

import "fmt"

func main() {
	ch := make(chan int) //channel
	//匿名函式並且應用goroutine
	go func() {
		ch <- dosomething(5)
	}() //無參數
	fmt.Println(<-ch)
}

func dosomething(x int) int {
	return x * 5
}
