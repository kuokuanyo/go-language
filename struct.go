package main

import "fmt"

type Point struct { X, Y int}

type Employee struct {
	ID int
	Name, Address string
	Dob time.Time
	Position string
	Salary int
	ManagerID int
}

func main() {
	fmt.Println(Scale(Point{1, 2}, 5))
}

func Scale(p Point, factor int) Point {
	return Point {p.X * factor, p.Y * factor}
}

//為了效率，較大的struct通常會使用指標傳入函式或從函式回傳
func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}