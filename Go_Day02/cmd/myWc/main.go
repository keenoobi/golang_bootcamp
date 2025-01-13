package main

import (
	"Go_Day02/internal/wc"
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

	if l && m || l && w || m && w {
		fmt.Fprintln(os.Stderr, "error: too many flags, only one at a time")
		os.Exit(1)
	}

	files := flag.Args()

	if len(files) == 0 {
		fmt.Fprintln(os.Stderr, "error: no file provided")
		os.Exit(1)
	}

	if !l && !m && !w {
		w = true
	}

	var wg sync.WaitGroup
	wc.Counter(files, &wg, l, m, w)

	wg.Wait()
}
