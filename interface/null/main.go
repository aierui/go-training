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

	var n person
	rtn, err := json.Marshal(n)
	fmt.Println(err)
	fmt.Printf("rtn Type:%T, Value:%v\n", rtn, string(rtn))
	/*
		<nil>
		rtn Type:[]uint8, Value:{}
	*/

	var z person
	var zI interface{}
	zI = z
	zIRtn, err := json.Marshal(zI)
	fmt.Println(err)
	fmt.Printf("zIRtn Type:%T, Value:%v\n", zIRtn, string(zIRtn))

	nilRtn, err := json.Marshal(nil)
	fmt.Println(err)
	fmt.Printf("nilRtn Type:%T, Value:%v\n", nilRtn, string(nilRtn))

	/*
		<nil>
		zIRtn Type:[]uint8, Value:{}
		<nil>
		nilRtn Type:[]uint8, Value:null
	*/

}

func setVal(val interface{}) {

}

func assignVal(val interface{}) error {

	decoder := json.NewDecoder(bytes.NewReader([]byte(`{"test": {"unique": 1}}`)))
	decoder.UseNumber()

	return decoder.Decode(val)
}
