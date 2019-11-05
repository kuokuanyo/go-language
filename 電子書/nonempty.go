package main

import "fmt"

//回傳非空字串
//底層陣列在呼叫時被修改
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			//將入元素
			strings[i] = s
			i++
		}
	}
	return strings[: i]
}

func main() {
	data := []string{"one","","three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)
}