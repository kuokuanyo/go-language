package main 

import "fmt"

type Point struct {X, Y int}

func main() {
	p := Point{1, 2}
	q := Point{2, 1}
	fmt.Println(p.X == q.X && p.Y == q.Y)
	fmt.Println(p.X == q.Y && p.Y == q.X)
}