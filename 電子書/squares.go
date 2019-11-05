//不具名函式
//回傳一個函式，回傳下一個平方數
package main

import "fmt"

//兩層函式
func squares() func() int {
	var x int
	//回傳 func()
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}