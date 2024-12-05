package main

import (
	"Go_Day02/internal/finder"
	"flag"
	"fmt"
	"os"
)

func main() {
	showDirs := flag.Bool("d", false, "Show directories")
	showFiles := flag.Bool("f", false, "Show files")
	showSymLinks := flag.Bool("sl", false, "Show symbolic links")
	ext := flag.String("ext", "", "Filter by extension")
	flag.Parse()

	if !*showDirs && !*showFiles && !*showSymLinks {
		fmt.Println("error: at least one of flags must be sprecified")
		return
	}
	if *ext != "" && !*showFiles {
		fmt.Println("error: -ext option can only be used with -f")
	}

	path := flag.Arg(0)
	if path == "" {
		path = "."
	}

	err := finder.WalkDir(path, *showDirs, *showFiles, *showSymLinks, *ext)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}
}
