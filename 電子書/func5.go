package main

import "fmt"

//雙重函式
func foo() func() int {
	return func() int {
		return 42
	}
}

func main() {
	f := foo()
	fmt.Println(f())
}
