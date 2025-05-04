package main

import "fmt"

// variable initialization can be included in var def
var i, j int = 1, 2

func main() {
	// note: if initializer is present then the type of the vars can be omitted
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
