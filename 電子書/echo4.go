//每次取用變數的位址或複製一個指標時會建構出新的別名
package main

import (
	"flag"
	"fmt"
	"strings"
)
//套件階級變數
//flag.Bool建構出bool型別的flag變數，三個參數分別為flag名稱、變數預設值(false)與使用者提出無效的參數訊息
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	//flag.Parse()必須在使用flag之前呼叫，目的將flag變數預設值更新
	flag.Parse()
	//非flag參數可從flag.Args()取得字串的slice
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}