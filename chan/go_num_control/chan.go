package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 灵活控制 goroutine 并发数量
// 灵活 chan + sync

func main() {
	goNum := 4
	channelNum := 10
	ch := make(chan int, 5)
	var wg sync.WaitGroup

	for i := 0; i < goNum; i++ {
		wg.Add(1)

		// 生成指定数量的 goroutine
		// 每个 goroutine 接受 chan 数据
		go func() {
			defer wg.Done()

			for d := range ch {
				fmt.Printf("index => %d, time => %d \n", d, time.Now().Unix())
				//time.Sleep(time.Second * time.Duration(d))
			}
		}()
	}

	for i := 0; i < channelNum; i++ {
		ch <- i
		fmt.Printf("index => %d, goroutine Num => %d\n", i, runtime.NumGoroutine())
		//time.Sleep(time.Second)
	}

	close(ch)
	wg.Wait()
}
