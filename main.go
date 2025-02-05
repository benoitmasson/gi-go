package main

import (
	"context"
	"fmt"
	"iter"
	"strings"
)

func main() {
	s := strings.Fields("Hello Touraine Tech 2025")
	for i, w := range Backward(s) {
		fmt.Println(w)
		if i == 2 {
			break
		}
	}

	fmt.Println()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c := make(chan string)

	go BackwardChannel(ctx, s, c)

	for w := range c {
		fmt.Println(w)
		if w == "Tech" {
			cancel()
			break
		}
	}

	fmt.Println()

	next, stop := iter.Pull2(Backward(s))
	defer stop()
	for {
		i, w, ok := next()
		if !ok {
			break
		}
		fmt.Println(i, w)
		if i == 2 {
			stop()
			break
		}
	}
}
