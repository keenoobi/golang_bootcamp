package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func parseInput(reader *bufio.Reader) ([]int, error) {
	var numbers []int
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}
		num, err1 := strconv.Atoi(input)
		if err1 != nil || num < -100000 || num > 1000000 {
			return nil, fmt.Errorf("invalid input: the number is incorrect")
		}

		numbers = append(numbers, num)

	}
	return numbers, nil
}
