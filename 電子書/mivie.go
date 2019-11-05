package main

import (
	"log"
	"fmt"
	"encoding/json"
)

type Movie struct {
	Title string
	//json欄位第一個標籤都是指定go的替代名稱
	Year int `json:"released"`
	//color有設定額外選項，omitempty代表欄位如為空值則不輸出(0 or false)
	Color bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie {
		{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	//json.Marshal出現結果會非常緊湊難以閱讀
	//將go的資料結構轉換成json(稱為marshaling)
	//只有匯出的欄位會被marchal
	data, err := json.Marshal(movies)
	//出現問題
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	//第二種方法較整齊且縮排
	//因之前指派過，現在用=直接宣告
	data, err = json.MarshalIndent(movies, "", "   ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	//unmarshaling為反向操作，將json解碼並產生go資料結構
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)
}