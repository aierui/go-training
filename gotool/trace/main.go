package main

import (
	"net/http"
	_ "net/http/pprof"
)

func main() {

	_ = http.ListenAndServe(":7777", nil)

	m()

	// web visual view （直接用浏览器查看 简直太方便了）
	// http://127.0.0.1:7777/debug/pprof/

	// https://golang.org/doc/diagnostics#profiling
	// trace
	//step1. go run main.go
	//step2. curl 'http://127.0.0.1:7777/debug/pprof/trace?seconds=20' > trace.out
	//step2. go tool trace trace.out


	// cpu
	// step2. curl 'http://127.0.0.1:7777/debug/pprof/profile?seconds=20' > cpu
	// step3. go tool pprof -http=:9090 cpu
	// or.    go tool pprof http://127.0.0.1:7777/debug/pprof/profile?seconds=20

	// mem
	//step2. curl 'http://127.0.0.1:7777/debug/pprof/heap' > heap
	//step3. go tool pprof -http=:9090 heap
}

func m() {
	rows := make([]string, 10)

	for i := 0; i < 10; i++ {
		rows = append(rows, "111")
	}

	for {
		rows = append(rows, "111")
	}
}
