package main

import (
	"fmt"
	"log"
	"runtime"
)

/**
探索golang程序启动过程
https://cbsheng.github.io/posts/%E6%8E%A2%E7%B4%A2golang%E7%A8%8B%E5%BA%8F%E5%90%AF%E5%8A%A8%E8%BF%87%E7%A8%8B/

1. go build chan.go
2. gdb chan (lldb chan)

*/

/*
	runtime.args
	runtime.osinit
	runtime.schedinit
	runtime.newproc
	runtime.mstart
*/

func main() {
	fmt.Println("hello world!")
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	cpu()
}

func cpu() {
	maxCPU := runtime.NumCPU()
	NumGoroutine := runtime.NumGoroutine()
	log.Println("maxCPU ", maxCPU)
	log.Println("NumGoroutine ", NumGoroutine)

	var num = 100

	var sign = make(chan struct{}, num)

	for i := 0; i < num; i++ {
		go func(i int) {
			log.Println("i ==> ", i)
			sign <- struct{}{}

			NumGoroutine := runtime.NumGoroutine()
			log.Println("NumGoroutine ii==>  ", NumGoroutine, i)

		}(i)
	}

	for j := 0; j < num; j++ {
		<-sign
	}

	maxCPU = runtime.NumCPU()
	NumGoroutine = runtime.NumGoroutine()
	log.Println("maxCPU ", maxCPU)
	log.Println("NumGoroutine ", NumGoroutine)
}
