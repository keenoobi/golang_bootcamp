package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"project/internal/parser"
	"project/internal/stats"
	"sort"
)

func main() {
	meanFlag := flag.Bool("mean", false, "Print the mean test")
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
	numbers, err := parser.ParseInput(reader)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if len(numbers) == 0 {
		fmt.Println("Error: no valid numbers")
		return
	}

	sort.Ints(numbers)

	if *meanFlag {
		fmt.Printf("Mean: %.2f\n", stats.Mean(numbers))
	}
	if *meadianFlag {
		fmt.Printf("Median: %.2f\n", stats.Meadian(numbers))
	}
	if *modeFlag {
		fmt.Printf("Mode: %d\n", stats.Mode(numbers))
	}
	if *sdFlag {
		fmt.Printf("SD: %.2f\n", stats.StandardDeviation(numbers, stats.Mean(numbers)))
	}
}
