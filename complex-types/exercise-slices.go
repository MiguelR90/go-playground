package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	// could use append instead of make
	// this will grow the slice / array incrementally instead of all at ones
	p := make([][]uint8, dx)

	for i := range p {
		p[i] = make([]uint8, dy)

		for j := range p[i] {
			p[i][j] = uint8(i * j)
		}
	}

	return p

}

func main() {

	// p := Pic(3, 3)
	// p := Pic(30, 30)
	// fmt.Println(p)

	pic.Show(Pic)
	// here we will implement pic
}
