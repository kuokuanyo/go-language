package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send(goroutine)
	go send(even, odd, quit)

	//receive
	receive(even, odd, quit)

	fmt.Println("about to exit.")
}

//send function
func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 { //偶數
			e <- i
		} else { //奇數
			o <- i
		}
	}
	close(e)
	close(o)
	q <- 0 //只傳遞一個數，不用close
	close(q)
}

//receive function
func receive(e, o, q <-chan int) {
	for { //無限循環
		select {
		case v := <-e:
			fmt.Println("from the even channel:", v)
		case v := <-o:
			fmt.Println("from the odd channel:", v)
		case v := <-q:
			fmt.Println("from the quit channel:", v)
			return
		}
	}
}
