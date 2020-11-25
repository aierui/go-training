package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	/*
		➜  declare git:(master) ✗ go vet main.go
		# command-line-arguments
		./main.go:8:16: loop variable i captured by func literal
	*/
	// 解法1
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	// 解法2
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			fmt.Println(i)
		}()
	}
}
