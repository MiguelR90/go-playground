package main

import "fmt"

func main() {

	// the "zero" value of a slice is nil
	// a nil slice has a length and capacity of 0 and has no underlying array

	// a nil slice
	var s []int
	fmt.Println(s, len(s), cap(s))

	if s == nil {
		fmt.Println("nil!")
	}
}
