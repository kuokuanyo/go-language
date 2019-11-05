package main

import "fmt"

func foo() {
	defer func() { //延遲函式(匿名函式))
		fmt.Println("defer ran")
	}() //無參數(匿名函式)
	fmt.Println("ran")
}

func main() {
	defer foo() //延遲
	fmt.Println("Hello")
}
