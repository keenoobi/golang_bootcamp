package main

import (
	"Go_Day01/internal/converter"
	"Go_Day01/internal/data"
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
	var printer func(data.Cakes)

	switch extension {
	case ".json":
		reader = parser.JSONReader{}
		printer = converter.PrintXML
		fmt.Println("This is json file")
	case ".xml":
		reader = parser.XMLReader{}
		printer = converter.PrintJSON
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
	printer(cakes)

}
