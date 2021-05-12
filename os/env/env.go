package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, s := range os.Environ() { //
		kv := strings.SplitN(s, "=", 2) // unpacks "key=value"
		fmt.Printf("key:%q value:%q\n", kv[0], kv[1])
	}
}
