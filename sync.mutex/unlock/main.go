package main

import (
	"fmt"
	"sync"
)

var x int

func main() {
	var (
		w sync.WaitGroup
		m sync.Mutex
	)

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &m)
	}

	w.Wait()
	fmt.Println("final value of x", x)
}

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()

	wg.Done()

}

/*
 Unlock 未加锁 或者 已解锁的 Mutex 会？ 前者 panic
 sync.mutex 是否可重入？ 不会死锁，会 hang 住
 到底怎样解锁 defer ? 业务上(复杂，逻辑多，不清楚是否会有 panic ) 推荐 defer，如果 用 defer 那么 锁的粒度不一样，defer 将是整个 func 的，可以最小单元化
 m.Lock()
 defer m.Unlock


*/
