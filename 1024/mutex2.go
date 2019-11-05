package main

import (
	"fmt"
	"sync"
)

func main() {
	//concurrency
	var wg sync.WaitGroup

	//計數器
	incrementer := 0
	//100個goroutine
	gs := 100

	wg.Add(gs)

	//互斥鎖
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		//匿名函式並且goroutine
		go func() {
			mu.Lock()
			v := incrementer
			v++
			incrementer = v
			fmt.Println(incrementer)
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("final:", incrementer)
}
