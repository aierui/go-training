package main

func main() {
	callAdd()
}

func add(a, b int32) (int32, bool) {
	return a + b, true
}

func callAdd() int32 {
	a, _ := add(10, 20)
	return a
}
