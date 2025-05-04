package main

import "fmt"

func main() {
	// variable declared without an explicit value are given their default
	// "zero" value which for each type it might vary. Example:
	// int is 0, float is 0.0, bool is false, string is the empty string
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
