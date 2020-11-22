package _map

import (
	"fmt"
	"testing"
)

func TestMapCount(t *testing.T) {
	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}

	counter := 0
	for range m {
		if counter == 0 {
			delete(m, "A")
		}
		counter++
	}
	fmt.Println(counter)
}
