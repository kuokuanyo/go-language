/*排序需要有三個條件:序列長度、比較元素的方法、交換兩元素的方式
因此sort.Interface有三個方法如下:
package sort

type Interface interface {
	Len() int
	Less(i, j int) bool  //i、j是元素的索引
	Swap(i, j int)
}
*/
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

//依照Artist欄位排序
//定義具有Len、Less、Swap方法的slice型別
type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//依年份排序
type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {
	//呼叫排序程式
	a := byArtist(tracks) //sort.Interface
	sort.Sort(a)
	printTracks(a)

	//反向排序
	sort.Sort(sort.Reverse(a))
	printTracks(a)

	//依年份
	b := byYear(tracks) //sort.Interface
	sort.Sort(b)
	printTracks(b)
}

/*反向排序利用到下列套件
package sort
type reverse struct { Interface } //上面的sort.Interface(未匯出)
func (r reverse) Less(i, j int) bool { return r.Interface.Less(j, i)}
func Reverse(data Interface) Interface {return reverse{data} }
*/
