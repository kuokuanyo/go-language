//以svg會3D圖
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320 //畫布尺寸
	cells = 100 //格數
	xyrange = 30.0 //軸範圍(-xyrange..+xyrange)
	xyscale = width / 2 / xyrange //x, y單位像素
	zscale = height * 0.4 //z單位像素
	angle = math.Pi / 6 //x與y角度(30)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30), cos(30)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' " + 
		"style='stroke: grey; fill: white; stroke-width: 0.7' " + 
		"width='%d' height='%d'>", width, height)
	for i :=0; i < cells; i++ {
		for j :=0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

//corner回傳兩個值(格的座標)
func corner(i, j int) (float64, float64) {
	//找出(i,j)格的點(x,y)
	x := xyrange * (float64(i) / cells - 0.5)
	y := xyrange * (float64(j) / cells - 0.5)

	//計算z高度
	z := f(x, y)

	//投射(x,y,z)到2D畫布(sx,sy)
	sx := width / 2 + (x - y) * cos30 * xyscale
	sy := height / 2 + (x + y) * sin30 * xyscale - z * zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) //與(0,0)的距離
	return math.Sin(r) / r
}