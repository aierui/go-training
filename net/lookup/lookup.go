package main

import (
	"fmt"
	"net"
)

func main() {

	addr := "10.15.42.17"

	rtn, err := net.LookupAddr(addr)

	fmt.Println(rtn, err)
}



