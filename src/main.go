package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	p "project/parser"
	s "project/stats"
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
	numbers, err := p.ParseInput(reader)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	if len(numbers) == 0 {
		fmt.Println("error: no valid numbers")
		return
	}

	sort.Ints(numbers)

	if *meanFlag {
		fmt.Printf("Mean: %.2f\n", s.Mean(numbers))
	}
	if *meadianFlag {
		fmt.Printf("Median: %.2f\n", s.Meadian(numbers))
	}
	if *modeFlag {
		fmt.Printf("Mode: %d\n", s.Mode(numbers))
	}
	if *sdFlag {
		fmt.Printf("SD: %.2f\n", s.StandardDeviation(numbers, s.Mean(numbers)))
	}
}
