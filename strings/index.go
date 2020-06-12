package main

import (
	"fmt"
	"strings"
)

func main()  {

	s1 := "abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdef"
	s2 := "qweeqweeqweeqweeqweeqweeqweeqweeqweeqweeqweeqweeqweeqweeqweeqweeqweeqweeqweeqweeqweeqweeqabc"

	c1 := strings.Index(s1, s2)
	fmt.Printf("c1=%v \n", c1)

}
