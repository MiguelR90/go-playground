package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// type []T is a slice with elements of type T
	// (array) slices are dynamically sized and much more flexible
	// view into the elems of an array. slices are much more common than arrays
	// a[low: high] syntax is similar to python
	// includes the half open interval where low is included but not high
	// a[1:4] refers to elems 1, 2 and 3 but NOT 4
	var s []int = primes[1:4]
	fmt.Println(s)
}
