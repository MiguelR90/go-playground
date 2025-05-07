package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// a type implements and interface by implementing its methods
// there is no explict declaration of intent, no "implements" keyword
// implicit interfaces decouple the def of an interface from it's impl
// which could then appeear in any package without prearrangement!
// this is the polymorthism that Jose Valim talks so much about
// protocols in elixir and i guess (clojure)
// definitely the structural typing of python (which i guess they just borrowed from go)
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"helloworld"}
	i.M()
}
