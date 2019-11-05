//struct中有相似性或重複的欄位
//嵌入struct與不具名欄位
package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct{
	Point //不具名欄位
	Radius int
}

type Wheel struct {
	Circle //不具名欄位
	Spokes int
}

/*或用以下方是宣告
第一種:
w = Wheel{Circle{Point{8, 8}, 5}, 20}

第二種:
w = Wheel{
	Circle: Circle{
		Point: Point{X: 8, Y:8},
		Radius: 5, 這裡一定要逗點
	},
	Spokes: 20, 一定要逗點
}
*/

func main() {
	var w Wheel
	w.X = 8 //與w.Circle.Point.x = 8相同
	w.Y = 8 //與w.Circle.Point.y = 8相同
	w.Radius = 5 //同w.Circle.Radius = 5
	w.Spokes = 20
	//%v以自然格式顯示
	//#會顯示所有欄位名稱(完整資訊)
	fmt.Printf("%#v\n", w)
	w.X = 42
	fmt.Printf("%v\n", w)
}