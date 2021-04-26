package main

import (
	"os"
	"runtime/trace"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	ch := make(chan string)
	go func() {
		ch <- "like"
	}()

	<-ch

	//step1. go run main.go 2> trace.out
	//step2. go tool trace trace.out
}
