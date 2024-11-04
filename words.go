package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"iter"
	"unicode"
)

const bufSize = 1024

func Words(r io.Reader) iter.Seq[string] {
	return func(yield func(string) bool) {
		byteBuf := make([]byte, bufSize)
		wordBuf := bytes.Buffer{}
		for {
			n, err := r.Read(byteBuf)
			if err != nil {
				if errors.Is(err, io.EOF) {
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
				if !yield(wordBuf.String()) {
					return
				}
				wordBuf.Reset()
			}
		}
	}
}
