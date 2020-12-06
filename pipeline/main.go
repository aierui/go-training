package main

import (
	"fmt"
	"sync"
)

// go tool compile -S -N -l main.go
func main() {
	// 组装pipeline
	//c := gen(2, 3)
	//
	//out := sq(c)
	//
	//// 消费输出
	//fmt.Println(<-out) // 4
	//fmt.Println(<-out) // 9
	//
	//fmt.Println()

	done := make(chan struct{})
	defer close(done)

	in := gen1(done, 2, 3)
	// Distribute the sq work across two goroutines that both read from in
	// 多个函数可以从相同的 channel 中读取，直到 channel 关闭，这称为：fan-out
	//c1 := sq1(done, in)
	//c2 := sq1(done, in)

	sqWorkerNum := 5
	workers := make([]<-chan int, sqWorkerNum)
	for i := 0; i < sqWorkerNum; i++ {
		ch := sq1(done, in)
		workers[i] = ch
	}

	fmt.Printf("%#v\n", workers)

	//time.Sleep(1 * time.Second)

	// Consume the merge output from c1 and c2.
	// 一个函数可以从多个输入读取，然后处理直到所有的都关闭，然后关闭单个 channel。这称为 fan-in
	out1 := merge(done, workers...)
	for n := range out1 {
		fmt.Printf("merge result:%v\n", n)
	}
	fmt.Println("Finished!")
	//done <- struct{}{}
	//done <- struct{}{}
}

//func gen(nums ...int) <-chan int {
//	out := make(chan int, 0)
//
//	go func() {
//		for _, num := range nums {
//			out <- num
//		}
//	}()
//
//	return out
//}

func gen1(done chan struct{}, nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, num := range nums {
			out <- num
			select {
			case out <- num:
			case <-done:
				fmt.Printf("gen1 finished\n")
				return
			}
		}
	}()

	return out
}

//func sq(in <-chan int) <-chan int {
//	out := make(chan int)
//	go func() {
//		for n := range in {
//			out <- n * n
//		}
//		close(out)
//	}()
//	return out
//}

func sq1(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		// When you range over a channel, iterations only produce a single value,
		// the values that were sent on the channel.
		// There is no index or key value like in case of slices or maps.
		for n := range in {
			fmt.Printf("n:%v, T:%T\n", n, n)
			out <- n * n
			select {
			case out <- n * n:
			case <-done:
				fmt.Printf("sq1 finished\n")
				return
			}
		}
	}()
	return out
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	out := make(chan int, 0)

	// Start an output goroutine for each input channel in cs. output
	// copies values from c to out until c is closed, then calls wg.Done
	output := func(c <-chan int) {
		defer wg.Done()
		for v := range c {
			select {
			case out <- v:
				fmt.Printf("inbound put outbound, value:%v\n", v)
			case <-done:
				fmt.Printf("merge finished\n")
				return
			}
		}
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are done.
	// This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
