package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type city struct {
	Bali       string  `json:"Postal"`
	Kauai      float64 `json:"Latitude"`
	Maui       float64 `json:"Longitude"`
	Java       string  `json:"Address"`
	Newzealand string  `json:"City"`
	Skye       string  `json:"State"`
	Oahu       string  `json:"Zip"`
	Hawaii     string  `json:"Country"`
}

type cities []city

func main() {
	var data cities

	rcvd := `[{"Postal":"zip","Latitude":37.7668,"Longitude":-122.3959,"Address":"",
	"City":"SAN FRANCISCO","State":"CA","Zip":"94107","Country":"US"},
	{"Postal":"zip","Latitude":37.371991,"Longitude":-122.02602,"Address":"",
	"City":"SUNNYVALE","State":"CA","Zip":"94085","Country":"US"}]`

	//decode(unmarshal)
	//decode rcvd and store the result in pointer to data
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(data)
	//use tag
	fmt.Println(data[1].Kauai)
}
