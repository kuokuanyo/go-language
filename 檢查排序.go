package main

import (
	"fmt"
	"sort"
)

func main() {
	values := []int{3, 1, 4, 1} //slice
	//檢查是否已經排序
	fmt.Println(sort.IntsAreSorted(values))        //false
	sort.Ints(values)                              //排序
	fmt.Println(values)                            //1 1 3 4
	fmt.Println(sort.IntsAreSorted(values))        //True
	sort.Sort(sort.Reverse(sort.IntSlice(values))) //反轉
	fmt.Println(values)                            //4 3 1 1
	fmt.Println(sort.IntsAreSorted(values))        //false
}
