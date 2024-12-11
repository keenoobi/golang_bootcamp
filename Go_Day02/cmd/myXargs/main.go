package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("error: no command provided")
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	scanner := bufio.NewScanner(os.Stdin)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if len(input) == 0 {
		fmt.Println("error: no input lines provided")
		os.Exit(1)
	}

	allArgs := append(args, input...)

	cmd := exec.Command(command, allArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
