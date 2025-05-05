package main

import "fmt"

// go function may be closures
// a closure is a function value that reference variable from outside the function body
// the function may access and assign to the reference variables in the sense the function
// is "bound" to the variables

// adder is a closure that in itself holds state
// each instance of adder holds an internal sum function
// which keep track of internal summation
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {

	// these two instances hold two internal total summation
	// that can be modified indepently based on the calling args
	// not a concept that i am used to using regularly but useful to hold state
	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}
