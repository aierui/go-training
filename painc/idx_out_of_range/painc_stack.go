package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter function main.")
	caller1()
	fmt.Println("Exit function main.")

	/**
	Enter function main.
	Enter function caller1.
	Enter function caller2.
	panic: runtime error: index out of range [5] with length 5

	goroutine 1 [running]:
	main.caller2()
	        /Users/develop/aierui/go-training/painc/q1/painc_stack.go:22 +0x82
	main.caller1()
	        /Users/develop/aierui/go-training/painc/q1/painc_stack.go:15 +0x7e
	main.main()
	        /Users/develop/aierui/go-training/painc/q1/painc_stack.go:9 +0x7e
	*/
}

func caller1() {
	fmt.Println("Enter function caller1.")
	caller2()
	fmt.Println("Exit function caller1.")
}

func caller2() {
	fmt.Println("Enter function caller2.")
	s1 := []int{0, 1, 2, 3, 4}
	e5 := s1[5]
	_ = e5
	fmt.Println("Exit function caller2.")
}
