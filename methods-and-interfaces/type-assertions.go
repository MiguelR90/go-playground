package main

import "fmt"

func main() {
	var i interface{} = "string"

	// a type assertion provides access to an interface value's underlying concrete value
	// statement asserts that the interface value i holds the concrete type string
	// and assigns the underlying value to var s
	// if i doesn't hold a string it'll trigger a panic
	s := i.(string)
	fmt.Println(s)

	// to test whether an interface value holds a specific type, a type assertion can return two values
	// the underlying value and a boolean value that reports whether the assertion succeeded
	s, ok := i.(string)
	fmt.Println(s, ok)

	// error as value!!!
	// this type assertion results in an error but it's handle
	// by explicitly checking f, ok since it's return as a value
	f, ok := i.(float64)
	fmt.Println(f, ok)

	// this is an unhandles error. we don't check f, ok
	// type assertion results and a panic and halts
	f = i.(float64) // panic
	fmt.Println(f)
}
