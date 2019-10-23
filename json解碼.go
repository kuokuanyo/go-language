//解碼
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	//json格式
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	bs := []byte(s) //轉換slice
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	//people := []person{}
	var people []person

	//func Unmarshal(data []byte, v interface{}) error
	//解碼到people
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("all of the data", people)

	for i, v := range people {
		fmt.Println("---- PERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
