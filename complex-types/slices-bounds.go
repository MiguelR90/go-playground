package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}

	// when slicing you may omit the low and high bounds and they'll default
	// to their correspoing values low -> 0 and high -> lenght of array
	s = s[:]
	fmt.Println(s)

	// low is 0
	s = s[:4]
	fmt.Println(s)

	// high is lengh of slice???
	// yes high is the 4th index since that the length of this slices
	// just like python lists
	s = s[2:]
	fmt.Println(s)

}
