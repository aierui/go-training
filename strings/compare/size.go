package main

import (
	"fmt"
)

func main() {
	compare()
}

// env GOSSAFUNC=compare go build size.go
func compare() {
	fmt.Println(999 <= 1000)
	fmt.Println("903" <= "1000")
	fmt.Println("903" <= "1000")
	fmt.Println("abc" <= "abz")
	fmt.Println("abc" <= "abzd")
	fmt.Println("abc" <= "abcd")
}
