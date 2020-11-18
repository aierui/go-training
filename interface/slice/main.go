package main

import "fmt"

func printAll(vals []interface{}) { //1
	for _, val := range vals {
		fmt.Println(val)
	}
}

func main() {
	// 执行之后竟然会报 cannot use names (type []string) as type []interface {} in argument to printAll 错误，why？
	//names := []string{"stanley", "david", "oscar"}
	//printAll(names)

	// go 不会对 类型是interface{} 的 slice 进行转换 。
	// 为什么 go 不帮我们自动转换，一开始我也很好奇，
	// 在 go 的 wiki 中找到了答案 https://github.com/golang/go/wiki/InterfaceSlice

	// 手动实现

	dataSlice := []int{1, 2, 3}
	var interfaceSlice []interface{} = make([]interface{}, len(dataSlice))
	for i, d := range dataSlice {
		interfaceSlice[i] = d
	}
	printAll(interfaceSlice)

}
