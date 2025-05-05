package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {

	// map literals are like struct literals but the keys are requires
	// this is how you define a map statically, i.e., is not initializer
	// or grown dynamically
	m := map[string]Vertex{
		"bell-labs":  Vertex{1.0, 1.0},
		"other-labs": Vertex{1.0, 1.0},
		"stanford":   Vertex{1.0, 1.0},
		"harvard":    Vertex{1.0, 1.0},
	}

	fmt.Println(m)

	// if the top-level is just a type name, you can omit it from the elements
	// of the literal. so it'll know that the values are structs of type Vertex
	o := map[string]Vertex{
		"hello": {.1, .2},
		"world": {.3, .4},
	}

	fmt.Println(o)
}
