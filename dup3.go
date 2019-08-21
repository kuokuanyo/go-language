//輸出檔案中出現超過一次的行文字與次數
package main

import {
	"fmt"
	"io/ioutil"
	"os"
	"strings"
}

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		//ReadFile()回傳位元組的slice，必須轉換成string才能用split分開
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n",n, line)
		}
	}
}