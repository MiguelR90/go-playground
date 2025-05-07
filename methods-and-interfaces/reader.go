package main

import (
	"fmt"
	"io"
	"strings"
)

// io package specifies the io.Reader interface which represents the read end of a stream of data
// the interface has a read method:
// func (T) Read(b []bytes) (n int, err error)
// read populates the given byte slice with data and returns the number of bytes populated and an
// error value. it return an io.EOF error when the stream ends
func main() {

	// this type? has a method that implements that interface?
	r := strings.NewReader("hello reader!")

	// this is the byte string that get's populated by the reader interface
	// this totally controls the size of the byte slice that get's populated
	// the reader will fill this slice with data as we iterate over it
	b := make([]byte, 8)

	// this is like a while true construct
	for {

		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])

		// exit when we reach the end of the file stream via break keyword?
		if err == io.EOF {
			break
		}
	}
}
