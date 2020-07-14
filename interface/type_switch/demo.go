package main

import "fmt"

func main() {

	// 类型转换
	// 语法：<结果类型> := <目标类型> ( <表达式> )
	// 类型转换是用来在不同但相互兼容的类型之间的相互转换的方式，所以，当类型不兼容的时候，是无法转换的。如下：
	var var1 int = 7

	var2 := float32(var1)
	var3 := int64(var1)

	fmt.Printf("%T->%v\n", var1, var1)
	fmt.Printf("%T->%v\n", var2, var2)
	fmt.Printf("%T->%v\n", var3, var3)

	var4 := new(int32)
	fmt.Printf("%T->%v\n", var4, var4)
	// src/puzzlers/interface/type_switch/chan.go:22:16: cannot convert var4 (type *int32) to type int32
	//var5 := *int32(var4)
	//fmt.Printf("%T->%v\n", var5, var5)

	// 类型断言

	// <目标类型的值>，<布尔参数> := <表达式>.( 目标类型 ) // 安全类型断言
	// <目标类型的值> := <表达式>.( 目标类型 )　　//非安全类型断言
	// 类型断言的本质，跟类型转换类似，都是类型之间进行转换，不同之处在于，类型断言实在接口之间进行，
	// 相当于Java中，对于一个对象，把一种接口的引用转换成另一种。
}
