package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	showDirs := flag.Bool("d", false, "Show directories")
	showFiles := flag.Bool("f", false, "Show files")
	showSymLinks := flag.Bool("sl", false, "Show symbolic links")
	ext := flag.String("ext", "", "Filter by extension")

	flag.Parse()

	path := flag.Arg(0)
	if path == "" {
		path = "."
	}

	err := walkDir(path, *showDirs, *showFiles, *showSymLinks, *ext)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}
}

func walkDir(path string, showDirs, showFiles, showSymLinks bool, ext string) error {
	return filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			if os.IsPermission(err) {
				return nil
			}
			return err
		}

		if info.IsDir() && showDirs {
			fmt.Println(path)
		} else if !info.IsDir() && showFiles {
			if ext == "" || hasExtension(path, ext) {
				fmt.Println(path)
			}
		} else if info.Mode()&os.ModeSymlink != 0 && showSymLinks {
			processSymlynk(path)
		}
		return nil
	})
}

func hasExtension(filename, ext string) bool {
	return strings.HasSuffix(filename, "."+ext)
}

func processSymlynk(path string) {
	realPath, err := os.Readlink(path)
	if err != nil {
		fmt.Println(path, "-> [broken]")
	} else {
		fmt.Println(path, "->", realPath)
	}
}
