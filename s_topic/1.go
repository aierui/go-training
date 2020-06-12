package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(1000 * time.Millisecond)
			fmt.Printf("2")
		}

		fmt.Printf("1")
	}
}
