//不具名欄位的型別可以是具名型別的指標
package main

import (
	"image/color"
	"math"
	"fmt"
)

//欄位開頭必須要大寫才會匯出
type Point struct {
	X float64
	Y float64
}

type ColoredPoint struct {
	*Point //指標
	Color color.RGBA //struct
}

//Distance
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (p ColoredPoint) Distance(q Point) float64 {
	return p.Point.Distance(q)
}

//ScaleBy
func (p *Point) ScaleBy(factor float64) {
	//兩種表達方式一樣
	p.X = p.X * factor
	p.Y *= factor
}

func (p *ColoredPoint) ScaleBy(factor float64) {
	p.Point.ScaleBy(factor)
}

func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{&Point{1, 1}, red}
	var q = ColoredPoint{&Point{5, 4}, blue}

	//ColoredPoint(p)會自動找到Point
	//q呼叫Point，並且為指標，因此*q.Point
	fmt.Println(p.Distance(*q.Point))
	//指向同一個
	q.Point = p.Point
	p.ScaleBy(2)
	fmt.Println(*p.Point, *q.Point)
}