package main

import (
	"fmt"
)

func main() {

	for _, r := range `-\|/` {
		fmt.Printf("\n%c", r)
		//time.Sleep(1000 * time.Millisecond)
	}

	fmt.Println()

	for k, v := range []int{1, 2, 3} {
		fmt.Printf("k => %v v => %v \n", k, v)
	}

	// 循环永动机
	arr := []int{1, 2, 3}

	for k, v := range arr {
		arr = append(arr, v)
		fmt.Printf("k => %v v => %v \n", k, v)
	}

	fmt.Printf("arr => %v \n", arr)

	fmt.Println()

	// 上述代码的输出意味着循环只遍历了原始切片中的三个元素，我们在遍历切片时追加的元素不会增加循环的执行次数，所以循环最终还是停了下来。

	// 神奇的指针

	arr2 := []int{1, 2, 3}
	newArr2 := []*int{}
	newArr3 := []*int{}

	for k, v := range arr2 {
		fmt.Printf("k => %v v => %v \n", k, v)
		newArr2 = append(newArr2, &v)
		newArr3 = append(newArr3, &arr[k])
	}

	for k, v := range newArr2 {
		fmt.Printf("newArr2 k => %v v => %v \n", k, *v)
	}

	for k, v := range newArr3 {
		fmt.Printf("newArr3 k => %v v => %v \n", k, *v)
	}

	fmt.Println()

	// 遍历清空数组

	arr4 := []int{1, 2, 3}

	for k, v := range arr4 {
		arr4[k] = 0
		fmt.Printf("arr4 k => %v v => %v \n", k, v)
	}

	fmt.Println()

	// 随机遍历

	hash := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}

	for k, v := range hash {
		fmt.Printf("hash k => %v v => %v \n", k, v)
		//println(k, v)
	}

}
