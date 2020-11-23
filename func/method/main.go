package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string
}

func (a Person) Name() string {
	a.name = "hi " + a.name

	return a.name
}

func NameOfP(a Person) string {
	a.name = "hi " + a.name

	return a.name
}

func main() {

	m1 := Person{
		name: "aierui",
	}

	fmt.Println(m1.Name())
	// equal to
	fmt.Println(Person.Name(m1)) // 接受者就是函数默认的第一个参数
	/*
		hi aierui
		hi aierui
	*/

	t1 := reflect.TypeOf(Person.Name)
	t2 := reflect.TypeOf(NameOfP)
	fmt.Println(t1 == t2) // true

}
