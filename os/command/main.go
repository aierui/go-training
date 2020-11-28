package main

import (
	"fmt"
	"os"
	"os/exec"
)

// go run main.go run /bin/bash
func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("unknown command")

	}
}

func run() {
	fmt.Printf("Command Running: %v\n", os.Args[2:])
	cmd := exec.Command(os.Args[2], os.Args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("new errors %v\n", err)
	}

	fmt.Println("running successful!")
}
