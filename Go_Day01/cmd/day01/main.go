package main

import (
	"Go_Day01/internal/parser"
	"flag"
	"fmt"
	"path/filepath"
)

func main() {
	var filename string
	flag.StringVar(&filename, "f", "", "Read the file")
	flag.Parse()

	extension := filepath.Ext(filename)

	var reader parser.DBReader

	switch extension {
	case ".json":
		reader = parser.JSONReader{}
		fmt.Println("This is json file")
	case ".xml":
		reader = parser.XMLReader{}
		fmt.Println("This is xml file")
	default:
		fmt.Println("wrong file")
		return
	}

	cakes, err := reader.Read(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cakes.String())

}
