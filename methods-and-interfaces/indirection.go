package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	// NOTE: that this works even though v is a value and not a pointer
	// this is a convenience, go interprets v.Scale(2) as (&v).Scale(2)
	// since the Scale method has a pointer receiver!!!!!!
	// feels confusing and likely will take a little while to get used to
	v.Scale(2)
	ScaleFunc(&v, 10)

	// p is always a point so both method and func should work just fine
	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
