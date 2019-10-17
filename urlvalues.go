package main

import (
	"fmt"
)

//map
//key is string, value is slice
type Values map[string][]string

//回傳一個建的第一個關聯值，若無則回傳""
//方法宣告
func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0] //第一個值
	}
	return "" //無則回傳""
}

//加入值，加入該鍵現有關聯值
func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}

func main() {
	m := Values{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q")) //沒有q,回傳""
	fmt.Println(m.Get("item")) //回傳第一個值"1"
	fmt.Println(m["item"]) //["1", "2"]
	fmt.Println(m)

	//將m指派為空map
	m = nil
	fmt.Println(m.Get("item")) //""
	//m.Add("item", "1") //panic :對空map指派
	m = Values{"lang": {"en"}}
	fmt.Println(m)
}