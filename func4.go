package main

import (
	"fmt"
	"math"
)

//圓形
type circle struct {
	radius float64
}

//方形
type square struct {
	length float64
}

//method
func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

//method
func (s square) area() float64 {
	return s.length * s.length
}

//介面
//只要有area()方法就符合介面
type shape interface {
	area() float64
}

func info(st string, s shape) {
	fmt.Println(st, s.area())
}

func main() {
	circ := circle{
		radius: 12.345,
	}
	squa := square{15}

	info("circle", circ)
	info("square", squa)
}
