package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {

	// a slice has both a length and a capacity
	// length := number of elements slice contains
	// capacity := number of elements in the underlying array, COUNTING FROM THE FIRST ELEMENT IN THE SLICE
	// slices can only be extended / resliced from the end
	// once you start chopping from the front you cant extend from the beginning of the underlying array
	// you are not protected from extending beyond the array's capacity which results in an error
	// NOTE: this will take a while to get used to
	s := []int{0, 1, 2, 3, 4, 5}

	// slice the slice to give it zero length
	s = s[:0]
	printSlice(s)

	// extend the length of the slice from the high side
	s = s[:4]
	printSlice(s)

	// drop first two values; reduces the capacity of the slice
	s = s[2:]
	printSlice(s)

	// first two values are gone!!! but okay to expand from the high side
	// but the capacity is now 4 as opposed to 6 so the 3rd element is 5
	// so we should see 5 appear in the slice!!!
	s = s[:4]
	printSlice(s)

	// error slice bouds out of range
	// s = s[:10]
	// printSlice(s)

}
