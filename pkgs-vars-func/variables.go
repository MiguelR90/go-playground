package main

import "fmt"

// var keyword is used to defined a list of variables with type being the
// last entry in the list
var c, python, java bool

func main() {
	// vars can be defined at package scope or funciton scope
	// NOTE: no protection from readying and mutating global scope vars!!
	var i int
	fmt.Println(i, c, python, java)
}
