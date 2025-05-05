package main

import "fmt"

// lsp is complaing that i am redeclaring the struct here instead of just
// importing from the struct.go submodule? idk
type Vertex struct {
	X int
	Y int
}

func main() {

	v := Vertex{1, 2}
	// u is a pointer to v
	u := &v
	// to access the field X when we have a pointer u we could use (*u).X
	// that notation is cumbersome so the language permits to write u.X instead
	// syntax avoid the explicit deference of the pointer.
	// NOTE: easy to get this wrong and confused. potential foot gun.
	u.X = 1e3
	fmt.Println(*u)

}
