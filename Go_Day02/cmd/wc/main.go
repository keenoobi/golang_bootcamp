package main

import (
	"Go_Day02/internal/wc"
	"flag"
	"fmt"
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
	wc.Counter(files, &wg, l, m, w)

	wg.Wait()

}
