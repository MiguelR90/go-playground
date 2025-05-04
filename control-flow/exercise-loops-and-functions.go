package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {

	// Compute the sqrt of x in a loop by adjusting an initial guess
	z := 1.0 // initial guess for sqrt of x; alternatively we could use z := float64(1)
	// FIXME: tolarance is hardcoded
	tol := 1e-9 // precision tolerance. how close do we want out calculation to be

	// Consider using relative difference as an alternative stoping criteria
	// having to do a while loop with a for keywords is very unintuitive for me
	for math.Abs(z*z-x) > tol {
		// Adjust z based on how close Z^2 is to x
		// Upate step is given by the following:
		// z_{j+1} -= (z_j * z_j - 1) / (2 * z)
		z -= (z*z - x) / (2 * z) // update is newton's method which works very well for square root func

	}

	return z
}

func main() {
	fmt.Println(
		Sqrt(1),
		Sqrt(2),
		Sqrt(3),
		Sqrt(4),
		Sqrt(9),
		Sqrt(16),
	)
}
