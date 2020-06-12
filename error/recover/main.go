package main

import "fmt"

//calNumber
//以下代码展示了如何将 panic 转换成 error 并返回
func calNumber() (res int, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("%s", p)
		}
	}()

	div := 0
	res = 10 / div

	return
}

func main() {
	_, err := calNumber()
	if err != nil {
		fmt.Printf("Get error: `%s`\n", err)
	}
}