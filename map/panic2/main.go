package main

import "fmt"

type Resp struct {
	ErrNo  int                    `json:"errno"`
	ErrMsg string                 `json:"errmsg"`
	Data   map[string]interface{} `json:"data"`
}

func main() {
	rtn1 := &Resp{}
	fmt.Printf("rtn1 => %#v", rtn1)
	// data 中的 map 还是为 nil
	//rtn1 => &main.Resp{ErrNo:0, ErrMsg:"", Data:map[string]interface {}(nil)}
}
