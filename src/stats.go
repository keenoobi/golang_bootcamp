package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func mean(array []int) float64 {
	sum := 0
	for _, num := range array {
		sum += num
	}
	result := float64(sum) / float64(len(array))

	return result
}

func meadian(numbers []int) float64 {
	n := len(numbers)
	if n%2 == 1 {
		return float64(numbers[n/2])
	}
	return float64(numbers[n/2-1]) + float64(numbers[n/2])/2
}

func mode(numbers []int) int {
	maxCount := 0
	currentCount := 1
	mode := numbers[0]

	for i := 1; i < len(numbers); i++ {
		if numbers[i] == numbers[i-1] {
			currentCount++
		} else {
			if currentCount > maxCount {
				maxCount = currentCount
				mode = numbers[i-1]
			}
			currentCount = 1
		}
	}

	if currentCount > maxCount {
		mode = numbers[len(numbers)-1]
	}

	return mode
}

func standardDeviation(numbers []int, mean float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += math.Pow(float64(num)-mean, 2)
	}
	variance := sum / float64(len(numbers))
	return math.Sqrt(variance)
}

func main() {
	meanFlag := flag.Bool("mean", false, "Print the mean")
	meadianFlag := flag.Bool("median", false, "Print the median")
	modeFlag := flag.Bool("mode", false, "Print the mode")
	sdFlag := flag.Bool("sd", false, "Print the standard deviation")

	flag.Parse()

	if !*meanFlag && !*meadianFlag && !*modeFlag && !*sdFlag {
		*meanFlag = true
		*meadianFlag = true
		*modeFlag = true
		*sdFlag = true
	}

	reader := bufio.NewReader(os.Stdin)
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
			fmt.Println("Error: Invalid input. The number is incorrect.")
			return
		}

		numbers = append(numbers, num)

	}

	if len(numbers) == 0 {
		fmt.Println("Error: No valid numbers.")
		return
	}

	sort.Ints(numbers)

	if *meanFlag {
		fmt.Printf("Mean: %.2f\n", mean(numbers))
	}
	if *meadianFlag {
		fmt.Printf("Median: %.2f\n", meadian(numbers))
	}
	if *modeFlag {
		fmt.Printf("Mode: %d\n", mode(numbers))
	}
	if *sdFlag {
		fmt.Printf("SD: %.2f\n", standardDeviation(numbers, mean(numbers)))
	}
}
