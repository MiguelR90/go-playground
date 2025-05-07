package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// GO doesn't have classes, however it has methods on types
// a method is a function with a spacial receiver argument
// the receiver appears in its own arg list between the func keyboard and name
// I GUESS? we can dispatch on multiple receiver args?
// can these be defined anywhere?

// Abs is the name of the method
// v is the receiver
// Vertex is the type
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{0.1, 0.2}
	fmt.Println(v.Abs())
}
