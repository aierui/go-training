package main

import "fmt"

func main() {

	var i interface{}
	i = 1

	r, ok := i.(string)
	fmt.Println(r, ok) // r="",ok=false

	r1 := i.(string) // runtime.fatal panic interface conversion: interface {} is int, not string
	fmt.Println(r1)

}
