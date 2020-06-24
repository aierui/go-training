package main

import (
	"fmt"
	"unsafe"
)

type M struct {
	i, j int
}

func main() {
	a := &M{1, 12}
	add(a)
	fmt.Printf("[%p] %v\n", a, a)
}

func add(m *M) {
	// todo
	ptr := unsafe.Pointer(m)

	for i := 0; i < 2; i++ {
		//c := (*int)(unsafe.Pointer((uintptr(ptr) + uintptr(8*i))))
		c := (*int)(unsafe.Pointer(uintptr(ptr) + uintptr(8*i)))

		*c += i + 1

		fmt.Printf("[%p] %d\n", c, *c)
	}
}
