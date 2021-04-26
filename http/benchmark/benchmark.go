package main

import (
	"fmt"
	"net/http"
)

func main() {

	num := 1000
	sign := make(chan struct{}, num)

	for i := 0; i < num; i++ {
		go func(i int) {
			G()
			sign <- struct{}{}
		}(i)
	}

	// 办法2。
	for j := 0; j < num; j++ {
		<-sign
	}

}

func G() string {
	resp, err := http.Get("http://127.0.0.1:8020/api/health?query=1&test=1")
	if err != nil {
		// handle error
		fmt.Printf("request sending error: %v\n", err)
		return err.Error()
	}
	defer resp.Body.Close()
	//body, err := io.ReadAll(resp.Body)

	//url1 := "http://127.0.0.1:8020/api/health?query=1&test=1"
	//resp1, err := http.Get(url1)
	//if err != nil {
	//	fmt.Printf("request sending error: %v\n", err)
	//	return err.Error()
	//}
	//defer resp1.Body.Close()

	line1 := resp.Proto + " " + resp.Status
	return line1
}
