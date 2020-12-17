package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type person struct {
	conf map[string]item
}

type item struct {
	unique int `json:"unique"`
}

func main() {
	//var m *person
	m := &person{}
	fmt.Printf("m Type:%T, Value:%v\n", m, m)

	fmt.Printf("m.conf Type:%T, Value:%v\n", m.conf, m.conf)

	_ = assignVal(&m.conf)
	fmt.Printf("m.conf Type:%T, Value:%v\n", m.conf, m.conf)

}

func assignVal(val interface{}) error {

	decoder := json.NewDecoder(bytes.NewReader([]byte(`{"test": {"unique": 1}}`)))
	decoder.UseNumber()

	return decoder.Decode(val)
}
