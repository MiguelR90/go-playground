package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// the difference between a method and a function is just the receiver argument
// and the calling syntax... methods can use the type T.method "dot syntax"
// a function need to take the type / instance as an argument
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
