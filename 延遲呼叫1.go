//defer也可以用於複雜函式除錯上的進入與離開動作
//defer延遲函式在return陳述更新函式的姊果變數之後執行
package main

import(
	"time"
	"log"
)

func main() {
	bigSlowOperation()
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")() //一定要有括號
	//模擬長時間操作
	time.Sleep(10 * time.Second)
}

//雙層函式
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}