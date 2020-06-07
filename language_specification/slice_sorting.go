package main

import (
	"fmt"
	"sort"
)

type S struct {
	v int
}

func main() {
	s := []S{{1}, {3}, {5}, {2}}
	// A
	fmt.Printf("%#v", s)
	fmt.Println()
	sort.Slice(s, func(i, j int) bool { return s[i].v < s[j].v })
	fmt.Printf("%#v", s)
}
