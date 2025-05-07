package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {

	// interface that specifies zero methods is the empty interface
	// an empty interface may hold values of any type
	// every type implements at least zero methods
	// i guess this is the Any or dynamic construct in GO
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"

	describe(i)
}
