//計算unicode字元次數
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	//建立一個map計算unicode字元次數
	//rune unicode
	counts := make(map[rune]int) 
	var utflen [utf8.UTFMax + 1]int //計算utf8編碼長度
	invalid :=0 //計算無效utf8字元

	//NewReader相當於NewReaderSize，將變數封裝成擁有size緩存的對象
	in := bufio.NewReader(os.Stdin)
	for {
		//ReadRune讀取utf8編碼的unicode值
		//回傳該碼值、編碼長度、可能的錯誤
		r, n, err := in.ReadRune()
		//錯誤
		if err == io.EOF {
			break
		}
		//出現問題
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		//回傳rune是unicode.ReplacementChar且長度為1
		//unicode.ReplacementChar代表無效的unicode點
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		//都沒問題時
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invaild > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}