package main

import (
	"Go_Day01/internal/data"
	"Go_Day01/internal/parser"
	"flag"
	"fmt"
)

func main() {
	var filename string
	flag.StringVar(&filename, "f", "", "Read the file")
	flag.Parse()
	if filename == "" {
		fmt.Println("error: -f flag required")
		return
	}

	reader, extension, err := parser.GetReader(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	cakes, err := reader.Read(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	data.PrintData(extension, cakes)

}
