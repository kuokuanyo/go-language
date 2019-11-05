package main

import (
	//"bytes"

	"fmt"
	"t"
)

type IntSet struct {
	t.IntSet
}

func main() {
	x := IntSet{}
	y := t.IntSet{}
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())

	fmt.Println(x.Has(9), x.Has(123))

	fmt.Println(&x)         //指標
	fmt.Println(x.String()) //呼叫String()，編譯器會插入&變成指標
	fmt.Println(x)          //沒有String方法，不適指標
}
