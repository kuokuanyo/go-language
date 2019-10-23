package main

import (
	"fmt"
)

type person struct {
	name string
}

func main() {
	p1 := person{
		name: "James Bond", //struct最後要逗號
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *person) { //參數為指標
	p.name = "Miss Moneypenny"
	// (*p).name = "Miss Moneypenny" 也可以
}
