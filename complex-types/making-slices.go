package main

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func main() {

	// slices can be created with the builtin make function; this is how you create dynamically sized arrays
	// make allocates a zeroed array and returns a slice that refers to that array
	a := make([]int, 5)
	printSlice("a", a)

	// to specify a capacity we can pass a third argument to make
	// b is of length 0 and capacity 5
	b := make([]int, 0, 5)
	printSlice("b", b)

	// first two elems but capcity remains 5
	c := b[:2]
	printSlice("c", c)

	// drop the first two elements from the underlying array and capcity drops to 3
	d := c[2:5]
	printSlice("d", d)

}
