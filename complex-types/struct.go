package main

import "fmt"

// struct is a collection of fields
// same as struct in all other languages like: Elixir
// like type dicts (and data classes) in python
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{42, 99})
}
