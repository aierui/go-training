package main


import "fmt"

type coder interface {
	code()
	debug()
}

type Gopher struct {
	language string
}

// *Gopher 类型也没有实现 code 方法，但是因为 Gopher 类型实现了 code 方法，所以让 *Gopher 类型自动拥有了 code 方法。
// 隐式实现
func (p Gopher) code() {
	fmt.Printf("I am coding %s language\n", p.language)
}

func (p *Gopher) debug() {
	fmt.Printf("I am debuging %s language\n", p.language)
}

/**
q5 index.go 中 前面说过，不管接收者类型是值类型还是指针类型，都可以通过值类型或指针类型调用，这里面实际上通过语法糖起作用的。


先说结论：
1、实现了接收者是值类型的方法，相当于自动实现了接收者是指针类型的方法；
2、而实现了接收者是指针类型的方法，不会自动生成对应接收者是值类型的方法。


接收者是指针类型的方法，很可能在方法中会对接收者的属性进行更改操作，
从而影响接收者；而对于接收者是值类型的方法，在方法中不会对接收者本身产生影响。

使用指针作为方法的接收者的理由：

1、方法能够修改接收者指向的值。
2、避免在每次调用方法时复制该值，在值的类型为大型结构体时，这样做会更加高效。

*/
func main() {
	//var c coder = &Gopher{"Go"}
	//c.code()
	//c.debug()

	//var c coder = &Gopher{"Go"}
	//c.code()

	//var c coder = &Gopher{"Go"}
	//c.debug()

	//var c = Gopher{"Go"}
	//c.code()

	//var c = Gopher{"Go"}
	//c.debug()


	var c coder = Gopher{"Go"} // ./index.go:41:6: cannot use Gopher literal (type Gopher) as type coder in assignment:
	c.code()
	c.debug()

}