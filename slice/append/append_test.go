package main

import (
	"fmt"
	"testing"
)

func TestSliceAppend(t *testing.T) {
	s := make([]int, 1)
	appendSlice(s)
	fmt.Println(s)

	/*
		=== RUN   TestSliceAppend
		[0 0 1 2 3 4 5 6 7 8 9]
		[0]
		--- PASS: TestSliceAppend (0.00s)
		PASS
	*/
}

func appendSlice(s []int) {
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	fmt.Println(s)
}

func TestSliceAppendLenAndCap(t *testing.T) {
	s1 := make([]int, 10)
	for i := 0; i < 10; i++ {
		s1 = append(s1, i)
		fmt.Printf("s1 len=%v cap=%v\n", len(s1), cap(s1))
	}

	s2 := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		s2 = append(s2, i)
		fmt.Printf("s2 len=%v cap=%v\n", len(s2), cap(s2))
	}

	fmt.Println(s1, s2)

	/*
		=== RUN   TestSliceAppendLenAndCap
		[0 0 0 0 0 0 0 0 0 0 0 1 2 3 4 5 6 7 8 9] [0 1 2 3 4 5 6 7 8 9]
		--- PASS: TestSliceAppendLenAndCap (0.00s)
		PASS
	*/
}

func TestAppendP(t *testing.T) {
	a := make([]int, 32)
	a = append(a, 1)

	fmt.Println(a)

	b := make([]int, 32)
	c := b[1:16]
	fmt.Println(b, c)

	b[3] = 9
	fmt.Println(b, c)

	fmt.Printf("The length of b: %d, The capacity of b: %d\n", len(b), cap(b))
	fmt.Printf("The length of c: %d, The capacity of c: %d\n", len(c), cap(c))

	b = append(b, 1)
	b[2] = 8
	fmt.Println(b, c)

	fmt.Printf("The length of b: %d, The capacity of b: %d\n", len(b), cap(b))
	fmt.Printf("The length of c: %d, The capacity of c: %d\n", len(c), cap(c))

	/*
		[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]
		[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
		[0 0 0 9 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] [0 0 9 0 0 0 0 0 0 0 0 0 0 0 0]
		The length of b: 32, The capacity of b: 32
		The length of c: 15, The capacity of c: 31
		[0 0 8 9 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1] [0 0 9 0 0 0 0 0 0 0 0 0 0 0 0]
		The length of b: 33, The capacity of b: 64
		The length of c: 15, The capacity of c: 31
	*/
}
