package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	//concurrency
	var wg sync.WaitGroup //如果沒使用，每次結果都會不同
	wg.Add(2)             //two goroutine

	//匿名函式，goroutine
	go func() {
		fmt.Println("hello from thing one.")
		wg.Done()
	}()
	//匿名函式，goroutine
	go func() {
		fmt.Println("hello from thing two.")
		wg.Done()
	}()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("about to exit.")

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
}
