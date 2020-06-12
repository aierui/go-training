package main

import (
	"fmt"
	"github.com/panjf2000/ants"
	"runtime"
	"sync"
	"time"
)

var (
	aChan = make(chan int, runtime.NumCPU())
	bChan = make(chan int, runtime.NumCPU())
	quit  = make(chan struct{})
	_     = startProcess()
	wg    sync.WaitGroup
)

func startProcess() struct{} {
	go func() {
		fmt.Println("process goroutine stared")
		defer func() {
			close(quit)
			close(aChan)
			fmt.Println("process goroutine stopped")
		}()
		pool, err := ants.NewPool(runtime.NumCPU())
		otherPool, err := ants.NewPool(runtime.NumCPU())
		if err != nil {
			fmt.Printf("can not init pool because of %#v\n", err)
			return
		}
		ticker := time.NewTicker(5 * time.Second)
		defer func() {
			fmt.Println("release pool")
			_ = pool.Release()
			_ = otherPool.Release()
			ticker.Stop()
		}()
		for {
			select {
			case a, ok := <-aChan:
				if !ok {
					return
				}
				aTmp := a // aTmp not affected by a
				//go func() {
				_ = pool.Submit(func() {
					fmt.Printf("the result is %d\n", aTmp)
					// todo with time
					time.Sleep(4 * time.Second)
					bChan <- aTmp + 1000
				})
				//}()
			case b, ok := <-bChan:
				if !ok {
					return
				}
				bTmp := b
				//go func() { // 非协程 导致阻塞
				//time.Sleep(10 * time.Second)
				_ = otherPool.Submit(func() {
					//_ = pool.Submit(func() { // 同otherPool 非协程导致阻塞
					fmt.Printf("======the result is %d\n", bTmp)
					// todo with time
					time.Sleep(2 * time.Second)
					wg.Done()
				})
				//}()
			case <-ticker.C:
				fmt.Printf("current pool size: %d/%d\n", pool.Running(), pool.Cap())
			case <-quit:
				fmt.Println("receive quit request")
				return
			}
		}
	}()
	return struct{}{}
}

// Stop will cause goroutine processing stopped
func Stop() {
	quit <- struct{}{}
	<-quit
}

func main() {
	i := 0
	for {
		if i > 100 {
			break
		}
		i += 1
		wg.Add(1)
		go func(num int) {
			//fmt.Println(num)
			aChan <- num
		}(i)
	}
	wg.Wait()
}
