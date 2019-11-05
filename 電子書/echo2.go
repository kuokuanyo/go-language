package main

import(
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	//_為空辨識符，用於語法需要但程式邏輯不需要變數的情況
	//假設跟echo1一樣設i := 1，後面沒使用到i，會出現錯誤
	for _, arg := range os.Args[1:] { //range產出一對值
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}