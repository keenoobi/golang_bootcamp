package main

import (
	"Go_Day01/internal/fscomparator"
	"flag"
	"fmt"
)

func main() {
	old := flag.String("old", "", "for an old version of filesystem dump")
	new := flag.String("new", "", "for a new version of filesystem dump")

	flag.Parse()

	if *old == "" || *new == "" {
		fmt.Println("Both --old and --new flags are required")
		return
	}

	err := fscomparator.CompareFiles(*old, *new)
	if err != nil {
		fmt.Println(err)
		return
	}
}
