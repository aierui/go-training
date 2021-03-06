package main

import "fmt"

func main() {
	// 示例1。
	s1 := make([]int, 5)
	fmt.Printf("The length of s1: %d\n", len(s1))
	fmt.Printf("The capacity of s1: %d\n", cap(s1))
	fmt.Printf("The value of s1: %d\n", s1)
	s2 := make([]int, 5, 8)
	fmt.Printf("The length of s2: %d\n", len(s2))
	fmt.Printf("The capacity of s2: %d\n", cap(s2))
	fmt.Printf("The value of s2: %d\n", s2)
	fmt.Println()

	str2 := make([]string, 5, 8)
	fmt.Printf("The length of str2: %v\n", len(str2))
	fmt.Printf("The capacity of str2: %v\n", cap(str2))
	fmt.Printf("The value of str2: %v\n", str2)
	fmt.Println()

	// 示例2。
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	fmt.Printf("The length of s4: %d\n", len(s4))
	fmt.Printf("The capacity of s4: %d\n", cap(s4))
	fmt.Printf("The value of s4: %d\n", s4)
	fmt.Println()

	// 示例3。
	s5 := s4[:cap(s4)]
	fmt.Printf("The length of s5: %d\n", len(s5))
	fmt.Printf("The capacity of s5: %d\n", cap(s5))
	fmt.Printf("The value of s5: %d\n", s5)

	/*
		The length of s1: 5
		The capacity of s1: 5
		The value of s1: [0 0 0 0 0]
		The length of s2: 5
		The capacity of s2: 8
		The value of s2: [0 0 0 0 0]

		The length of s4: 3
		The capacity of s4: 5
		The value of s4: [4 5 6]

		The length of s5: 5
		The capacity of s5: 5
		The value of s5: [4 5 6 7 8]
	*/
}
