package main

import "fmt"

type myInt1 = int
type myInt2 = int

type myInt3 int
type myInt4 int

func main() {

	var I1 myInt1
	var I2 myInt2
	var I3 myInt3
	var I4 myInt4

	I1 = 1
	I2 = 2
	I3 = 3
	I4 = 4

	fmt.Printf("Type:%T->Value:%v\n", I1, I1)
	fmt.Printf("Type:%T->Value:%v\n", I2, I2)
	fmt.Printf("Type:%T->Value:%v\n", I3, I3)
	fmt.Printf("Type:%T->Value:%v\n", I4, I4)

	/*
		Type:int->Value:1
		Type:int->Value:2
		Type:main.myInt3->Value:3
		Type:main.myInt4->Value:4
	*/

}
