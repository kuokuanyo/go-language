package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cities []struct {
	Precision string  `json:"precision"`
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
	Address   string  `json:"Address"`
	City      string  `json:"City"`
	State     string  `json:"State"`
	Zip       string  `json:"Zip"`
	Country   string  `json:"Country"`
}

func main() {
	var data cities
	rcvd := `[{"precision":"zip","Latitude":37.7668,"Longitude":-122.3959,
	"Address":"","City":"SAN FRANCISCO","State":"CA","Zip":"94107","Country":"US"},
	{"precision":"zip","Latitude":37.371991,"Longitude":-122.02602,"Address":"",
	"City":"SUNNYVALE","State":"CA","Zip":"94085","Country":"US"}]`
	//decode(unmarshal)
	//decode rcvd and store the result in data(pointer)
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(data)
}
