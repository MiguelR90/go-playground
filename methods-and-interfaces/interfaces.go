package main

import (
	"fmt"
	"math"
)

// an interface is defined as a set of method signatures
// this is very similar to structural subtyping in python
// as long as a struct adheres to this interface then functions
// can operate against it
// value of interface type can hold any value that implements those methods

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // a MyFloat implements Abser

	// NOTE: that this is the pointer and not the value
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// NOTE: this is not something the compiler catches either
	// a = v

	fmt.Println(a.Abs())

}
