package main

import (
	"fmt"
	"os"
	"time"
)

func main3(inputFile string, memSize int) {
	// setup
	stat, err := os.Stat(inputFile)
	if err != nil {
		fmt.Printf("Failed to stat input file: %v\n", err)
		return
	}
	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Failed to open input file: %v\n", err)
		return
	}
	fmt.Printf("Read from file %q (%dkB)\n", inputFile, stat.Size()/1024)
	fmt.Println()

	var (
		m0, m                     uint64
		start                     time.Time
		totalCount, distinctCount int
	)
	m0 = disableGC()

	// basic
	start = time.Now()
	totalCount, distinctCount = countExactDistinctWords(f)
	fmt.Printf("[ExactCount] Found %d distinct words (out of %d total words)\n", distinctCount, totalCount)
	m = clearMemory(f)
	fmt.Printf("[ExactCount] Used %dkB heap memory in %dms\n\n", m-m0, time.Since(start).Milliseconds())
}
