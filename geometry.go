//OOP(物件導向)
//方法宣告是在一般宣告的函示名稱錢加上一個額外參數
package geometry

import(
    "math"
)

//struct
//要大寫才會輸出
type Point struct { X, Y float64}

//傳統函式
func Distance(p, q Point) float64 {
    return math.Hypot(q.X - p.X, q.Y - p.Y)
}

//同樣功能用方法宣告(Point型別)
//Point.Distance方法
func (p Point) Distance(q Point) float64 {
    return math.Hypot(q.X - p.X, q.Y - p.Y)
}

//Path代表一系列線段
//slice(Point型別)
type Path []Point

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