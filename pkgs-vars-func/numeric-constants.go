package main

import "fmt"

// numeric constants are high precision value
// an untyped constant takes the type needed by its context

const (
	// create a huge number by shifting a 1 bit left 100 places
	// i.e., the binary number that is 1 followed by 100 zeros
	Big = 1 << 100
	// shift it right again 99 places so we end up with 1<<1, or 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	fmt.Println(needInt(Small))
	// cannot be used because it's too large?
	// fmt.Println(needInt(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
