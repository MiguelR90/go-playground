package main

import "fmt"

// fibonacci is a function that returns a function that returns an int.
func fibonacci() func() int {

	// don't know how to do it without manually triggering the first return
	// with some sort of initizlizer marker
	var current, next int = 0, 1

	return func() int {
		// the current fibonacci number however i can't return yet
		// since i need to perform an update to the current and next variables
		// this is a temp variable just to hold the return value for this iteration
		fib := current

		// this is the update step setting up the vars for the next call
		current, next = next, current+next

		// the ability to yield in python make this so much easier
		// generator and closures are similar since they are able to hold
		// internal state but perhaps generators a bit superior because
		// of their explicitness
		return fib
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
