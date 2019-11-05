package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "and I am", p.age, "years")
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32, //struct最後一定要記得加，
	}

	p1.speak()
}
