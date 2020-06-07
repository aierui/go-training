package main

import "fmt"

func main() {
	ch1 := make(chan int, 2)
	// 发送方。
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Sender: sending element %v...\n", i)
			ch1 <- i
		}
		fmt.Println("Sender: close the channel...")
		close(ch1)
	}()

	// 接收方。
	// Loop
	for {
		elem, ok := <-ch1
		if !ok {
			fmt.Println("Receiver: closed channel")
			break
		}
		fmt.Printf("Receiver: received an element: %v\n", elem)
	}

	fmt.Println("End.")

	/*
		Sender: sending element 0...
		Sender: sending element 1...
		Sender: sending element 2...
		Sender: sending element 3...
		Receiver: received an element: 0
		Receiver: received an element: 1
		Receiver: received an element: 2
		Receiver: received an element: 3
		Sender: sending element 4...
		Sender: sending element 5...
		Sender: sending element 6...
		Sender: sending element 7...
		Receiver: received an element: 4
		Receiver: received an element: 5
		Receiver: received an element: 6
		Receiver: received an element: 7
		Sender: sending element 8...
		Sender: sending element 9...
		Receiver: received an element: 8
		Sender: close the channel...
		Receiver: received an element: 9
		Receiver: closed channel
		End.
	*/
}
