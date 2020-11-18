package main

import "fmt"

// go build -gcflags="-m"

func main() {
	fmt.Println("Called stackAnalysis", stackAnalysis())
}

//go:noinline
func stackAnalysis() int {
	data := 55
	return data
}

/*
# github.com/aierui/go-training/func/escape
./escape.go:8:13: inlining call to fmt.Println
./escape.go:8:14: "Called stackAnalysis" escapes to heap
./escape.go:8:51: stackAnalysis() escapes to heap
./escape.go:8:13: []interface {} literal does not escape
<autogenerated>:1: .this does not escape
*/
