package main


// In line ABCD, which ones of them have syntax issues?

type S struct {
}

func f(x interface{}) {
}

func g(x *interface{}) {
}

func m(x *S) {
}

func n(x S) {
}

//https://www.flysnow.org/2018/02/24/golang-function-parameters-passed-by-value.html

// 传值（值传递）、传引用（引用传递）

// 引用类型 和 传引用 是两个概念

// 在 golang 中 函数的参数只有传值，没有引用传递
// Go 语言选择了传值的方式，无论是传递基本类型、结构体还是指针，都会对传递的参数进行拷贝
// 将指针作为参数传入某一个函数时，在函数内部会对指针进行复制，也就是会同时出现两个指针指向原有的内存空间，所以 Go 语言中『传指针』也是传值。
// 当我们对 Go 语言中大多数常见的数据结构进行验证之后，其实就能够推测出 Go 语言在传递参数时其实使用的就是传值的方式，接收方收到参数时会对这些参数进行复制；
// 了解到这一点之后，在传递数组或者内存占用非常大的结构体时，我们在一些函数中应该尽量使用指针作为参数类型来避免发生大量数据的拷贝而影响性能。

func main() {
	s := S{}
	p := &s
	f(s) //A correct
	n(s) //A correct

	g(s) //B incorrect
	f(p) //C correct
	g(p) //D incorrect

	m(s) //
	m(p) //
}