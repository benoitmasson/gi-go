package main

import (
	"context"
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

func BackwardChannel[E any](ctx context.Context, s []E, c chan<- E) {
	defer close(c)

	for i := len(s) - 1; i >= 0; i-- {
		select {
		case <-ctx.Done():
			fmt.Println("-- break")
			return
		default:
			c <- s[i]
		}
	}

	fmt.Println("-- done")
}
