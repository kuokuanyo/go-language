package main

import (
	"math"
	"fmt"
)

type Point struct {
	X float64
	Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (p *Point) ScaleBy(factor float64) Point{
	//兩種表達方式一樣
	p.X = p.X * factor
	p.Y *= factor
	return *p
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	//Distance
	//第一種方式
	distanceFromP := p.Distance
	fmt.Println(distanceFromP(q))
	var origin Point //{0, 0}
	fmt.Println(distanceFromP(origin))
	fmt.Printf("%T\n", distanceFromP)

	//第二種方式
	distance := Point.Distance
	fmt.Println(distance(p, q))
	fmt.Printf("%T\n", distance)

	//ScaleBy
	//第一種方式
	scaleP := p.ScaleBy
	fmt.Println(scaleP(2)) //{2, 4}
	fmt.Println(scaleP(3)) //{6, 12}
	fmt.Println(scaleP(10)) //{60, 120}

	//第二種方式
	scale := (*Point).ScaleBy
	fmt.Println(scale(&p, 10))
}