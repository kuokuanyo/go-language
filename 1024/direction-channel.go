//單向channel
package main

import "fmt"

func main() {
	//channel
	c := make(chan int)
	//receive(單向)
	cr := make(<-chan int)
	//send(單向)
	cs := make(chan<- int)

	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

	go foo(c)
	bar(c)

	fmt.Println("about to exit.")
}

//send function
func foo(c chan<- int) {
	c <- 42
}

//receive function
func bar(c <-chan int) {
	fmt.Println(<-c)
}
