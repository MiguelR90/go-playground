package main

import "fmt"

func main() {

	// type [n]T is an array of n values of type T
	// an array's length if part of it's type so they can't be resize
	// GO provides other conveniences for resizing and working with arrays

	// a is an array of string with size = 2
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// declaration and assignment syntax
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

}
