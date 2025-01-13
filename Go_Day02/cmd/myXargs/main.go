package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "error: no command provided")
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	_, err := exec.LookPath(command)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: command not found: %s\n", command)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	var input []string
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			input = append(input, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}

	if len(input) == 0 {
		fmt.Fprintln(os.Stderr, "error: no input lines provided")
		os.Exit(1)
	}

	allArgs := append(args, input...)

	cmd := exec.Command(command, allArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
