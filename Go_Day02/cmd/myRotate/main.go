package main

import (
	"Go_Day02/internal/rotater"
	"flag"
	"fmt"
	"os"
	"sync"
)

func main() {
	var archiveDir string
	flag.StringVar(&archiveDir, "a", "", "Directory to store archived logs")

	flag.Parse()

	logFiles := flag.Args()
	if len(logFiles) == 0 {
		fmt.Fprintln(os.Stderr, "error: no log files provided")
		os.Exit(1)
	}

	if archiveDir != "" {
		if _, err := os.Stat(archiveDir); os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "error: archive directory %s does not exist\n", archiveDir)
			os.Exit(1)
		}
	}

	var wg sync.WaitGroup

	for _, logFile := range logFiles {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			err := rotater.ArchiveLog(file, archiveDir)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
			}
		}(logFile)
	}
	wg.Wait()
}
