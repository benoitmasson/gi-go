package main

import (
	"io"
	"math/rand/v2"
	"strings"
)

/* Count total */

func countExactTotalSheep(r io.Reader) int {
	data, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	sheepNames := strings.Fields(string(data))
	count := 0
	for _, name := range sheepNames {
		admireSheep(name)
		count++
	}
	return count
}

func admireSheep(name string) {
	/*
		        __  _
		    ,-:'  `; `-._
		   (_,           )
		 ,'o"(            )>
		(__,-'            )
		   (             )
		    `-'._.--._.-'
		       || || ||
	*/
}

func countExactTotalSheepIterator(r io.Reader) int {
	count := 0
	for sheepName := range ReadFields(r) {
		admireSheep(sheepName)
		count++
	}
	return count
}

/* Count distinct */

func countExactDistinctSheep(r io.Reader) (int, int) {
	count := 0
	sheepNames := make(map[string]bool)
	for sheepName := range ReadFields(r) {
		sheepNames[sheepName] = true
		count++
	}
	return count, len(sheepNames)
}

// countApproxDistinctSheep implements CVM algorithm
// See https://www.quantamagazine.org/computer-scientists-invent-an-efficient-new-way-to-count-20240516/
func countApproxDistinctSheep(r io.Reader, memSize int) (int, int, int) {
	count := 0
	sheepNames := make(map[string]bool, memSize)
	currentRound := 0
	for sheepName := range ReadFields(r) {
		count++

		// fmt.Println(w, currentRound, len(words))
		if rand.Uint64N(1<<currentRound) > 0 {
			// randomly clear memory
			// keeping/storing words becomes more and more difficult as rounds go by (two times harder by round)
			delete(sheepNames, sheepName)
		} else {
			sheepNames[sheepName] = true
		}

		if len(sheepNames) >= memSize {
			// memory full: cleanup half of it and move to next round
			cleanup(sheepNames)
			currentRound++
		}
	}

	return count, len(sheepNames) << currentRound, currentRound + 1
}

// cleanup randomly clears half of the keys
func cleanup(m map[string]bool) {
	for k := range m {
		if rand.IntN(2) == 0 {
			delete(m, k)
		}
	}
}
