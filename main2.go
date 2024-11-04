package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main2(inputFile string) {
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
		m0, m uint64
		start time.Time
		count int
	)
	m0 = disableGC()

	start = time.Now()
	count = readNaive(f)
	fmt.Printf("[Naive] Found %d words\n", count)
	m = clearMemory(f)
	fmt.Printf("[Naive] Used %dkB heap memory in %dms\n\n", m-m0, time.Since(start).Milliseconds())
}

func readNaive(r io.Reader) int {
	count := 0
	return count
}
