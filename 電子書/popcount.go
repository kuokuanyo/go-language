//套件初始化
//使用func init() { /* ... */}初始化
package popcount

//pc[i] 是i的人口計算
var pc [256]byte

//初始化
func init() {
	for i := range pc { //也可以寫成 for i, _ := range pc
		pc[i] = pc[i / 2] + byte(i & 1)
	}
}

//popcount回傳x的人口計算(值為1的位元素)
func PopCount(x uint64) int {
	return int(pc[byte(x >> ( 0 * 8))] +
		pc[byte(x >> (1 * 8))] +
		pc[byte(x >> (2 * 8))] +
		pc[byte(x >> (3 * 8))] + 
		pc[byte(x >> (4 * 8))] + 
		pc[byte(x >> (5 * 8))] +
		pc[byte(x >> (6 * 8))] +
		pc[byte(x >> (7 * 8))] )
}