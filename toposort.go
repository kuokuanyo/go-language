//電腦課程與其必修課程
//排序出修課順序
//prereqs的值為string，鍵為slice
package main

import ("fmt"
		"sort"
)

var prereqs = map[string][]string {
	"algorithms": {"data structures"},
	"calculus"  : {"linear algebra"},

	"compilers" : {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures"      : {"discrete math"},
	"databases"            : {"data structures"},
	"discrete math"        : {"intro to programming"},
	"formal languages"     : {"discrete math"},
	"networks"             : {"operating systems"},
	"operating systems"    : {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	//toposort為slice，因此i為數字，course為迴圈參數
	for i, course := range toposort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

//回傳slice
func toposort(m map[string][]string) []string {
	//slice
	var order []string

	//map
	seen := make(map[string]bool)

	//slice
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	//排序
	sort.Strings(keys)

	/*function(不具名)
	當不具名函式需要遞迴時，必須先宣告一個變數，再指派不具名函式給該變數
	假設直接合併(visitAll := func(items []string) {visitAll(m[item])}，
	會出現visitAll位定義的錯誤) */
	var visitAll func(items []string)

	visitAll = func(items []string) {
		//因不用用到slice索引，因此使用_
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	visitAll(keys)
	return order
}