package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}
	people := []person{
		p1,
		p2,
	}

	fmt.Println(people)

	//func Marshal(v interface{}) ([]byte, error)
	bs, err := json.Marshal(people)

	//檢查錯誤
	if err != nil { //錯誤時
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
