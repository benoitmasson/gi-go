package main

import (
	"io"
	"math/rand/v2"
)

func countExactDistinctWords(r io.Reader) (int, int) {
	count := 0
	words := make(map[string]bool)
	for w := range Words(r) {
		words[w] = true
		count++
	}
	return count, len(words)
}

// countApproxDistinctWords implements CVM algorithm
// See https://www.quantamagazine.org/computer-scientists-invent-an-efficient-new-way-to-count-20240516/
func countApproxDistinctWords(r io.Reader, memSize int) (int, int, int) {
	count := 0
	words := make(map[string]bool, memSize)
	currentRound := 0
	for w := range Words(r) {
		count++

		// fmt.Println(w, currentRound, len(words))
		if rand.Uint64N(1<<currentRound) > 0 {
			// randomly clear memory
			// keeping/storing words becomes more and more difficult as rounds go by (two times harder by round)
			delete(words, w)
		} else {
			words[w] = true
		}

		if len(words) >= memSize {
			// memory full: cleanup half of it and move to next round
			cleanup(words)
			currentRound++
		}
	}

	return count, len(words) << currentRound, currentRound + 1
}

// cleanup randomly clears half of the keys
func cleanup(m map[string]bool) {
	for k := range m {
		if rand.IntN(2) == 0 {
			delete(m, k)
		}
	}
}
