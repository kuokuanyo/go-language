package main

import(
	"fmt"
	"os"
)

//可變函式(通常用於字串格式化)
//...表示函式可用任意數量的參數呼叫
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}

func main() {
	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(1, 2, 3, 4))

	//slice
	values := []int{1, 2, 3, 4} 
	fmt.Println(sum(values...))

	linenum, name := 12, "count"
	errorf(linenum, "undefied: %s", name)
}