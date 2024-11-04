package main

import (
	"flag"
	"fmt"
)

func main() {
	inputFile := flag.String("input", "logoden-biniou/small.txt", "File to read words from")
	flag.Parse()
	if inputFile == nil || *inputFile == "" {
		fmt.Println("No input file given")
		return
	}

	main2(*inputFile)
}
