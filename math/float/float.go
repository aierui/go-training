package main

import (
	"fmt"
	"strconv"
)

func main() {
	f := 100.12345678901234567890123456789
	fmt.Println(strconv.FormatFloat(f, 'b', 5, 32))
	// 13123382p-17
	fmt.Println(strconv.FormatFloat(f, 'e', 5, 32))
	// 1.00123e+02
	fmt.Println(strconv.FormatFloat(f, 'E', 5, 32))
	// 1.00123E+02
	fmt.Println(strconv.FormatFloat(f, 'f', 5, 32))
	// 100.12346
	fmt.Println(strconv.FormatFloat(f, 'g', 5, 32))
	// 100.12
	fmt.Println(strconv.FormatFloat(f, 'G', 5, 32))
	// 100.12
	fmt.Println(strconv.FormatFloat(f, 'b', 30, 32))
	// 13123382p-17
	fmt.Println(strconv.FormatFloat(f, 'e', 30, 32))
	// 1.001234588623046875000000000000e+02
	fmt.Println(strconv.FormatFloat(f, 'E', 30, 32))
	// 1.001234588623046875000000000000E+02
	fmt.Println(strconv.FormatFloat(f, 'f', 30, 32))
	// 100.123458862304687500000000000000
	fmt.Println(strconv.FormatFloat(f, 'g', 30, 32))
	// 100.1234588623046875
	fmt.Println(strconv.FormatFloat(f, 'G', 30, 32))
	// 100.1234588623046875

	fmt.Println()
	f1 := strconv.FormatFloat(1.56*100, 'f', 10, 64)
	fmt.Println(f1)
	fmt.Println(strconv.ParseInt(f1, 10, 64))

	fmt.Println()
	f2 := strconv.FormatFloat(1.506*100, 'f', 0, 64)
	fmt.Println(f2)
	fmt.Println(strconv.ParseInt(f2, 10, 64))

}
