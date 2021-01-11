package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Item struct {
	Status int
	Lv1    int
	Lv2    int
	Lv3    int
	STime  string
	ETime  string `json:"e_time"`
}

func main() {

	js := `{
		"lv1": 3000,
		"status": 1,
		"lv3": 4000,
		"lv2": -1,
		"stime": "2021-01-01 12:00:00",
		"e_time": "2021-12-01 12:00:00"
	  }`

	var itemValue Item

	decode := json.NewDecoder(bytes.NewBufferString(js))
	decode.UseNumber()
	err := decode.Decode(&itemValue)
	fmt.Println(err)

	fmt.Printf("%#v\n", itemValue)




}
