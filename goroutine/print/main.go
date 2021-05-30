package main

import (
	"fmt"
	"runtime"
)

func main() {
	example1()
	//example2()
}

func example1() {
	go say("world")
	say("hello")
}

func example2() {
	go say("world")
	say2("hello")
}

func say(s string) {
	for i := 0; i < 2; i++ {
		fmt.Println(s)
	}
}

func say2(s string) {
	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}
