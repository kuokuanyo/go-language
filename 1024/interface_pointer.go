//interface
//pointer
package main

import "fmt"

func main() {
	p1 := person{first: "James"}
	/*
		saysomething(p1) 會出現錯誤，因為是person的指標才有speak()method
	*/
	saysomething(&p1) //參數給指標，正確
	p1.speak()        //method雖然參數要給定指標，但它會自動轉換成指標
}

type person struct {
	first string
}

//method
func (p *person) speak() { //參數必須要是指標
	fmt.Println("hello")
}

//interface
type human interface {
	speak()
}

func saysomething(h human) { //參數是只要包含speak() method即可
	h.speak()
}
