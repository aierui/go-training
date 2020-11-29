package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func main() {
	fmt.Println("start")

	// goroutine exception panic
	Go(func() {
		fmt.Println("go routine throw new panic")
		panic("Suppose there is a panic")
	})

	time.Sleep(5 * time.Second)
}

func Go(x func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				errData := make(map[string]interface{}, 0)
				errData["err"] = err
				errData["stack"] = string(debug.Stack())
				fmt.Printf("recover_panic %#v\n", errData)
			}
		}()
		x()
	}()
}
