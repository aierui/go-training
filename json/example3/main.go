package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

type ActionTypeMapping struct {
	Conf map[string]ItemMapping
}

type ItemMapping struct {
	Unique string `json:"unique"`
}

func main() {
	fmt.Println("Hello, playground")

	s := `{"ark":{"unique":"100"},"caller":{"unique":"1"}}`

	ret2 := ActionTypeMapping{}
	err2 := jsoniter.Unmarshal([]byte(s), &ret2.Conf)
	fmt.Printf("err %v\n", err2)
	fmt.Println(ret2)
}
