package main

import (
	"flag"
	"fmt"
	"io"
	"maps"
	"math"
	"os"
	"slices"
)

func main_stats() {
	inputFile := flag.String("input", "logoden-biniou/small.txt", "File to read words from")
	outputFile := flag.String("output", "", "File to write results to")
	outputBuckets := flag.Int("buckets", 1000, "Size of buckets for output data")
	memSize := flag.Int("size", 100, "Memory size limit")
	n := flag.Int("n", 1000, "Number of tries")
	flag.Parse()
	if inputFile == nil || *inputFile == "" {
		fmt.Println("No input file given")
		return
	}
	if memSize == nil || *memSize <= 0 {
		fmt.Println("Invalid memory size given")
		return
	}
	if n == nil || *n <= 0 {
		fmt.Println("Invalid number of tries given")
		return
	}
	if outputBuckets == nil || *outputBuckets <= 0 {
		fmt.Println("Invalid number of output buckets given")
		return
	}

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

	nums := make(map[int]int, *n)
	sum := int64(0)
	for i := range *n {
		if (i+1)%100 == 0 {
			fmt.Printf("round %d/%d…\n", i+1, *n)
		}
		_, k, _ := countApproxDistinctSheep(f, *memSize)
		nums[k]++
		sum += int64(k)
		_, err := f.Seek(0, io.SeekStart)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println()

	real := 427430
	average := float64(sum) / float64(*n)
	deviation := float64(0)
	for k, i := range nums {
		deviation += (float64(k) - average) * (float64(k) - average) * float64(i)
	}
	deviation /= float64(*n)
	deviation = math.Sqrt(deviation)

	fmt.Printf(
		"- average = %.2f (real = %d, %.3f%% error)\n- deviation = %.2f (%2f%% of average)\n",
		average, real, 100*math.Abs(average-float64(real))/float64(real),
		deviation, 100*deviation/average,
	)
	fmt.Println()

	fmt.Printf("66.7%% have less than %.2f%% error\n", 100-100*(average-deviation)/average)
	fmt.Printf("95%% have less than %.2f%% error\n", 100-100*(average-1.96*deviation)/average)
	fmt.Printf("99.6%% have less than %.2f%% error\n", 100-100*(average-3*deviation)/average)

	if outputFile != nil && *outputFile != "" {
		of, err := os.Create(*outputFile)
		if err != nil {
			panic(err)
		}

		buckets := make(map[int]int, int(average)/(*outputBuckets))
		for k, i := range nums {
			buckets[k/(*outputBuckets)*(*outputBuckets)] += i
		}
		keys := slices.Collect(maps.Keys(buckets))
		slices.Sort(keys)
		for _, k := range keys {
			_, err = fmt.Fprintf(of, "%d,%d\n", k, buckets[k])
			if err != nil {
				panic(err)
			}
		}
	}
}
