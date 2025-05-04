package main

import (
	"fmt"
	"math"
)

func main() {
	// math.pi is undefined since package api only exposes public methods
	// the public api is denoted by a Capital letter as opposed to
	// the private only have lower case latters instead
	// fmt.Println(math.pi) // undefined error
	fmt.Println(math.Pi)
}
