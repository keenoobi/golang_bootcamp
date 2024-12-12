package finder

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func WalkDir(path string, showDirs, showFiles, showSymLinks bool, ext string) error {
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
		return
	}

	// Проверяем, существует ли файл, на который указывает ссылка
	if _, err := os.Stat(realPath); os.IsNotExist(err) {
		fmt.Println(path, "-> [broken]")
	} else {
		fmt.Println(path, "->", realPath)
	}
}
