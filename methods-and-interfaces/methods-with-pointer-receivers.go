package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// there are two reasons to use pointer receivers
// 1. method can modify the value that it's receiver points to
// 2. avoid copying the value on each method call. more eff for large struct

// RULE: ONE OR THE OTHER BUT NOT BOTH!!!!
// all methods of a given type should have either value or pointer receivers
// but not a mixture of both

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// NOTE: doesn't need to be a pointer receiver because it doesn't mod the receiver in place
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
