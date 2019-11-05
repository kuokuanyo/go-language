//互斥鎖
//防止競爭條件
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	//計數器
	counter := 0
	//數量
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex //互斥鎖

	for i := 0; i < gs; i++ {
		//匿名函式並且使用goroutine
		go func() {
			mu.Lock()
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}() //無參數
		fmt.Println("Goroutines\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
