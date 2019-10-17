//當counter的100個元素迴圈後，關閉naturals導致squarer結束迴圈並關閉squares
//關閉上面的naturals會導致平方計算迴圈無止境的接收、處理與發送零值
//沒有方法可以檢測channel是否關閉，要自己建立檢測方式
package main

import "fmt"

func main() {
	naturals := make(chan int) //自然數(channel)
	squares := make(chan int)  //平方(channel)

	//counter
	//goroutine
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals) //執行迴圈後關閉
	}()

	//squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}
}
