//輸出出現超過一次以上的行，讀取stdin或一系列檔案
package main

import(
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countlines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			//os.Open()回傳兩個值，第一個是開啟的檔案(*os.File)，用於後續的Scanner
			//第二個值是內建的error型別，如果err等於內建的特殊值nil，則檔案成功開啟
			f, err := os.Open(arg) //回傳兩個值，存給f, err
			if err != nil { //不等於nil，發生問題
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countlines(f, counts)
			f.close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
func countlines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}