package main

// go tool compile -S -N -l main.go > add.txt
func main() {
	a, b := 1, 2
	add(a, b)
}

func add(a, b int) int {

	return a + b
}
