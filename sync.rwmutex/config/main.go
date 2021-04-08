package main

import (
	"fmt"
	"sync"
	"time"
)

var x int

func increment(m *sync.RWMutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
}

func readX(m *sync.RWMutex) {
	m.RLock()
	defer m.RUnlock()
	fmt.Printf("x is %d\n", x)
}

func main() {
	var (
		m sync.RWMutex
	)

	for i := 0; i < 10; i++ {
		go readX(&m)
	}

	increment(&m)

	time.Sleep(5 * time.Second)

	fmt.Println("final value of x", x)
}
