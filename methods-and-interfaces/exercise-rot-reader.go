package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// brute force mapping
// func Rot13Map() map[byte]byte {
// 	X := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
//
// 	Y := []byte{'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm'}
//
// 	m := make(map[byte]byte)
// 	for i := range X {
// 		m[X[i]] = Y[i]
// 	}
// 	return m
// }

// this is a much better way to accomplish this mapping
// 1. need to understand the ordering of a-z and A-Z
// 2. what is byte subtraction? not quite intuitive for me yet
// 3. mod makes sense but why subtract?
func rot13(c byte) byte {
	switch {
	case 'A' <= c && c <= 'Z':
		return 'A' + (c-'A'+13)%26
	case 'a' <= c && c <= 'z':
		return 'a' + (c-'a'+13)%26
	default:
		return c
	}
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)

	// TODO: do we check this here or after the loop?
	if err == io.EOF {
		return n, err
	}

	for i := range b {
		b[i] = rot13(b[i])
	}

	// return n or the len(b)? does is matter?
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
