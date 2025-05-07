package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// we could save the copy by making this a reference pointer as well???
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {

	v := Vertex{3, 4}
	// the pointer method we had before rewritten as a function
	// this makes it very clear that this function could mutate
	// value methods and i guess functions always receive copy of the args
	Scale(&v, 10)
	fmt.Println(Abs(v))

}
