package main

import (
	"flag"
	"fmt"
)

func main() {
	old := flag.String("old", "", "for an old version of db")
	new := flag.String("new", "", "for a new version of db")

	flag.Parse()

	fmt.Println(*old, *new)

}
