//定義一系列線段的Path型別授予Distance方法
package main

import(
	"fmt"
	"math"
)

type Point struct { x, y float64}

//Path代表一系列線段
//slice(Point型別)
type Path []Point

//回傳路徑長度
//方法宣告
//此Distance為Path.Distance
func (path Path) Distance() float64 {
	sum := 0.0
	//path為slice
	//i為第幾個位置
	for i := range path { 
		if i > 0 {
			sum += path[i - 1].Distance(path[i]) //此Distance為Point.Distance
		}
	}
	return sum
}

//同樣功能用方法宣告(Point型別)
//Point.Distance方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x - p.x, q.y - p.y)
}

func main() {
	perim := Path {
	{1, 1},
	{5, 1},
	{5, 4},
	{1, 1}, //一定要逗號
}
fmt.Println(perim.Distance())
}