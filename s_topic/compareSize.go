package main

import (
	"fmt"
	"math"
)

func main() {

	a := math.Pow(5, 4321);
	b := math.Pow(4, 5321);
	c := a > b;
	fmt.Printf("a=%v\n", a)
	fmt.Printf("b=%v\n", b)
	fmt.Printf("c=%v\n", c)

}
