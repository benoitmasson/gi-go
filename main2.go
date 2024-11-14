package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
	"unicode"
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

	start = time.Now()
	count = readIterator(f)
	fmt.Printf("[Iterator] Found %d words\n", count)
	m = clearMemory(f)
	fmt.Printf("[Iterator] Used %dkB heap memory in %dms\n\n", m-m0, time.Since(start).Milliseconds())

	start = time.Now()
	count = readChannel(f)
	fmt.Printf("[Channel] Found %d words\n", count)
	m = clearMemory(f)
	fmt.Printf("[Channel] Used %dkB heap memory in %dms\n\n", m-m0, time.Since(start).Milliseconds())
}

func readNaive(r io.Reader) int {
	count := 0

	contents, err := io.ReadAll(r)
	if err != nil {
		fmt.Printf("Failed to read input file: %v\n", err)
		return 0
	}

	for _, w := range strings.Fields(string(contents)) {
		// fmt.Println(w)
		_ = w
		count++
	}

	return count
}

func readIterator(r io.Reader) int {
	count := 0
	for w := range Words(r) {
		// fmt.Println(w)
		_ = w
		count++
	}
	return count
}

func readChannel(r io.Reader) int {
	words := make(chan string, 128)
	go func() {
		byteBuf := make([]byte, bufSize)
		wordBuf := bytes.Buffer{}
		for {
			n, err := r.Read(byteBuf)
			if err != nil {
				if errors.Is(err, io.EOF) {
					close(words)
					return
				}
				panic(fmt.Errorf("failed to read bytes: %w", err))
			}

			for i := range n {
				b := byteBuf[i]

				if !unicode.IsSpace(rune(b)) {
					wordBuf.WriteByte(b)
					continue
				}
				// b is a space

				if wordBuf.Len() == 0 {
					// ignore empty words
					continue
				}
				words <- wordBuf.String()
				// to interrupt treatment (in case of break): context should be used (explicitely)
				wordBuf.Reset()
			}
		}
	}()

	count := 0
	for w := range words {
		// fmt.Println(w)
		_ = w
		count++
	}
	return count
}
