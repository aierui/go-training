package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("./test.txt")
	//file, err := os.OpenFile("./test.txt", os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		fmt.Printf(" throw error %s \n", err)
	}

	buf := bufio.NewReader(file)

	for {
		// delim
		line, err := buf.ReadString('\n')
		//line, _, err := buf.ReadLine()
		if err != nil || err == io.EOF {
			break
		} else {
			//fmt.Print(line)
			fmt.Println(line)
		}
	}

}
