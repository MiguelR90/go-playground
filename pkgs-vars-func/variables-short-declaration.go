package main

import "fmt"

func main() {
	// := is shorthand declaration for variables
	// used instead of var but likely need to be used with initalizers
	// this returns an error: t := int error int is not an expression
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
