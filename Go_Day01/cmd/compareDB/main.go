package main

import (
	"Go_Day01/internal/comparator"
	"Go_Day01/internal/parser"
	"flag"
	"fmt"
)

func main() {
	old := flag.String("old", "", "for an old version of db")
	new := flag.String("new", "", "for a new version of db")

	flag.Parse()

	if *old == "" || *new == "" {
		fmt.Println("error: flag requires a file")
		return
	}

	oldDBRreader, _, err := parser.GetReader(*old)
	if err != nil {
		fmt.Println(err)
		return
	}
	newBDReader, _, err := parser.GetReader(*new)
	if err != nil {
		fmt.Println(err)
		return
	}

	oldDBRecipes, err := oldDBRreader.Read(*old)
	if err != nil {
		fmt.Println(err)
		return
	}
	newDBRecipes, err := newBDReader.Read(*new)
	if err != nil {
		fmt.Println(err)
		return
	}

	comparator.CompareRecipes(oldDBRecipes, newDBRecipes)

}
