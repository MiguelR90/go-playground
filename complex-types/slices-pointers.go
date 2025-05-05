package main

import "fmt"

func main() {

	// slices are references / pointers into arrays
	// chaning the elements of a slice modifies the corresponding elems of the underlying array
	// other silces will also see the changes
	// NOTE: mutability os always a foot fun ( where is elixir )
	names := [4]string{
		"miguel",
		"emmie",
		"lydia",
		"mimi",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
