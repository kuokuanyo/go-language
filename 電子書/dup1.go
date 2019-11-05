//輸出標準輸入中出現超過一次的行，前面加上次數
//使用了if、map、bufio
package main

import(
	//bufio套件幫助輸出輸入的效率與方便性
	//最重要的功能之一為讀取輸入並將其猜分成多行的Scanner型別
	"bufio"
	"fmt"
	"os"
)

func main() {
	//map保存一組鍵/值並提供儲存讀取檢測元素操作
	//鍵可以是任何 == 比較值的型別，通常是string，值可以是任何型別
	//內建的make建構出新的空map
	counts := make(map[string]int) 
	input := bufio.NewScanner(os.Stdin)
	//scan()在有下一行的時候會傳true，而沒有更多輸入時回傳false
	for input.Scan(){
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			//Printf產生格式化輸出
			//函式名稱以f(如Printf)結束，通常都是格式化函式
			//函式名稱以ln(如Println)結束，則是一般函式
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}