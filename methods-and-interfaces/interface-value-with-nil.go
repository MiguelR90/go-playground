package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")

		// naked return? does this return nil or just void?
		return
	}

	fmt.Println(t.S)

}

func describe(i I) {
	fmt.Println("(%v, %T)\n", i, i)
}

func main() {

	var i I
	var t *T

	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}
