package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {

	if v := math.Pow(x, n); v < lim {
		fmt.Printf("if: %g < %g\n", v, lim)
		return v
	} else {
		// variables decleared inside and if short statement are also available
		// inside any of the else blocks
		fmt.Printf("else: %g >= %g\n", v, lim)
	}

	// v is no longer available since outside of the conditional scope
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 10),
	)
}
