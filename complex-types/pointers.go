package main

import "fmt"

func main() {

	i, j := 42, 42069

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer (NOTE: to modify the value you need to get a hold of the value)
	fmt.Println(i)  // see the new value of i (should print 21 since we've mutated original)

	p = &j         // point to j (NOTE: p has already been declared so we don't need :=)
	*p = *p / 37   // divide j through the pointer (NOTE: integer division thus no upcasting like python)
	fmt.Println(j) // see the new value of j
}
