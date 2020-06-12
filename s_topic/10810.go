package main

import "fmt"

//需求 计算1到20000的2倍
//就是用一个go程inputer来产生数据，写入ch中，然后用多个go程从ch中读走数据，然后乘以2操作
//当某个go程把ch最后一个数据读走后，将quit通道写入数据
//主go程读走quit通道数据后，结束程序

func inputer(ch chan<- int) {
	for i := 1; i <= 20000; i++ {
		ch <- i
	}
	close(ch)
}

func counter(ch <-chan int, quit chan<- bool) {
	num, ok := <-ch
	if !ok {
		quit <- true
	}
	fmt.Println(num * 2)
}

func main() {
	ch := make(chan int, 100)
	quit := make(chan bool)
	go inputer(ch)

	for i := 0; i < 500; i++ {
		go counter(ch, quit)
	}
	<-quit
}
