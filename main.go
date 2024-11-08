package main

import (
	"flag"
	"fmt"
)

func main() {
	inputFile := flag.String("input", "logoden-biniou/small.txt", "File to read words from")
	memSize := flag.Int("size", 100, "Memory size limit")
	flag.Parse()
	if inputFile == nil || *inputFile == "" {
		fmt.Println("No input file given")
		return
	}
	if memSize == nil || *memSize <= 0 {
		fmt.Println("Invalid memory size given")
		return
	}

	main1()
	main2(*inputFile)
	main3(*inputFile, *memSize)
}
