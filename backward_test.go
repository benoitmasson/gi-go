package main

import (
	"iter"
	"slices"
	"strings"
	"testing"
)

func TestBackward(t *testing.T) {
	input := strings.Fields("Hello Touraine Tech 2025")
	expected := slices.Clone(input)
	slices.Reverse(expected)

	next, stop := iter.Pull2(Backward(input))
	defer stop()

	for i, expectedWord := range expected {
		expectedIndex := len(expected) - 1 - i
		index, word, ok := next()
		if !ok {
			t.Fatalf("Iterator stopped, expected word %q at index %d", expectedWord, expectedIndex)
		}
		if expectedIndex != index {
			t.Errorf("Expected index %d, got index %d", expectedIndex, index)
		}
		if expectedWord != word {
			t.Errorf("Expected word %q at index %d, got word %q", expectedWord, expectedIndex, word)
		}
	}

	index, word, ok := next()
	if ok {
		t.Errorf("Unexpected word %q at index %d", word, index)
	}
}
