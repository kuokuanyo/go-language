//OOP(物件導向)
//方法宣告是在一般宣告的函示名稱錢加上一個額外參數
package main

import(
	"math"
	"fmt"
)

//struct
type Point struct { x, y float64}

//傳統函式
func Distance(p, q Point) float64 {
	return math.Hypot(q.x - p.x, q.y - p.y)
}

//同樣功能用方法宣告(Point型別)
//Point.Distance方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x - p.x, q.y - p.y)
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))
}