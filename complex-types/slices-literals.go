package main

import "fmt"

func main() {

	// slice literal is like an array literal without the lenght
	// this is an array literal and it creates the same array as before
	// and a slice that references it.
	// this is a much more convenient way of creating arrays without having
	// to specify the length
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, true, false, true}
	fmt.Println(r)

	// okay this wild sytax. create a struct on the fly so the result is an array
	// of structs but we get a slice? (this feels confusing and complex)
	// FIXME: a potential point of friction for my brain
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, false},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
