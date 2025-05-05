package main

import (
	"fmt"
	"math"
)

// take a function of type float, float -> float and returns a float
// type decoration here might take a bit to get used to or we'll just
// have to read it carefully every time. it is quite combursome and verbose
func compute(fn func(float64, float64) float64) float64 {

	// this will just apply these hard-coded values to the function
	// that is passes in, i.e., fn
	return fn(3, 4)
}
func main() {

	// i guess we just got introduced to a lambda
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	// we can pass function as arguments
	// not just function we define but other function of the same type
	// math.Pow matches the function signature
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

}
