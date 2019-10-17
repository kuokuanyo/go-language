//輸出水的沸點
package main

import "fmt"

//常數boilingF是套件層級的宣告
//不只在宣告檔案可見，也貫穿整個套件下所有檔案
const boilingF = 212.0

func main() {
	//變數f與c是函式的區域宣告
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gF or %gC\n", f, c)
}fmt.