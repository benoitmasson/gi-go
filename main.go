package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"time"
)

func main() {
	// config
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

	// setup
	stat, err := os.Stat(*inputFile)
	if err != nil {
		fmt.Printf("Failed to stat input file: %v\n", err)
		return
	}
	f, err := os.Open(*inputFile)
	if err != nil {
		fmt.Printf("Failed to open input file: %v\n", err)
		return
	}
	fmt.Printf("Read from file %q (%dkB)\n", *inputFile, stat.Size()/1024)
	fmt.Println()

	var (
		m0, m                             uint64
		start                             time.Time
		totalCount, distinctCount, rounds int
	)
	m0 = disableGC()

	/* Count total */
	// basic
	start = time.Now()
	totalCount = countExactTotalSheep(f)
	fmt.Printf("[ExactCountTotal] Found %d total sheep\n", totalCount)
	m = clearMemory(f)
	fmt.Printf("[ExactCountTotal] Used %dkB heap memory in %dms\n\n", m-m0, time.Since(start).Milliseconds())

	// iterator
	start = time.Now()
	totalCount = countExactTotalSheepIterator(f)
	fmt.Printf("[ExactCountTotalIterator] Found %d total sheep\n", totalCount)
	m = clearMemory(f)
	fmt.Printf("[ExactCountTotalIterator] Used %dkB heap memory in %dms\n\n", m-m0, time.Since(start).Milliseconds())

	/* Count distinct */
	// basic
	start = time.Now()
	totalCount, distinctCount = countExactDistinctSheep(f)
	fmt.Printf("[ExactCountDistinct] Found %d distinct sheep (out of %d total sheep)\n", distinctCount, totalCount)
	m = clearMemory(f)
	fmt.Printf("[ExactCountDistinct] Used %dkB heap memory in %dms\n\n", m-m0, time.Since(start).Milliseconds())
	exactCount := distinctCount

	// CVM
	start = time.Now()
	totalCount, distinctCount, rounds = countApproxDistinctSheep(f, *memSize)
	fmt.Printf("[ApproxCountDistinct(%d)] Estimated %d distinct sheep in %d rounds (out of %d total sheep)\n", *memSize, distinctCount, rounds, totalCount)
	m = clearMemory(f)
	fmt.Printf("[ApproxCountDistinct(%d)] Error rate: %.2f%%\n", *memSize, math.Abs(100*(float64(distinctCount-exactCount)/float64(exactCount))))
	fmt.Printf("[ApproxCountDistinct(%d)] Used %dkB heap memory in %dms\n\n", *memSize, m-m0, time.Since(start).Milliseconds())
}
