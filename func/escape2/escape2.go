package main

import "fmt"

// go build -gcflags="-m"

func main() {
	fmt.Println("Called stackAnalysis", stackAnalysis())
}

//go:noinline
func stackAnalysis() *int {
	data := 55
	return &data
}

/*
# github.com/aierui/go-training/func/escape2
./escape2.go:8:13: inlining call to fmt.Println
./escape2.go:13:2: moved to heap: data
./escape2.go:8:14: "Called stackAnalysis" escapes to heap
./escape2.go:8:13: []interface {} literal does not escape
<autogenerated>:1: .this does not escape
*/