package main

import (
	"fmt"
	"iter"
	"strings"
)

func main1() {
	s := strings.Fields("Hello meetup Golang Rennes")
	for i, w := range Backward(s) {
		fmt.Println(i, w)
		if i == 2 {
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
