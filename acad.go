package acad

func Sum(xi ...int) int {
	s := 0
	//xi is []int
	for _, v := range xi {
		s += v
	}
	return s
}
