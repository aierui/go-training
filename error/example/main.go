package main

import "fmt"

func main() {
	var err error
	fmt.Printf("err:%v\n", err.Error()) // panic: runtime error: invalid memory address or nil pointer dereference
}
