package main

import "fmt"

func main() {
	var a, b, c, d int
	b = incr(a)
	fmt.Println(a, b) // 0 1

	d = incr2(c)
	fmt.Println(c, d) // 0 3
}

// 匿名返回值（局部变量 b）
func incr(a int) int {
	var b int

	defer func() {
		a++
		b++
	}()

	a++
	b = a

	return b
}

// 命名返回值
func incr2(a int) (b int) {
	defer func() {
		a += 5
		b += 2
	}()

	a++

	// equal to
	// b = a
	// return b
	return a
}
