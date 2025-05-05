package main

import "fmt"

func main() {
	m := make(map[string]int)

	// assignments are trivial
	m["answer"] = 42
	fmt.Println("the answer is:", m["answer"])

	// mutating maps is trivial too (mutation is the devil but okay :/)
	m["answer"] = 49
	fmt.Println("the answer is:", m["answer"])

	// removing elems can be done via delete built in
	delete(m, "answer")
	fmt.Println("the answer is:", m["answer"]) // should return nil?
	// btw this is not "nil" this is the zero value of the value type which in case of int is zero
	// this could be very confusing because zero could be a perfectly valid value
	// and we would sometime expect nil to indicate not available and not zero
	// i guess this is why go error handles all over the place

	// test that the key is present with a two value assignment
	// first time we're introduces to "errors as values"
	elem, ok := m["answer"]
	fmt.Println("the answer is:", elem, "Present?", ok)

}
