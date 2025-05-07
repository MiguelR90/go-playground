package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// func that take a value must take a value of that specific type
// NOTE: cannot call AbsFunc(&v) with a pointer!! complier error
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// while methods with value receivers can take a value or pointer
// it still makes a copy if we pass an explicit pointer? or does it
// modify in place?? (copy would be my guess)... this stuff is a bit confusing
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
