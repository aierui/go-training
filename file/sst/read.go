package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	read2()
}

func read1() {
	f, err := ioutil.ReadFile("file/sst/sst_0_0.sst")
	if err != nil {
		fmt.Println("read fail", err)
	}
	rtn := string(f)
	fmt.Print(rtn)
}

func read2() {

	f, err := os.Open("file/sst/sst_0_0.sst")
	if err != nil {
		fmt.Println("read file fail", err)
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd fail", err)
	}

	rtn := string(fd)
	fmt.Print(rtn)

}
