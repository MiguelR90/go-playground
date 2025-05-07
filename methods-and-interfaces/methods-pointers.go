package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// you can declare methods with pointer receivers
// with the literal syntax (t *T) where T cannot itself be a pointer such as *int
// methods with pointer receivers can modified the value otherwise mod will be ignored
// with a value receiver the method operates on a copy of the original T type
// other function have a similar behaviour (always workign with copies????)
// pointer receivers are more common than value receiver

// this method mutates the receiver in place
// bc it mutates then this needs to be a pointer method
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// this method is a computation and return a result
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {

	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())

}
