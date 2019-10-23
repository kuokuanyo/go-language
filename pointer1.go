//pointer
package main

import "fmt"

func main() {
	x := 0
	fmt.Println("x befor", &x) //回傳位址
	fmt.Println("x befor", x)  //value
	foo(&x)
	fmt.Println("x after", &x) //address
	fmt.Println("x after", x)  //value
}

func foo(y *int) { //y為位址（指標）
	fmt.Println("y befor", y)  //address
	fmt.Println("y befor", *y) //value
	*y = 43
	fmt.Println("y after", y)  //address
	fmt.Println("y after", *y) //value

}
