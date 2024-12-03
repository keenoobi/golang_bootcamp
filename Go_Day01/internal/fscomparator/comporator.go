package fscomparator

import (
	"bufio"
	"fmt"
	"os"
)

func readFileToMap(filename string) (map[string]bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	fileMap := make(map[string]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileMap[scanner.Text()] = true
	}

	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reding file: %w", err)
	}
	return fileMap, nil
}

func CompareFiles(oldDumpFile, newDumpFile string) error {
	oldDumpMap, err := readFileToMap(oldDumpFile)
	if err != nil {
		return err
	}

	file, err := os.Open(newDumpFile)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		path := scanner.Text()
		if _, exist := oldDumpMap[path]; exist {
			delete(oldDumpMap, path)
		} else {
			fmt.Printf("ADDED %s\n", path)
		}
	}

	if err = scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}
	for path := range oldDumpMap {
		fmt.Printf("REMOVED %s\n", path)
	}
	return nil
}
