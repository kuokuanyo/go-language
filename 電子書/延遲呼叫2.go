package main

import(
	"fmt"
)

func double(x int) (result int) { //return結果變數名稱result
	defer func() { fmt.Printf("double(%d) = %d\n", x, result)}() //defer後面要有括號
	return x + x
}

func main() {
	double(4)
}