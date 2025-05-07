package main

import "golang.org/x/tour/reader"

type MyReader struct {
	Value byte
}

func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = r.Value
	}

	// why do i return the len?
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{'A'})
}
