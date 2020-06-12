package main

import "fmt"

func main() {
	str := "Go爱好者"
	for i, c := range str {
		fmt.Printf("%d: %q [% x]\n", i, c, []byte(string(c)))
	}

	fmt.Printf("The string: %q\n", str)
	fmt.Printf("=> runes(char): %q\n", []rune(str))
	fmt.Printf("=> runes(hex): %x\n", []rune(str))
	fmt.Printf("=> bytes(hex): [% x]\n", []byte(str))

	/*
	   0: 'G' [47]
	   1: 'o' [6f]
	   2: '爱' [e7 88 b1]
	   5: '好' [e5 a5 bd]
	   8: '者' [e8 80 85]
	   The string: "Go爱好者"
	   => runes(char): ['G' 'o' '爱' '好' '者']
	   => runes(hex): [47 6f 7231 597d 8005]
	   => bytes(hex): [47 6f e7 88 b1 e5 a5 bd e8 80 85]

	*/
}
