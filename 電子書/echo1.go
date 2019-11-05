//輸出命令列參數
package main //每個原始檔開頭都由package宣告

import( //要匯入的其它套件
	"fmt"
	"os"
)

func main() {
	var s, sep string //宣告兩個型別為string的變數(s,sep)，如果無初始化，值會是0或""
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "		
	}
	fmt.Println(s)
}