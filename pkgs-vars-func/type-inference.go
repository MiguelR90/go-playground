package main

import "fmt"

func main() {
	// when declaring variables via var x = or x := type can be inferred by
	// the value in the right handside of the assignment
	v := 42
	fmt.Printf("v is of type %T\n", v)

	var t float64 = 8.0
	f := t
	fmt.Printf("f is of type %T\n", f)

	// FIXME: v := t doesn't work so does that mean that vars cannot be reasigned
}
