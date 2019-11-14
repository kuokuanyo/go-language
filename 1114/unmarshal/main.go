package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type img struct {
	Width     int    `json:"Width"`
	Height    int    `json:"Height"`
	Title     string `json:"Title"`
	Thumbnail struct {
		URL    string `json:"Url"`
		Height int    `json:"Height"`
		Width  int    `json:"Width"`
	} `json:"Thumbnail"`
	Animated bool  `json:"Animated"`
	IDs      []int `json:"IDs"`
}

func main() {
	//variable
	var data img
	//json
	rcvd := `{"Width":800,"Height":600,"Title":"View from 15th Floor",
	"Thumbnail":{"Url":"http://www.example.com/image/481989943","Height":125,"Width":100},
	"Animated":false,"IDs":[116,943,234,38793]}`

	//decode(unmarshal)
	//Unmarshal encode data and store the result in pointer v
	//func Unmarshal(data []byte, v interface{}) error
	//第二個參數必須為指標，否則會出現錯誤
	err := json.Unmarshal([]byte(rcvd), &data)
	//if error
	if err != nil {
		log.Fatalln("error unmarshalling", err)
	}
	//print
	fmt.Println(data)

	for i, v := range data.IDs {
		fmt.Println(i, v)
	}
	fmt.Println(data.Thumbnail.URL)
}
