package main

import "fmt"

//xi可以是任意數量的int
func foo(xi ...int) int { //xi的type為[]int
	total := 0
	for _, v := range xi { //xi is slice，因此_為索引
		total += v
	}
	return total
}

func bar(xi []int) int { //xi的type為[]int
	total := 0
	for _, v := range xi { //xi is slice，因此_為索引
		total += v
	}
	return total
}

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := foo(ii...) //一定要將slice使用...
	fmt.Println(n)

	m := bar(ii) //一定要將slice使用...
	fmt.Println(m)
}
