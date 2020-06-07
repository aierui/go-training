package main

import (
	"fmt"
	"time"
)

func main() {
	//ch := make(chan int, 1)
	//
	//ch <- 1
	//
	//select {
	//case i := <-ch:
	//	fmt.Println("The ch crruent case %v %v", ch, i)
	//
	//default:
	//	fmt.Println("The select default case  %v %v", ch)
	//}

	m := time.Tick(1 * time.Second)

	fmt.Printf("The select default case %v type %T", m, m)
	fmt.Println()

	ch := make(chan int)
	go func() {
		for range m {
			ch <- 0
		}
	}()

	// select {} // 空的 select 语句会直接阻塞当前的 Goroutine，导致 Goroutine 进入无法被唤醒的永久休眠状态。

	fmt.Printf("The ch value %v type %T", ch, ch)
	fmt.Println()

	//select {
	//case ch <- 1:
	//
	//}

	for {
		select {
		case k := <-ch:
			fmt.Printf("case1  kk ==>> %v type => %T \n", k, k)
		case m := <-ch:
			fmt.Printf("case2  mm ==>> %v type => %T \n", m, m)
		}
	}

}
