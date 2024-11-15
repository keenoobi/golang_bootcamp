package main

import (
	"flag"
	"fmt"
	"sort"
)

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

	numbers, err := parseInput()
	if err != nil {
		fmt.Println("Error: ", err)
		return
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
