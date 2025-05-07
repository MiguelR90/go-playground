package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// interesting that this works with value or pointer receiver
func (p *Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"miguel", 35}
	b := Person{"emmie", 3}
	fmt.Println(a, b)
}
