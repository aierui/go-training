package main

import (
	"fmt"
	"reflect"
)

type S struct {
	a, b, c string
}

// https://medium.com/golangspec/comparison-operators-in-go-910d9d788ec0

// https://flaviocopes.com/golang-comparing-values/

// @todo 这里可以画一张各个类型对比表，类似 https://www.php.net/manual/zh/types.comparisons.php


func main() {
	x := interface{}(&S{"a", "b", "c"})
	y := interface{}(&S{"a", "b", "c"})

	fmt.Printf("%T -> %v  pointer => %p \n", x, x, x)
	fmt.Printf("%T -> %v  pointer => %p \n", y, y, y)

	fmt.Println(x == y) //A why?

	fmt.Println(reflect.DeepEqual(x, y))

	var a int = 1
	var b rune = '1'
	fmt.Println(a == b)


	var a1 int = 1
	var b1 int = 2
	var c float32 = 3.3
	var d float32 = 4.4
	fmt.Println(a1 == b1) // false 1 !=2
	fmt.Println(c == d) // false
}
