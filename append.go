package main

import "fmt"

//append函式
func appendInt(x []int, y int) []int {
	//z為slice(目前為空)
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) { //假設小於等於x的容量(代表還有空間)
		//擴張slice
		z = x[: zlen]
	} else {
		//容量不夠，分配新陣列
		//倍增(容量)
		zcap := zlen
		if zcap < 2 * len(x) {
			//將容量倍增
			zcap = 2 * len(x)
		}
		//make建立slice
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	//新增元素
	z[len(x)] = y
	return z
}

func main() {
	var x, y []int
	for i :=0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}