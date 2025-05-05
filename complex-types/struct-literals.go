package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// multi var declaration syntax / block
// NOTE: = replaces the usual := declaration
var (
	// you can list and declare just a subset of the field by using Name: syntax
	// ordering of the field is irrelevant and for any missing field it will be
	// "zero" declared? (this might be a foot fun)
	// spacial prefix & returns a pointer to the struct value
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0 are implicit (no optional like in python)
	p  = &Vertex{1, 2} // has type *Vertex ( no idea what this means)

)

func main() {
	fmt.Println(v1, v2, v3, p)
}
