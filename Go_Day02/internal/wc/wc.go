package wc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func Counter(files []string, wg *sync.WaitGroup, l, m, w bool) {
	for _, file := range files {
		wg.Add(1)
		go func(filePath string) {
			defer wg.Done()
			lineCount, charsCount, wordsCount, err := CountAll(filePath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error processing %s: %v\n", filePath, err)
				return
			}
			if l {
				fmt.Printf("%d\t%s\n", lineCount, filePath)
			}
			if m {
				fmt.Printf("%d\t%s\n", charsCount, filePath)
			}
			if w {
				fmt.Printf("%d\t%s\n", wordsCount, filePath)
			}
		}(file)
	}
}

func CountAll(fileName string) (int, int, int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, 0, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineCount := 0
	wordsCount := 0
	charsCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		charsCount += len(line)
		words := strings.Fields(line)
		wordsCount += len(words)
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, 0, err
	}
	return lineCount, charsCount, wordsCount, nil
}
