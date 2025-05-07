package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// this causes an infinite loop and crashes
	// return fmt.Sprintf("cannot Sqrt negative number: %v", e)
	// The Go runtime uses type assertions to determine how to handle values passed
	// to fmt.Sprint. When the value is of type error, fmt.Sprint calls its Error()
	// method. By converting the value to another type, such as float64, you effectively
	// tell fmt.Sprint to treat it as a plain number rather than an error.
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	// just using the math implementation since the important concept
	// is the error handling and raising
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
