package main

import "fmt"

func main() {
	var any interface{}
	any = 1
	fmt.Println(any)

	any = "aierui"
	fmt.Println(any)

	any = 5.21
	fmt.Println(any)

	any = true
	fmt.Println(any)

	/*
		1
		aierui
		5.21
		true
	*/
	// 将值保存到空接口
	// 从空接口获取值
	// 空接口的值比较

}
