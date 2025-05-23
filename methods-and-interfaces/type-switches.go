package main

import "fmt"

func do(i interface{}) {
	// a type switch permits several type assertions in series
	// a regular switch but the case in a type switch specify type (not values)
	// and those values are compared against the type of the value held by the given
	// interface value

	// same declaration and syntax as a type assertion but the specific type T
	// is replaced with the keyword type
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
