//runtime套件可使用相同機制堆疊
package main

import(
	"runtime"
	"os"
	"fmt"
)

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

//多個函式defer時，會反序的執行
//發生panic時，所有延遲以反序的方式執行
func f(x int) {
	fmt.Printf("f(%d)\n", x +0 / x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func main() {
	defer printStack() //延遲呼叫
	f(3)
}