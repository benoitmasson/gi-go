package main

import (
	"fmt"
	"strings"
)

func main() {
	s := strings.Fields("Hello Touraine Tech 2025")

	for w := range Backward(s) {
		fmt.Println(w)
	}
}
