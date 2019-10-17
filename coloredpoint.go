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
	Point
	Color color.RGBA //struct
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

/*
等同於
func (p ColoredPoint) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}
*/

func (p *Point) ScaleBy(factor float64) {
	//兩種表達方式一樣
	p.X = p.X * factor
	p.Y *= factor
}

/*
等同於
func (p *ColoredPoint) ScaleBy(factor float64) {
	//兩種表達方式一樣
	p.X = p.X * factor
	p.Y *= factor
}
*/

func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}

	//ColoredPoint(p)會自動找到Point
	fmt.Println(p.Distance(q.Point))
	//與可以fmt.Println(p.Point.Distance(q.Point))

	/*q不會自動找到Point，因此不能直接呼叫q
	fmt.Println(p.Distance(q)) 錯誤
	*/

	//ColoredPoint(p, q)會自動找到指標給ScaleBy函式
	p.ScaleBy(2)
	//也可以直接指派Point
	q.Point.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))
}