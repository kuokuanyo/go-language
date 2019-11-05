//介面排序的程式碼只要元素型別改變，就必須宣告新的sort.Interface
//只有Less函式會改變，而Len與Swap在不同型別上有相同的定義
package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

//struct
//大寫才會匯出
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

//slice
var tracks []*Track = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")}, //要逗號
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s) //有問題
	}
	return d
}

//struct
//sort.Interface的具體型別不一定是slice
type customSort struct {
	t    []*Track               //slice
	less func(x, y *Track) bool //需要改變less函式
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

//printTracks以表格輸出撥放清單
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)          //指標
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length") //格式化
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------") //格式化
	for _, t := range tracks {                                            //不會用到索引位置，因此 _
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() //計算欄寬並輸出表格
}

func main() {
	//定義多層排序函式
	a := customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title //bool
		}
		if x.Year != y.Year {
			return x.Year < y.Year //bool
		}
		if x.Length != y.Length {
			return x.Length < y.Length //bool
		}
		return false
	}}
	sort.Sort(a)
	printTracks(a.t)
}
