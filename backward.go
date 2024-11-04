package main

import (
	"fmt"
	"iter"
)

func Backward[E any](s []E) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(i, s[i]) {
				fmt.Println("-- break")
				return
			}
		}
	}
}
