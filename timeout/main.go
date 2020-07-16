package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	asyncCall1()
	fmt.Println()
	asyncCall2()
	fmt.Println()
	asyncCall3()
}

func asyncCall1() {
	printTime("asyncCall1 starting")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*800))
	defer cancel()
	go func(ctx context.Context) {
		// 发送HTTP请求
	}(ctx)

	select {
	case <-ctx.Done():
		printTime("asyncCall1 finished")
		return
	case <-time.After(time.Duration(time.Millisecond * 805)):
		printTime("asyncCall1 timeout")
		return
	}
}

func asyncCall2() {
	printTime("asyncCall2 starting")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*800))
	defer cancel()
	timer := time.NewTimer(time.Duration(time.Millisecond * 900))

	go func(ctx context.Context) {
		// 发送HTTP请求
	}(ctx)

	select {
	case <-ctx.Done():
		timer.Stop()
		timer.Reset(time.Second)
		printTime("asyncCall2 finished")
		return
	case <-timer.C:
		printTime("asyncCall2 timeout")
		return
	}
}

func asyncCall3() {
	printTime("asyncCall3 starting")
	ctx := context.Background()
	done := make(chan struct{}, 1)

	go func(ctx context.Context) {
		// 发送HTTP请求
		done <- struct{}{}
	}(ctx)

	select {
	case <-done:
		printTime("asyncCall3 finished")
		return
	case <-time.After(time.Duration(800 * time.Millisecond)):
		printTime("asyncCall3 timeout")
		return
	}
}

func printTime(flag string) {
	fmt.Printf("flag : %s, time: %s \n", flag, time.Now().UnixNano())
}
