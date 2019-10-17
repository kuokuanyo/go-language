/*goroutine
程式啟動時，main函式為主goroutine
f() 呼叫f()，等待它被返回
go f() 建構新的goroutine來呼叫f()，不等待
*/
package main

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func main() {
	go spinner(100 * time.Millisecond) //goroutine
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
