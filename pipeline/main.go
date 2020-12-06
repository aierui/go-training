package main

import (
	"fmt"
	"sync"
)

func main() {
	// 组装pipeline
	c := gen(2, 3)

	out := sq(c)

	// 消费输出
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9

	fmt.Println()

	done := make(chan struct{}, 2)
	defer close(done)

	in := gen(2, 3)
	// Distribute the sq work across two goroutines that both read from in
	// 多个函数可以从相同的 channel 中读取，直到 channel 关闭，这称为：fan-out
	c1 := sq(in)
	c2 := sq(in)

	//time.Sleep(1 * time.Second)

	// Consume the merge output from c1 and c2.
	// 一个函数可以从多个输入读取，然后处理直到所有的都关闭，然后关闭单个 channel。这称为 fan-in
	for n := range merge(done, c1, c2) {
		fmt.Println(n) // 4 then 9, then 4
	}
}

func gen(nums ...int) <-chan int {
	out := make(chan int, 0)

	go func() {
		for _, num := range nums {
			out <- num
		}
	}()

	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		// When you range over a channel, iterations only produce a single value,
		// the values that were sent on the channel.
		// There is no index or key value like in case of slices or maps.
		for n := range in {
			fmt.Printf("n:%v, T:%T\n", n, n)
			out <- n * n
		}
		close(out)
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
				fmt.Printf("inbound put outbound")
			case <-done:
				fmt.Printf("Finished")
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
