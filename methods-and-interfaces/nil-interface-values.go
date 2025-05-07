package main

import "fmt"

type I interface {
	M()
}

func describe(i I) {
	fmt.Println("(%v, %T)\n", i, i)
}

func main() {
	var i I
	describe(i)

	// error!! call a method of a nil type
	// a nil interface holds neither a value nor concrete type
	// returns runtime error since there's no type inside the interface tuple
	i.M()

}
