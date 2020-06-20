# Go 范型示例


[Run](http://go2goplay.golang.org/)

```go2
package main

import (
	"fmt"
)

func Print(type T)(s []T) {
	for _, v := range s {
		fmt.Print(v)
	}
}

func main() {
	Print([]string{"Hello, ", "go-training\n"})
}
```