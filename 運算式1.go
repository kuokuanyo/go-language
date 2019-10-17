package main

import(

)

//Point
type Point struct { X, Y float64}

//func
func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

//path is slice
type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point //空函數
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for i := range path {
		//呼叫path[i].Add(offset)或path[i].Sub(offset)
		path[i] = op(path[i], offset)
	}
}