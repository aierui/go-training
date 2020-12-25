package main

func main() {
	callAdd()
}

//go:noinline
func add(a, b int32) (int32, bool) {
	return a + b, true
}

//go:noinline
func callAdd() int32 {
	a, _ := add(10, 20)
	return a
}
