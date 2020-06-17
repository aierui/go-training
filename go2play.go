package main

import (
	"fmt"
)

func Print(type T)(s []T) {
	for _, v := range s {
		fmt.Print(v)
	}
}

func main() {
	Print([]string{"Hello, ", "go-training\n"})
}
