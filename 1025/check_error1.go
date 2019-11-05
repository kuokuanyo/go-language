package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	//延遲函式（匿名）
	//無引數傳遞
	defer func() {
		//假設正常執行，recover返回nil
		if r := recover(); r != nil { //修復
			fmt.Println("Recoverd in f", r)
		}
	}() //無參數
	fmt.Println("Calling g.")
	g(0)                                     //會出現panic
	fmt.Println("Returned normally from g.") //跳到defer函式，因此沒被執行
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	//defer函式輸出順序：先進後出
	//defer傳遞參數，保持當下的參數
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
