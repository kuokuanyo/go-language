//自定義排序
package main

import (
	"fmt"
	"sort"
)

type person struct {
	First string
	Age   int
}

type ByAge []person

func main() {
	p1 := person{"James", 32}
	p2 := person{"Moneypenny", 27}
	p3 := person{"Q", 64}
	p4 := person{"M", 56}

	people := ByAge{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(people)
	fmt.Println(people)
}

//method
func (p person) String() string {
	return fmt.Sprintf("%s: %d", p.First, p.Age)
}

//func Sort(data Interface)

/*sort套件
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}
*/

//下列三個方法符合Interface
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
