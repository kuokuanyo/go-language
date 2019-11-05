package main

import "fmt"

func foo(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}

func main() {
	//創建foo的第一個參數
	g := func(xi []int) int {
		if len(xi) == 0 {
			return 1
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi)-1]
	}
	x := foo(g, []int{1, 2, 3, 4, 5})
	fmt.Println(x)
}
