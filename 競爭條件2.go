package main

import (
	"fmt"
	"runtime"
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

	for i := 0; i < gs; i++ {
		//匿名函式並且goroutine
		go func() {
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Println(incrementer)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("final:", incrementer)
}
