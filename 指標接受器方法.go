//指標接受器方法
package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

//方法(指標)
func (p *Point) ScaleBy(factor float64) {
	//兩種表達方式一樣
	p.X = p.X * factor
	p.Y *= factor
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func main() {
	//回傳 *Point型別
	//可以這樣表示:var r *Point = &Point{1, 2}
	//var r = &Point{1, 2}
	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)

	//第二種方式
	p := Point{10, 100} //一般宣告
	(&p).ScaleBy(50) //放入指標到函式方法
	//也可以寫成(p).ScaleBy(50)，編譯器自動轉成指標
	fmt.Println(p)

	q := Point{5, 6}
	pptr := &q //指標

	//可以對Distance使用*Point接受器
	fmt.Println(pptr.Distance(p))
	//也可以
	fmt.Println((*pptr).Distance(p))
}

