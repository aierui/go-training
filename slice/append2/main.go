package main

import (
	"log"
	"net/http"
	"net/http/pprof"
)

func main() {

	s()

	rows := make([]string, 10)

	for i := 0; i < 10; i++ {
		rows = append(rows, "111")
	}

	for {
		rows = append(rows, "111")
	}

}

func s() {
	mux := http.NewServeMux()
	mux.HandleFunc("/profile", pprof.Profile)
	log.Fatal(http.ListenAndServe(":7777", mux))
}
