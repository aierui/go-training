package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Enter function main.")
	caller()
	fmt.Println("Exit function main.")

	/**
	Enter function main.
	Enter function caller.
	panic: something wrong

	goroutine 1 [running]:
	main.caller()
	        /Users/develop/aierui/go-training/painc/call/painc.go:16 +0xb5
	main.main()
	        /Users/develop/aierui/go-training/painc/call/painc.go:10 +0x7e
	*/
}

func caller() {
	fmt.Println("Enter function caller.")
	panic(errors.New("something wrong")) // 正例。
	panic(fmt.Println)                   // 反例。
	fmt.Println("Exit function caller.")
}
