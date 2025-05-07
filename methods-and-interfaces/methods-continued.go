package main

import "fmt"

// you can declared methods on non-struct types too
// however the receiver type needs to be defined in the same package
// as the method. this includes the built-in types
// I GUESS? we can't extend the methods of the built-in types

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f) // casting here seems unnecessary?
}

func main() {
	f := MyFloat(-8.0)
	fmt.Println(f.Abs())

	f = MyFloat(42.0)
	fmt.Println(f.Abs())
}
