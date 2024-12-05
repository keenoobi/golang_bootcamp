package wc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func Counter(files []string, wg *sync.WaitGroup, countFunc func(string) (int, error)) {
	for _, file := range files {
		wg.Add(1)
		go func(filePath string) {
			defer wg.Done()
			count, err := countFunc(filePath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error processing %s: %v\n", filePath, err)
				return
			}
			fmt.Printf("%d\t%s\n", count, filePath)
		}(file)
	}
}

func CountLines(fileName string) (int, error) {
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

func CountWords(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	wordsCount := 0

	for scanner.Scan() {
		word := strings.Fields(scanner.Text())
		wordsCount += len(word)
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return wordsCount, nil
}

func CountChars(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	charsCount := 0
	for scanner.Scan() {
		charsCount += len(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return charsCount, nil
}
