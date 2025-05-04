package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {

	// GO's short statement which executes before the if condition
	// Similar to python's walrus operator (except for var scoping rules)
	// variables decleared by the statement are only in scope until the end of the if
	if v := math.Pow(x, n); v < lim {
		return v
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
