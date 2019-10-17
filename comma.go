//對整數插入逗號
//ex:12345 -> 12,345
func comma(s string) string {
	n := len(s)
	//假設長度小於三，直接貼上
	if n <= 3 {
		return s
	}
	//大於三
	return comma(s[: n-3]) + "," + s[n-3:]
}