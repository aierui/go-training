package main

import "fmt"

type Person struct {
	age int
}

func (p Person) howOld () int {
	return p.age
}

func (p *Person) growUp() {
	p.age ++
}

/**

在调用方法的时候，
值类型既可以调用值接收者的方法，也可以调用指针接收者的方法；
指针类型既可以调用指针接收者的方法，也可以调用值接收者的方法。

也就是说，不管方法的接收者是什么类型，该类型的值和指针都可以调用，不必严格符合接收者的类型。

示例如下：
 */
func main() {
	A := Person{age: 18}
	fmt.Println(A.howOld())
	A.growUp()
	fmt.Println(A.howOld())

	B := &Person{age: 25}

	fmt.Println(B.howOld())
	B.growUp()
	fmt.Println(B.howOld())
}
