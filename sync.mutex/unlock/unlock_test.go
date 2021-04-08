package main

import (
	"fmt"
	"sync"
	"testing"
)

func BenchmarkUnLock(b *testing.B) {

	var (
		w sync.WaitGroup
		m sync.Mutex
	)

	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000; i++ {
			w.Add(1)
			go increment(&w, &m)
		}

		w.Wait()
	}

	fmt.Println("final value of x", x)
}
