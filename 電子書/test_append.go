package main

import "fmt"

//...讓該函式變可變參數(接受任何數量的參數)
func appendInt(x []int, y ...int) []int {
	var z []int
	//zlen := len(x) + len(y)
	copy(z[len(x):], y)
	return z
}

func main() {
	var x []int
	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	x = append(x, x...)
	fmt.Println(x)
}