package main

import (
	"flag"
	"fmt"
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
		m0, m      uint64
		start      time.Time
		totalCount int
	)
	m0 = disableGC()

	// Count total //
	// basic
	start = time.Now()
	totalCount = countTotalSheep(f)
	fmt.Printf("[CountTotal] Found %d total sheep\n", totalCount)
	m = clearMemory(f)
	fmt.Printf("[CountTotal] Used %dkB heap memory in %dms\n\n", m-m0, time.Since(start).Milliseconds())

	/*
















		// iterator
		start = time.Now()
		totalCount = countTotalSheepIterator(f)
		fmt.Printf("[CountTotalIterator] Found %d total sheep\n", totalCount)
		m = clearMemory(f)
		fmt.Printf("[CountTotalIterator] Used %dkB heap memory in %dms\n\n", m-m0, time.Since(start).Milliseconds())

		/*


















		// Count distinct //
		var distinctCount int

		// basic
		start = time.Now()
		totalCount, distinctCount = countExactDistinctSheep(f)
		fmt.Printf("[ExactCountDistinct] Found %d distinct sheep (out of %d total sheep)\n", distinctCount, totalCount)
		m = clearMemory(f)
		fmt.Printf("[ExactCountDistinct] Used %dkB heap memory in %dms\n\n", m-m0, time.Since(start).Milliseconds())

		/*





















		exactCount := distinctCount

		// CVM
		var rounds int

		start = time.Now()
		totalCount, distinctCount, rounds = countApproxDistinctSheep(f, *memSize)
		fmt.Printf("[ApproxCountDistinct(%d)] Estimated %d distinct sheep in %d rounds (out of %d total sheep)\n", *memSize, distinctCount, rounds, totalCount)
		m = clearMemory(f)
		fmt.Printf("[ApproxCountDistinct(%d)] Error rate: %.2f%%\n", *memSize, math.Abs(100*(float64(distinctCount-exactCount)/float64(exactCount))))
		fmt.Printf("[ApproxCountDistinct(%d)] Used %dkB heap memory in %dms\n\n", *memSize, m-m0, time.Since(start).Milliseconds())

	/* */
}
