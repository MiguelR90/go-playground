package main

import "fmt"

func main() {
	// defers execution of a function until the surrounding function returns
	// func call + arguments are eval immediately but func is not executed until the
	// calling function returns

	// NOTE: this is novel to me. I haven't encountered defer before.
	defer fmt.Println("world")
	fmt.Println("hello")
}
