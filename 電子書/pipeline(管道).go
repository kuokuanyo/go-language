/*channel可用於連接goroutine，使一個goroutine輸出作為另一個的goroutine輸入，稱為管道
以下程式展示第一個goroutine產生自然數並透過channel發送給第二個goroutine，接受值後計算平方，
然後透過另一個channel發送給第三個goroutine輸出
*/
package main

import "fmt"

func main() {
	naturals := make(chan int) //自然數(channel)
	squares := make(chan int)  //平方(channel)

	//counter
	//goroutine
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	//squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	for {
		fmt.Println(<-squares)
	}
}
