package main

import(
	"geometry"
	"fmt"
)


func main() {
	perim := geometry.Path {{1, 1}, {5, 1}, {5, 4}, {1, 1}, } //一定要逗號
	//fmt.Println(geometry.Distance(perim))
	fmt.Println(perim.Distance())
}
