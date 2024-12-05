package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sync"
)

func main() {
	var l, m, w bool
	flag.BoolVar(&l, "l", false, "Count all lines in a given file")
	flag.BoolVar(&m, "m", false, "Count all characters in a given file")
	flag.BoolVar(&w, "w", false, "Count all words in a given file")

	flag.Parse()
	files := flag.Args()
	if len(files) == 0 {
		fmt.Println("error: no file provided")
		return
	}

	if !l && !m && !w {
		w = true
	}

	var wg sync.WaitGroup
	if l {
		LineCounter(files, &wg)
	}

	wg.Wait()

}

func LineCounter(files []string, wg *sync.WaitGroup) {
	for _, file := range files {
		wg.Add(1)
		func(fileName string) {
			defer wg.Done()
			count, err := countLinesInFile(fileName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error processing %s: %v\n", fileName, err)
				return
			}
			fmt.Printf("%d\t%s\n", count, fileName)
		}(file)
	}
}

func countLinesInFile(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return lineCount, nil
}
