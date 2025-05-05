package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{42, 99}

	// struct fields are accessed using a dot
	v.X = 4
	fmt.Println(v.X)
}
