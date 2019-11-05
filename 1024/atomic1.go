package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	//concurrency
	var wg sync.WaitGroup
	var incrementer int64

	//100個goroutine
	gs := 100

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		//匿名函式並且goroutine
		go func() {
			atomic.AddInt64(&incrementer, 1)
			fmt.Println(atomic.LoadInt64(&incrementer))
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("final:", incrementer)
}
