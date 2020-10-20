package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s1 := map[string]interface{}{
		"name": "aierui",
		"age":  25,
	}

	for k, v := range s1 {
		fmt.Printf("key:%v value:%v type of value:%T\n", k, v, v) //此处age是int
	}
	b, _ := json.Marshal(s1)

	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)

	for k, v := range m {
		fmt.Printf("key:%v value:%v type of value:%T\n", k, v, v) //此处的age是float64
	}

	/*
		https://draveness.me/golang/docs/part4-advanced/ch09-stdlib/golang-json/
		key:name value:aierui type of value:string
		key:age value:25 type of value:int
		key:name value:aierui type of value:string
		key:age value:25 type of value:float64
	*/
}
