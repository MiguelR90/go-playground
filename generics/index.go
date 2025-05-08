package main

import "fmt"

// Type parameters
// GO funcs can be written to work on multiple types using type parameters
// Type param appear between brackets before the func args
// func Index[T comparable](s []T, x T) int {}
// comparable is "constraint" that makes it possible to use the == and != operators
// Index works for any type that supports comparable
// NOTE: constraints haven't been introduced yet. Are these like interfaces?
// yes that are exactly like interfaces!!

func Index[T comparable](s []T, x T) int {

	for i, v := range s {
		if v == x {
			return i
		}
	}

	// if no match is found we return a negavite index?
	// we could also consider returning an error value here
	return -1
}

func main() {

	// Index works for a slice of ints
	s := []int{1, 2, 3, 4, 5}
	x := 4
	fmt.Printf("Value=%v is at index=%v\n", x, Index(s, x))

	// Index also works for slice of strings
	r := []string{"cat", "dog", "bunny"}
	y := "bunny"
	fmt.Printf("Value=%v is at index=%v\n", y, Index(r, y))
}
