package main

import "fmt"

// a map is a key value pair similar to dicts in python
// the "zero" value of a map is nil
// nil map has no keys nor can keys be added
// make func return a map of a given type initialized and ready to use

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	m["Stanford"] = Vertex{40.68433, -74.39967}
	m["MIT"] = Vertex{40.68433, -74.39967}

	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
}
